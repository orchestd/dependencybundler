package context

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/orchestd/dependencybundler/bundler/contextHeader"
	"github.com/orchestd/dependencybundler/depBundler/middlewares/trace"
	"github.com/orchestd/dependencybundler/interfaces/configuration"
	"github.com/orchestd/dependencybundler/interfaces/log"
	"github.com/orchestd/tokenauth"
)

func HeadersToContext(config configuration.Config, jwToken tokenauth.TokenBase, logger log.Logger) gin.HandlerFunc {
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

		ok, journeytoken, err := trace.GetJourneyTokenFromRequestToken(c, jwToken)
		if err != nil {
			logger.Error(context.Background(), "can't exec HeadersToContext for journeytoken. err: "+err.Error())
		}
		if ok {
			ctx := context.WithValue(c.Request.Context(), "journeytoken", journeytoken)
			c.Request = c.Request.WithContext(ctx)
		}

		c.Next()
		return
	}
}
