package log

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/orchestd/dependencybundler/interfaces/log"
	"github.com/orchestd/dependencybundler/interfaces/transport"
	"google.golang.org/api/logging/v2"
	"net/http/httputil"
	"os"
	"time"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func GinLogHandlerMiddleware(logger log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		useReqRespLogger := os.Getenv("useReqRespLogger")
		if useReqRespLogger != "true" {
			c.Next()
			return
		}

		reqJson, _ := httputil.DumpRequest(c.Request, true)
		blw := &bodyLogWriter{
			body:           new(bytes.Buffer),
			ResponseWriter: c.Writer,
		}
		c.Writer = blw
		start := time.Now().UTC()

		c.Next()

		end := time.Now().UTC()
		latency := end.Sub(start)

		httpRequest := map[string]interface{}{
			"response": blw.body.Bytes(),
		}
		if c.Request.Method != "GET" && c.Request.Method != "DELETE" {
			httpRequest["request"] = reqJson
		}
		if d, ok := c.Deadline(); ok {
			httpRequest["deadline"] = d
		}
		entry := logger.WithFields(httpRequest)
		httpRequestStruct := logging.HttpRequest{
			Latency:         latency.String(),
			Protocol:        c.Request.Proto,
			Referer:         c.Request.Referer(),
			RemoteIp:        c.ClientIP(),
			RequestMethod:   c.Request.Method,
			RequestUrl:      c.Request.URL.String(),
			ResponseSize:    int64(c.Writer.Size()),
			Status:          int64(c.Writer.Status()),
			UserAgent:       c.Request.UserAgent(),
			ForceSendFields: nil,
			NullFields:      nil,
		}
		entry = entry.WithField("httpRequest", httpRequestStruct)
		var errorMsg string
		if len(c.Errors.String()) > 0 {
			srvErr, ok := c.Errors[0].Meta.(transport.IHttpLog)
			if ok {
				errorMsg = srvErr.GetAction()
				if srvErr.GetLogMessage() != nil {
					errorMsg += " - " + *srvErr.GetLogMessage()
				}
				entry = entry.WithError(c.Errors[0].Err)
			}
			entry = entry.WithField("source", srvErr.GetSource()).WithField("logValues", srvErr.GetLogValues())
		}
		if c.Writer.Status() >= 500 {
			entry.Error(c.Request.Context(), errorMsg)
		} else {
			entry.Info(c.Request.Context(), "%s finished ", c.FullPath())
		}
	}
}

var allowedContentTypes = map[string]bool{
	"application/json": true,
	"application/xml":  true,
}
