package trace

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

const TokenHeaderName  = "token"
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.Request.Header)
		token := c.GetHeader(TokenHeaderName)
		if token != "" {
			c.Set(TokenHeaderName,token)
		}
		c.Writer.Header().Set(TokenHeaderName, token)
		c.Next()
	}
}
