package trace

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"io"
	"io/ioutil"
	"net/http/httputil"
)



func HttpTracingUnaryServerInterceptor(deps tracingDeps) gin.HandlerFunc {
	return func(c *gin.Context) {

		if deps.Tracer == nil {
			c.Next()
			return
		}
		carrier := opentracing.HTTPHeadersCarrier(c.Request.Header)
		ctx, _ := deps.Tracer.Extract(opentracing.HTTPHeaders, carrier)
		op := "HTTP " + c.Request.Method
		sp := deps.Tracer.StartSpan(op, ext.RPCServerOption(ctx))
		ext.HTTPMethod.Set(sp, c.Request.Method)

		ext.HTTPUrl.Set(sp, (c.Request.URL.String()))
		componentName,_ := deps.Config.Get("DOCKER_NAME").String()
		ext.Component.Set(sp, componentName)
		defer sp.Finish()
		if v ,err := httputil.DumpRequest(c.Request , true);err == nil {
			addBodyToSpan(sp, "request", v)
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

		c.Next()
		ext.HTTPStatusCode.Set(sp, uint16(c.Writer.Status()))
		rawBody := blw.body
		if len(c.Errors.String()) > 0 {
			c.Errors.String()
			if c.Errors[0].Err != nil {
				ext.LogError(sp, c.Errors[0].Err)
			}
		}
		if s := c.Request.Context().Value("status");s != nil {
			sp.SetTag("dependencyBundler.status" , s)
		}
		if um := c.Request.Context().Value("userMessageId");um != nil {
			sp.SetTag("dependencyBundler.id" , um)
		}
		addBodyToSpan(sp , "response-headers" , c.Writer.Header())
		addBodyToSpan(sp, "response", rawBody.Bytes())
	}
}