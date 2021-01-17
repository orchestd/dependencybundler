package trace

import (
	"context"
	"github.com/gin-gonic/gin"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check for incoming header, use it if exists
		requestID :=c.GetHeader("Uber-Trace-Id")

		// Expose it for use in the application
		ctx := context.WithValue(c.Request.Context(),"Uber-Trace-Id",requestID)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
