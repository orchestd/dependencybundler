package context

import (
	"context"
	"github.com/gin-gonic/gin"
)

// IpToContext middleware adds the client ip to the context
func IpToContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIp := c.ClientIP()
		ctx := context.WithValue(c.Request.Context(), "clientIp", clientIp)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
		return
	}
}
