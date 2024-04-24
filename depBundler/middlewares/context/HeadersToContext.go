package context

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/orchestd/dependencybundler/bundler/contextHeader"
	"github.com/orchestd/dependencybundler/interfaces/configuration"
	"github.com/orchestd/dependencybundler/interfaces/log"
	"github.com/orchestd/tokenauth"
	"time"
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

		token := c.Request.Header.Get("Token")
		if len(token) > 0 {
			_, protectedData, err := jwToken.ValidateAndGetData(context.Background(), time.Now(), token)
			if err != nil {
				logger.Error(context.Background(), "can't ValidateAndGetData token err:"+err.Error())
			} else {
				if journeytoken, ok := protectedData["journeytoken"]; ok {
					if s, ok := journeytoken.(string); ok {
						ctx := context.WithValue(c.Request.Context(), "journeytoken", s)
						c.Request = c.Request.WithContext(ctx)
					} else {
						logger.Error(context.Background(), "journeytoken is not string")
					}
				}
			}
		}

		c.Next()
		return
	}
}
