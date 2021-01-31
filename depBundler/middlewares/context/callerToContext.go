package context

import (
	"context"
	"github.com/gin-gonic/gin"
)

func CallerToContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check for incoming header, use it if exists
		caller :=c.GetHeader("Caller")

		if len(caller) > 0 {
			ctx := context.WithValue(c.Request.Context(),"caller",caller)
			c.Request = c.Request.WithContext(ctx)
		}
		c.Next()
	}
}
