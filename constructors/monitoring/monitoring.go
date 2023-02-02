package monitoring

import (
	"context"
	"github.com/orchestd/dependencybundler/interfaces/configuration"
	log2 "github.com/orchestd/dependencybundler/interfaces/log"
	"github.com/orchestd/dependencybundler/interfaces/monitoring"
	"github.com/orchestd/log"
	monitor "github.com/orchestd/monitoring"
	monitoringWrapper "github.com/orchestd/monitoring/wrapper"
	"github.com/orchestd/sharedlib/consts"

	"go.uber.org/fx"
)

const (
	FxGroupMonitorContextExtractors = "monitorContextExtractors"
)

type monitorDeps struct {
	fx.In
	Lf                fx.Lifecycle
	Config            configuration.Config
	Log               log2.Logger
	Builder           monitor.Builder
	ContextExtractors []monitor.ContextExtractor `group:"monitorContextExtractors"`
}

// DefaultMonitor is a constructor that will create a Metrics reporter based on values from the Config Map
// such as
//
// 	- Tags: we will look for default tags using mortar.MonitorTagsKey within the configuration map
//
func DefaultMonitor(deps monitorDeps) (monitoring.Metrics, error) {
	tags, _ := deps.Config.Get(consts.MonitorTags).StringMapString() // can be empty
	reporter := monitoringWrapper.Builder().SetTags(tags).AddExtractors(deps.ContextExtractors...).DoOnError(func(err error) {
		deps.Log.WithError(err).Custom(nil, log.WarnLevel, 2, "monitoring error")
	}).Build(deps.Builder)

	deps.Lf.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return reporter.Connect(ctx)
		},
		OnStop: func(ctx context.Context) error {
			return reporter.Close(ctx)
		},
	})
	return reporter.Metrics(), nil
}
