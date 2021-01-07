package context

import (
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/configuration"
	"github.com/gin-gonic/gin"
)

func HeadersToContext(config configuration.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		headers, err := config.Get("contextHeaders").StringSlice()
		if err != nil {
			c.Next()
			return
		}
		for _, header := range headers {
			h := c.GetHeader(header)
			if len(h) > 0 {
				c.Set(header, h)
			}
		}
		c.Next()
		return
	}
}
