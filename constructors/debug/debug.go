package debug

import (
	"github.com/orchestd/debug"
	"github.com/orchestd/dependencybundler/interfaces/configuration"
	log2 "github.com/orchestd/dependencybundler/interfaces/log"
)

func DefaultDebug(traceApi debug.TraceApi, builder debug.Builder, config configuration.Config, logger log2.Logger) (debug.Debug, error) {
	if debug, err := config.Get("debugMode").Bool(); err != nil {
		return nil, nil
	} else if debug {
		return builder.SetTraceRepo(traceApi).Build()
	} else {
		return nil, nil
	}
}
