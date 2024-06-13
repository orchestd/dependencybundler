package emptyLogger

import (
	"context"
	log2 "github.com/orchestd/dependencybundler/interfaces/log"
	"github.com/orchestd/log"
)

func NewEmptyLogger() log2.Logger {
	return emptyLogger{}
}

type emptyLogger struct{}

func (e emptyLogger) Trace(ctx context.Context, format string, args ...interface{}) {

}

func (e emptyLogger) Debug(ctx context.Context, format string, args ...interface{}) {

}

func (e emptyLogger) Info(ctx context.Context, format string, args ...interface{}) {

}

func (e emptyLogger) Warn(ctx context.Context, format string, args ...interface{}) {

}

func (e emptyLogger) Error(ctx context.Context, format string, args ...interface{}) {

}

func (e emptyLogger) Custom(ctx context.Context, level log.Level, skipAdditionalFrames int, format string, args ...interface{}) {

}

func (e emptyLogger) WithError(err error) log.Fields {
	return e
}

func (e emptyLogger) WithField(name string, value interface{}) log.Fields {
	return e
}

func (e emptyLogger) WithFields(fields map[string]interface{}) log.Fields {
	return e
}

func (e emptyLogger) Configuration() log.LoggerConfiguration {
	return nil
}
