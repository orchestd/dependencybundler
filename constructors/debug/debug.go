package debug

import (
	"bitbucket.org/HeilaSystems/debug"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/configuration"
	log2 "bitbucket.org/HeilaSystems/dependencybundler/interfaces/log"
)

func DefaultDebug(traceApi debug.TraceApi,builder debug.Builder,config configuration.Config,logger log2.Logger) (debug.Debug,error) {
	if debug, err := config.Get("debugMode").Bool();err !=nil {
		return nil, nil
	} else if debug {
		return builder.SetTraceRepo(traceApi).Build()
	} else  {
		return nil, nil
	}
}
