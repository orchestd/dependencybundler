package context

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/orchestd/dependencybundler/bundler/contextHeader"
	"github.com/orchestd/dependencybundler/interfaces/configuration"
)

func HeadersToContext(config configuration.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		headers, err := config.Get("contextHeaders").StringSlice()
		if err != nil {
			c.Next()
			return
		}
		headers = append(headers, contextHeader.AlwaysCopyHeaders...)
		for _, header := range headers {
			h := c.GetHeader(header)
			if len(h) > 0 {
				ctx := context.WithValue(c.Request.Context(), header, h)
				c.Request = c.Request.WithContext(ctx)
			}
		}
		c.Next()
		return
	}
}
