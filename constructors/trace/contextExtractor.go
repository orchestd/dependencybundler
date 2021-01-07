package trace

import (
	"bitbucket.org/HeilaSystems/dependencybundler/depBundler"
	"bitbucket.org/HeilaSystems/log"
	"bitbucket.org/HeilaSystems/trace/bjaeger"
	"go.uber.org/fx"
)

func TraceInfoContextExtractorFxOption() fx.Option {
	return fx.Provide(
		fx.Annotated{
			Group: depBundler.FxGroupLoggerContextExtractors,
			Target: func() log.ContextExtractor {
				return bjaeger.TraceInfoExtractorFromContext
			},
		},
	)
}