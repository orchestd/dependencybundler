package bundler

import (
	"bitbucket.org/HeilaSystems/dependencybundler/depBundler"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/configuration"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/log"
	"bitbucket.org/HeilaSystems/trace/bjaeger"
	"context"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/fx"
)


func TracerFxOption() fx.Option {
	return fx.Provide(JaegerBuilder)
}

func JaegerBuilder(lc fx.Lifecycle, config configuration.Config, logger log.Logger) (opentracing.Tracer, error) {
	dockerName, err := config.Get(depBundler.DockerNameEnv).String()
	if err != nil {
		return nil, err
	}
	openTracer, err := bjaeger.Builder().
		SetServiceName(dockerName).
		AddOptions(bjaeger.BricksLoggerOption(logger)). // verbose logging,
		Build()
	if err != nil {
		return nil, err
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return openTracer.Connect(ctx)
		},
		OnStop: func(ctx context.Context) error {
			return openTracer.Close(ctx)
		},
	})
	return openTracer.Tracer(), nil
}
