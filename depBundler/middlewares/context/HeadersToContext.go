package context

import (
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/configuration"
	"bitbucket.org/HeilaSystems/session/sessionresolver"
	"context"
	"github.com/gin-gonic/gin"
)

func HeadersToContext(config configuration.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		headers, err := config.Get("contextHeaders").StringSlice()
		if err != nil {
			c.Next()
			return
		}
		headers = append(headers, sessionresolver.DataNowKey)
		for _, header := range headers {
			h := c.GetHeader(header)
			if len(h) > 0 {
				ctx := context.WithValue(c.Request.Context(),header,h)
				c.Request = c.Request.WithContext(ctx)
			}
		}
		c.Next()
		return
	}
}
