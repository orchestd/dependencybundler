package bundler

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/orchestd/dependencybundler/interfaces/configuration"
	"github.com/orchestd/dependencybundler/interfaces/log"
	"github.com/orchestd/sharedlib/consts"
	"github.com/orchestd/trace/bjaeger"
	"go.uber.org/fx"
)

type Tracer opentracing.Tracer

func TracerFxOption() fx.Option {
	return fx.Provide(JaegerBuilder)
}

func JaegerBuilder(lc fx.Lifecycle, config configuration.Config, logger log.Logger) (opentracing.Tracer, Tracer, error) {
	dockerName, err := config.Get(consts.ServiceNameEnv).String()
	if err != nil {
		return nil, nil, err
	}
	openTracer, err := bjaeger.Builder().
		SetServiceName(dockerName).
		AddOptions(bjaeger.BricksLoggerOption(logger)). // verbose logging,
		Build()
	if err != nil {
		return nil, nil, err
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return openTracer.Connect(ctx)
		},
		OnStop: func(ctx context.Context) error {
			return openTracer.Close(ctx)
		},
	})
	t := openTracer.Tracer()
	return t, Tracer(t), nil
}
