package trace

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/textproto"
	"net/url"
	"os"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/orchestd/servicereply"
	"github.com/uber/jaeger-client-go"
)

func HttpTracingUnaryServerInterceptor(deps tracingDeps) gin.HandlerFunc {
	return func(c *gin.Context) {
		disableTracer := os.Getenv("disableTracer")
		if disableTracer == "true" {
			c.Next()
			return
		}
		if deps.Tracer == nil {
			c.Next()
			return
		}
		carrier := opentracing.HTTPHeadersCarrier(c.Request.Header)
		ctx, _ := deps.Tracer.Extract(opentracing.HTTPHeaders, carrier)
		URLQueryUnescape, _ := url.QueryUnescape(c.Request.URL.String())
		op := "HTTP " + c.Request.Method + URLQueryUnescape
		sp := deps.Tracer.StartSpan(op, ext.RPCServerOption(ctx))
		ext.HTTPMethod.Set(sp, c.Request.Method)
		host := c.Request.Host
		if host == "" && c.Request.URL != nil {
			host = c.Request.URL.Host
		}
		ext.PeerHostname.Set(sp, host)
		ext.HTTPUrl.Set(sp, URLQueryUnescape)
		componentName, _ := deps.Config.GetServiceName()
		ext.Component.Set(sp, componentName)
		defer sp.Finish()

		_, requestTracePayload, err := getRequestTracePayload(c.Request, deps)
		if err == nil {
			addBodyToSpan(sp, "request", requestTracePayload)
		}

		if c.Request.URL.Query().Has("journeytoken") {
			journeytoken := c.Request.URL.Query().Get("journeytoken")
			sp.SetTag("journeytoken", journeytoken)
		} else {
			ok, journeytoken, err := GetJourneyTokenFromRequestToken(c, deps.JWToken)
			if err != nil {
				deps.Logger.Error(context.Background(), "can't exec HttpTracingUnaryServerInterceptor for journeytoken. err: "+err.Error())
			}
			if ok {
				sp.SetTag("journeytoken", journeytoken)
			}
		}

		blw := &bodyLogWriter{body: bytes.NewBuffer([]byte{}), ResponseWriter: c.Writer}
		c.Writer = blw

		// call handler
		c.Request = c.Request.WithContext(
			opentracing.ContextWithSpan(c.Request.Context(), sp))
		if debugMode, err := deps.Config.Get("debugmode").Bool(); err == nil {
			if debugMode {
				if len(c.Writer.Header().Get("Uber-Trace-Id")) == 0 {
					if sc, ok := sp.Context().(jaeger.SpanContext); ok {
						c.Header("Uber-Trace-Id", sc.TraceID().String())
					}
				}
			}
		}
		c.Next()
		ext.HTTPStatusCode.Set(sp, uint16(c.Writer.Status()))
		rawBody := blw.body
		if len(c.Errors.String()) > 0 {
			c.Errors.String()
			if c.Errors[0].Err != nil {
				ext.LogError(sp, c.Errors[0].Err)
			}
		}

		if s := c.Request.Context().Value("traceValues"); s != nil {
			if traceValues, ok := s.(servicereply.ValuesMap); ok {
				for k, v := range traceValues {
					sp.SetTag(k, v)
				}
			}
		}

		if s := c.Request.Context().Value("status"); s != nil {
			sp.SetTag("dependencyBundler.status", s)
		}
		if um := c.Request.Context().Value("userMessageId"); um != nil {
			sp.SetTag("dependencyBundler.id", um)
		}
		addBodyToSpan(sp, "response-headers", c.Writer.Header())
		addBodyToSpan(sp, "response", rawBody.Bytes())
	}
}

func getRequestTracePayload(req *http.Request, deps tracingDeps) ([]byte, interface{}, error) {
	bodyData, err := readAndRestoreRequestBody(req)
	if err != nil {
		return nil, nil, err
	}

	if isMultipartFormRequest(req) {
		reqDump, err := buildMultipartTraceDump(req, bodyData, deps)
		return bodyData, reqDump, err
	}

	reqDump, err := httputil.DumpRequest(cloneRequestWithBody(req, bodyData), true)
	if err != nil {
		return bodyData, nil, err
	}

	return bodyData, reqDump, nil
}

func readAndRestoreRequestBody(req *http.Request) ([]byte, error) {
	if req.Body == nil {
		return nil, nil
	}

	bodyData, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	req.Body = io.NopCloser(bytes.NewReader(bodyData))

	return bodyData, nil
}

func isMultipartFormRequest(req *http.Request) bool {
	contentType := req.Header.Get("Content-Type")
	if contentType == "" {
		return false
	}

	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		return false
	}

	return mediaType == "multipart/form-data"
}

