package logger

import (
	"bitbucket.org/HeilaSystems/dependencybundler/depBundler"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/configuration"
	log2 "bitbucket.org/HeilaSystems/dependencybundler/interfaces/log"
	"bitbucket.org/HeilaSystems/log"
	"go.uber.org/fx"
)

const compensateDefaultLogger = 1

type loggerDeps struct {
	fx.In
	Config            configuration.Config
	LoggerBuilder     log.Builder
	ContextExtractors []log.ContextExtractor `group:"loggerContextExtractors"`
}


func DefaultLogger(deps loggerDeps) log2.Logger {
	var logLevel = log.DebugLevel
	if levelValue := deps.Config.Get(depBundler.MinimumSeverityLevel); levelValue.IsSet() {
		if key  ,err := levelValue.String();err == nil {
			logLevel = log.ParseLevel(key)
		}
	}

	builder := deps.LoggerBuilder.SetLevel(logLevel).IncrementSkipFrames(compensateDefaultLogger)

	return log.CreateMortarLogger(builder, append(deps.ContextExtractors)...)
}