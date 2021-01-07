package server

import (
	"bitbucket.org/HeilaSystems/dependencybundler/depBundler/middlewares/context"
	log2 "bitbucket.org/HeilaSystems/dependencybundler/depBundler/middlewares/log"
	"bitbucket.org/HeilaSystems/dependencybundler/depBundler/middlewares/trace"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/configuration"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/log"
	"github.com/gin-gonic/gin"
)

func DefaultLogHandlerMiddleware(logImpl log.Logger) gin.HandlerFunc {
	return log2.GinLogHandlerMiddleware(logImpl)
}
func DefaultHeadersToContext(conf configuration.Config) gin.HandlerFunc{
	return context.HeadersToContext(conf)
}
func DefaultBasicRequestId() gin.HandlerFunc{
	return trace.RequestId()
}

func DefaultJwtToken() gin.HandlerFunc{
	return trace.JwtToken()
}