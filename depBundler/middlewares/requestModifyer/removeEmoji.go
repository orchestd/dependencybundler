package requestModifyer

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/orchestd/dependencybundler/interfaces/configuration"
	"github.com/orchestd/sharedlib/stringHelpers"
	"io"
	"io/ioutil"
)

func RequestModifyer(config configuration.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyCopy := new(bytes.Buffer)
		io.Copy(bodyCopy, c.Request.Body)
		bodyData := bodyCopy.Bytes()
		bodyData = []byte(stringHelpers.RemoveAllEmojis(string(bodyData)))
		c.Request.Body = ioutil.NopCloser(bytes.NewReader(bodyData))

		c.Next()
		return
	}
}
