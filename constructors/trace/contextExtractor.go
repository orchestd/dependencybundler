package trace

import (
	"github.com/orchestd/log"
	"github.com/orchestd/sharedlib/consts"
	"github.com/orchestd/trace/bjaeger"
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
