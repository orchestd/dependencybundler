package bundler

import (
	"bitbucket.org/HeilaSystems/debug"
	"bitbucket.org/HeilaSystems/debug/debugHandlers"
	"bitbucket.org/HeilaSystems/debug/debugHandlers/repos/trace/jaeger"
	debug2 "bitbucket.org/HeilaSystems/dependencybundler/constructors/debug"
	"go.uber.org/fx"
)

func DebugFxOption() fx.Option {
	return fx.Options(
		fx.Provide(jaeger.NewJaegerApiClient),
		fx.Provide(func() debug.Builder{
			builder := debugHandlers.DebugBuilder()
			return builder
		}),
		fx.Provide(debug2.DefaultDebug),
	)
}

