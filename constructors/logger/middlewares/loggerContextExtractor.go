package middlewares

import (
	"github.com/orchestd/dependencybundler/interfaces/configuration"
	"github.com/orchestd/log"
	"github.com/orchestd/log/middlewares"
)

func DefaultLoggerIncomingContextExtractor(config configuration.Config) log.ContextExtractor {
	headers, err := config.Get("contextHeaders").StringSlice()
	if err != nil {
		return middlewares.LoggerIncomingContextExtractor([]string{})
	}
	return middlewares.LoggerIncomingContextExtractor(headers)
}
