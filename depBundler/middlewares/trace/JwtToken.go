package trace

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
)

const TokenHeaderName = "token"
const UberTraceId = "Uber-Trace-Id"

func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(TokenHeaderName)
		if token != "" {
			ctx := context.WithValue(c.Request.Context(), TokenHeaderName, token)
			c.Request = c.Request.WithContext(ctx)
		}
		c.Writer.Header().Set(TokenHeaderName, token)
		c.Next()
	}
}

func GetUberTraceId(c context.Context) string {
	return fmt.Sprint(c.Value(UberTraceId))
}
