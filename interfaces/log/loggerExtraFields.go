package log

import (
	"context"
	"github.com/orchestd/log"
)

type loggerWithExtraFields struct {
	Logger
	extraFields []string
}

func NewLoggerWithExtraFields(logger Logger, extraFields []string) Logger {
	return loggerWithExtraFields{logger, extraFields}
}

func (l loggerWithExtraFields) Debug(ctx context.Context, format string, args ...interface{}) {
	l.log(ctx, log.DebugLevel, format, args...)
}

func (l loggerWithExtraFields) Info(ctx context.Context, format string, args ...interface{}) {
	l.log(ctx, log.InfoLevel, format, args...)
}

func (l loggerWithExtraFields) Warn(ctx context.Context, format string, args ...interface{}) {
	l.log(ctx, log.WarnLevel, format, args...)
}

func (l loggerWithExtraFields) Error(ctx context.Context, format string, args ...interface{}) {
	l.log(ctx, log.ErrorLevel, format, args...)
}

func (l loggerWithExtraFields) log(ctx context.Context, level log.Level, format string, args ...interface{}) {
	if fields, ok := l.contextExtraValuesFields(ctx); ok {
		l.Logger.WithFields(fields).Custom(ctx, level, 2, format, args...)
	} else {
		l.Logger.Custom(ctx, level, 2, format, args...)
	}
}

func (l loggerWithExtraFields) contextExtraValuesFields(ctx context.Context) (map[string]interface{}, bool) {
	var hasExtraField bool
	fields := make(map[string]interface{})
	for _, key := range l.extraFields {
		fieldValue := ctx.Value(key)
		if fieldValue != nil {
			hasExtraField = true
			fields[key] = fieldValue
		}
	}

	return fields, hasExtraField
}
