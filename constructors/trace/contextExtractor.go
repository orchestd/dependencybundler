package trace

import (
	"bitbucket.org/HeilaSystems/log"
	"bitbucket.org/HeilaSystems/sharedlib/consts"
	"bitbucket.org/HeilaSystems/trace/bjaeger"
	"go.uber.org/fx"
)

func TraceInfoContextExtractorFxOption() fx.Option {
	return fx.Provide(
		fx.Annotated{
			Group: consts.FxGroupLoggerContextExtractors,
			Target: func() log.ContextExtractor {
				return bjaeger.TraceInfoExtractorFromContext
			},
		},
	)
}
