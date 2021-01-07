package log

import (
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/log"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/transport"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
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
		bodyCopy := new(bytes.Buffer)
		io.Copy(bodyCopy, c.Request.Body)
		bodyData := bodyCopy.Bytes()
		c.Request.Body = ioutil.NopCloser(bytes.NewReader(bodyData))
		blw := &bodyLogWriter{body: bytes.NewBuffer([]byte{}), ResponseWriter: c.Writer}
		c.Writer = blw
		reqJson := json.RawMessage(bodyData)
		start := time.Now().UTC()
		path := c.Request.URL.Path
		c.Next()
		transactionID, _ := c.Get("RequestId")

		rawBody := blw.body
		jsonmsg := json.RawMessage(rawBody.Bytes())

		end := time.Now().UTC()
		latency := end.Sub(start)
		entry := logger.WithFields(map[string]interface{}{
			"request"  : reqJson,
			"response" : jsonmsg,
			"status":     c.Writer.Status(),
			"method":     c.Request.Method,
			"path":       path,
			"ip":         c.ClientIP(),
			"requestId": transactionID,
			"duration":   latency,
			"user_agent": c.Request.UserAgent(),
		})
		if d, ok := c.Deadline(); ok {
			entry = entry.WithField("deadline", d)
		}
		var errorMsg string
		if len(c.Errors.String()) > 0 {
			srvErr,ok := c.Errors[0].Meta.(transport.IHttpLog)
			if ok{
				errorMsg = srvErr.GetAction()
				if srvErr.GetLogMessage() != nil{
					errorMsg += " - " + *srvErr.GetLogMessage()
				}
				entry = entry.WithError(c.Errors[0].Err)
			}
			entry = entry.WithField("source",srvErr.GetSource()).WithField("logValues",srvErr.GetLogValues())
		}
		if c.Writer.Status() >= 500{
			entry.Error(c,errorMsg)
		}else {
			entry.Info(c,"%s finished " , c.FullPath())
		}
	}
}

var allowedContentTypes = map[string]bool{
	"application/json": true,
	"application/xml":  true,
}

