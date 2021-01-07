package trace

import "github.com/gin-gonic/gin"

const TokenHeaderName  = "token"
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get(TokenHeaderName)
		if token != "" {
			c.Set(TokenHeaderName,token)
		}
		c.Writer.Header().Set("Token", token)
		c.Next()
	}
}
