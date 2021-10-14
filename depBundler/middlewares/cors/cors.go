package cors

import (
	ginCors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	corsConfig := ginCors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, []string{"token", "access-token", "refresh-token"}...)
	corsConfig.ExposeHeaders = append(corsConfig.AllowHeaders, []string{"token", "access-token", "refresh-token"}...)
	return ginCors.New(corsConfig)
}
