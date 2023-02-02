package logger

import (
	"context"
	"github.com/orchestd/dependencybundler/depBundler"
	"github.com/orchestd/dependencybundler/interfaces/configuration"
	log2 "github.com/orchestd/dependencybundler/interfaces/log"
	"github.com/orchestd/log"
	"github.com/orchestd/sharedlib/consts"
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
	if levelValue := deps.Config.Get(consts.MinimumSeverityLevel); levelValue.IsSet() {
		if key, err := levelValue.String(); err == nil {
			logLevel = log.ParseLevel(key)
		}
	}

	builder := deps.LoggerBuilder.SetLevel(logLevel).IncrementSkipFrames(compensateDefaultLogger)

	return log.CreateMortarLogger(builder, append(deps.ContextExtractors, deps.selfStaticFieldsContextExtractor)...)
}

func (d loggerDeps) selfStaticFieldsContextExtractor(_ context.Context) map[string]interface{} {
	output := make(map[string]interface{})
	if dockerName, err := d.Config.GetServiceName(); err == nil {
		output["app"] = dockerName
	}
	info := depBundler.GetBuildInformation()
	if len(info.Hostname) > 0 {
		output["host"] = info.Hostname
	}
	if len(info.Version) > 0 {
		output["version"] = info.Version
	}
	if len(info.BuildTag) > 0 {
		output["buildNo"] = info.BuildTag
	}
	if len(info.GitCommit) > 0 {
		output["gitCommit"] = info.GitCommit
	}
	return output
}
