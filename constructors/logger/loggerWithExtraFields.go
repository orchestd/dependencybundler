package logger

import (
	"context"
	log2 "github.com/orchestd/dependencybundler/interfaces/log"
	"github.com/orchestd/log"
	"github.com/orchestd/sharedlib/consts"
)

const extraFieldsKey = "loggerExtraFields"

func ExtraFieldsLogger(deps LoggerDeps) log2.Logger {
	var logLevel = log.DebugLevel

	if levelValue := deps.Config.Get(consts.MinimumSeverityLevel); levelValue.IsSet() {
		if key, err := levelValue.String(); err == nil {
			logLevel = log.ParseLevel(key)
		}
	}
	if extraFields, err := deps.Config.Get(extraFieldsKey).StringSlice(); err == nil {
		deps.ContextExtractors = append(deps.ContextExtractors, extraFieldsContextExtractors(extraFields).Extract)
	}

	builder := deps.LoggerBuilder.SetLevel(logLevel).IncrementSkipFrames(compensateDefaultLogger)

	return log.CreateMortarLogger(builder, append(deps.ContextExtractors, deps.selfStaticFieldsContextExtractor)...)
}

type extraFieldsContextExtractors []string

func (h extraFieldsContextExtractors) Extract(ctx context.Context) map[string]interface{} {
	var output = make(map[string]interface{})
	for _, s := range h {
		output[s] = ctx.Value(s)
	}
	return output
}
