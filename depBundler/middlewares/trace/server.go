package trace

import (
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	"io"
	"io/ioutil"
	"net/http/httputil"
	"net/url"
)

func HttpTracingUnaryServerInterceptor(deps tracingDeps) gin.HandlerFunc {
	return func(c *gin.Context) {

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
		if v, err := httputil.DumpRequest(c.Request, true); err == nil {
			addBodyToSpan(sp, "request", v)
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

		bodyCopy := new(bytes.Buffer)
		io.Copy(bodyCopy, c.Request.Body)

		bodyData := bodyCopy.Bytes()
		c.Request.Body = ioutil.NopCloser(bytes.NewReader(bodyData))
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
