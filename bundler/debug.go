package bundler

import (
	"github.com/orchestd/debug"
	"github.com/orchestd/debug/debugHandlers"
	"github.com/orchestd/debug/debugHandlers/repos/trace/jaeger"
	debug2 "github.com/orchestd/dependencybundler/constructors/debug"
	"go.uber.org/fx"
)

func DebugFxOption() fx.Option {
	return fx.Options(
		fx.Provide(jaeger.NewJaegerApiClient),
		fx.Provide(func() debug.Builder {
			builder := debugHandlers.DebugBuilder()
			return builder
		}),
		fx.Provide(debug2.DefaultDebug),
	)
}
