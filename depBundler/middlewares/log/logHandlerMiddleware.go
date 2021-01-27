package log

import (
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/log"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/transport"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/logging/v2"
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
		//path := c.Request.URL.Path
		c.Next()

		rawBody := blw.body
		jsonmsg := json.RawMessage(string(rawBody.Bytes()))

		end := time.Now().UTC()
		latency := end.Sub(start)
		//gEntry := logging.Entry{HTTPRequest:&logging.HTTPRequest{
		//	Request:                        c.Request,
		//	Status:                         c.Writer.Status(),
		//	ResponseSize:                   int64(c.Writer.Size()),
		//	Latency:                        latency,
		//	RemoteIP:                       c.ClientIP(),
		//}}
		httpRequest := map[string]interface{}{
			"response" : jsonmsg,
		}
		if c.Request.Method != "GET" && c.Request.Method != "DELETE" {
			httpRequest["request"] = reqJson
		}
		if d, ok := c.Deadline(); ok {
			httpRequest["deadline"] =  d
		}
		entry := logger.WithFields(httpRequest)
		httpRequestStruct := logging.HttpRequest{
			Latency:                        latency.String(),
			Protocol:                      	c.Request.Proto,
			Referer:                        c.Request.Referer(),
			RemoteIp:                       c.ClientIP(),
			RequestMethod : c.Request.Method,
			RequestUrl:                     c.Request.URL.String(),
			ResponseSize:                   int64(c.Writer.Size()),
			Status:                         int64(c.Writer.Status()),
			UserAgent:                      c.Request.UserAgent(),
			ForceSendFields:                nil,
			NullFields:                     nil,
		}
		entry = entry.WithField("httpRequest" , httpRequestStruct)
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
			entry.Error(c.Request.Context(),errorMsg)
		}else {
			entry.Info(c.Request.Context(),"%s finished " , c.FullPath())
		}
	}
}

var allowedContentTypes = map[string]bool{
	"application/json": true,
	"application/xml":  true,
}