func buildMultipartTraceDump(req *http.Request, bodyData []byte, deps tracingDeps) ([]byte, error) {
	clonedRequest := cloneRequestWithBody(req, bodyData)
	if err := clonedRequest.ParseMultipartForm(32 << 20); err != nil {
		return nil, err
	}
	if clonedRequest.MultipartForm == nil {
		return httputil.DumpRequest(clonedRequest, true)
	}
	defer clonedRequest.MultipartForm.RemoveAll()

	traceBody := bytes.NewBuffer(nil)
	traceWriter := multipart.NewWriter(traceBody)

	if boundary, ok := getMultipartBoundary(req); ok {
		if err := traceWriter.SetBoundary(boundary); err != nil {
			return nil, fmt.Errorf("set multipart boundary: %w", err)
		}
	}

	for _, fieldName := range extractKeys(clonedRequest.MultipartForm.Value) {
		for _, value := range clonedRequest.MultipartForm.Value[fieldName] {
			if err := traceWriter.WriteField(fieldName, value); err != nil {
				return nil, fmt.Errorf("write multipart field %q: %w", fieldName, err)
			}
		}
	}

	for _, fieldName := range extractKeys(clonedRequest.MultipartForm.File) {
		files := clonedRequest.MultipartForm.File[fieldName]
		sort.SliceStable(files, func(i, j int) bool {
			return files[i].Filename < files[j].Filename
		})

		for _, file := range files {
			part, err := traceWriter.CreatePart(buildMultipartTracePartHeader(fieldName, file))
			if err != nil {
				return nil, fmt.Errorf("create multipart file part %q: %w", file.Filename, err)
			}

			if _, err := io.WriteString(part, saveFile(file, deps)); err != nil {
				return nil, fmt.Errorf("write multipart file link %q: %w", file.Filename, err)
			}

		}
	}

	if err := traceWriter.Close(); err != nil {
		return nil, fmt.Errorf("close multipart trace writer: %w", err)
	}

	traceRequest := cloneRequestWithBody(req, traceBody.Bytes())
	traceRequest.Header.Set("Content-Type", traceWriter.FormDataContentType())

	return httputil.DumpRequest(traceRequest, true)
}

func saveFile(fileHeader *multipart.FileHeader, deps tracingDeps) string {
	type UploadToBucketResponse struct {
		Data struct {
			UploadedTo string `json:"uploadedTo"`
		} `json:"data"`
	}

	saveFileForTrace, _ := deps.Config.Get("saveFileForTrace").Bool()
	if saveFileForTrace {
		file, err := fileHeader.Open()
		if err != nil {
			deps.Logger.Error(context.Background(), "cannot open file: "+err.Error())
			return "error saving file"
		}

		fileBytes, err := io.ReadAll(file)
		closeErr := file.Close()
		if err != nil {
			deps.Logger.Error(context.Background(), "cannot read file: "+err.Error())
			return "error saving file"
		}
		if closeErr != nil {
			deps.Logger.Error(context.Background(), "cannot close file: "+closeErr.Error())
			return "error saving file"
		}

		req := map[string]interface{}{
			"fileName": fileHeader.Filename,
			"content":  fileBytes,
		}
		bucketUrl, err := deps.Config.Get("bucketUrl").String()
		if err != nil {
			deps.Logger.Error(context.Background(), "can't get bucketUrl from conf: "+err.Error())
			return "error saving file"
		}
		err = deps.Client.ExternalPost(context.Background(), req, bucketUrl, "google/storage/byte", &req, nil, "json")
		if err != nil {
			deps.Logger.Error(context.Background(), "can't save file to bucket service: "+err.Error())
			return "error saving file"
		}
	}
	return "saveFileForTrace is not true"
}

func cloneRequestWithBody(req *http.Request, bodyData []byte) *http.Request {
	clonedRequest := req.Clone(req.Context())
	clonedRequest.Body = io.NopCloser(bytes.NewReader(bodyData))
	clonedRequest.ContentLength = int64(len(bodyData))

	return clonedRequest
}

func buildMultipartTracePartHeader(fieldName string, file *multipart.FileHeader) textproto.MIMEHeader {
	header := make(textproto.MIMEHeader, len(file.Header))
	for k, v := range file.Header {
		header[k] = append([]string(nil), v...)
	}

	return header
}

func getMultipartBoundary(req *http.Request) (string, bool) {
	_, params, err := mime.ParseMediaType(req.Header.Get("Content-Type"))
	if err != nil {
		return "", false
	}

	boundary, ok := params["boundary"]
	return boundary, ok
}

func extractKeys[T any](data map[string]T) []string {
	keys := make([]string, 0, len(data))
	for key := range data {
		keys = append(keys, key)
	}

	return keys
}
