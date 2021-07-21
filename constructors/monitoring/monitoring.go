package monitoring

import (
	"bitbucket.org/HeilaSystems/dependencybundler/depBundler"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/configuration"
	log2 "bitbucket.org/HeilaSystems/dependencybundler/interfaces/log"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/monitoring"
	"bitbucket.org/HeilaSystems/log"
	monitor "bitbucket.org/HeilaSystems/monitoring"
	monitoringWrapper "bitbucket.org/HeilaSystems/monitoring/wrapper"
	"context"

	"go.uber.org/fx"
)

const (
	FxGroupMonitorContextExtractors = "monitorContextExtractors"
)

type monitorDeps struct{
	fx.In
	Lf fx.Lifecycle
	Config            configuration.Config
	Log log2.Logger
	Builder monitor.Builder
	ContextExtractors []monitor.ContextExtractor `group:"monitorContextExtractors"`
}

// DefaultMonitor is a constructor that will create a Metrics reporter based on values from the Config Map
// such as
//
// 	- Tags: we will look for default tags using mortar.MonitorTagsKey within the configuration map
//
func DefaultMonitor(deps monitorDeps) (monitoring.Metrics,error) {
	tags, _ := deps.Config.Get(depBundler.MonitorTags).StringMapString() // can be empty
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
	return reporter.Metrics(),nil
}