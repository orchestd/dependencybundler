package server

import (
	"github.com/gin-gonic/gin"
	"github.com/orchestd/dependencybundler/depBundler/middlewares/context"
	log2 "github.com/orchestd/dependencybundler/depBundler/middlewares/log"
	"github.com/orchestd/dependencybundler/depBundler/middlewares/trace"
	"github.com/orchestd/dependencybundler/interfaces/configuration"
	"github.com/orchestd/dependencybundler/interfaces/log"
)

func DefaultLogHandlerMiddleware(logImpl log.Logger) gin.HandlerFunc {
	return log2.GinLogHandlerMiddleware(logImpl)
}
func DefaultHeadersToContext(conf configuration.Config) gin.HandlerFunc {
	return context.HeadersToContext(conf)
}
func DefaultBasicRequestId() gin.HandlerFunc {
	return trace.RequestId()
}

func DefaultJwtToken() gin.HandlerFunc {
	return trace.JwtToken()
}
