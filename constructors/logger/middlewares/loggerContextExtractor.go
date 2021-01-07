package middlewares

import (
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/configuration"
	"bitbucket.org/HeilaSystems/log"
	"bitbucket.org/HeilaSystems/log/middlewares"
)

func DefaultLoggerIncomingContextExtractor(config configuration.Config)log.ContextExtractor {
	headers , err := config.Get("contextHeaders").StringSlice()
	if err != nil {
		return middlewares.LoggerIncomingContextExtractor([]string{})
	}
	return middlewares.LoggerIncomingContextExtractor(headers)
}