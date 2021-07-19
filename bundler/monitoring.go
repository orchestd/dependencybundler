package bundler

import (
	"bitbucket.org/HeilaSystems/dependencybundler/constructors/monitoring"
	"bitbucket.org/HeilaSystems/dependencybundler/depBundler"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/configuration"
	monitoring2 "bitbucket.org/HeilaSystems/monitoring"
	"bitbucket.org/HeilaSystems/monitoring/bprometheus"
	"go.uber.org/fx"
)

func MonitoringFxOption() fx.Option {
	return fx.Options(
		fx.Provide(func(config configuration.Config)(monitoring2.Builder,error){
			name,err :=  config.Get(depBundler.DockerNameEnv).String()
			if err != nil {
				return nil,err
			}
			return bprometheus.Builder().SetNamespace(name),nil
		}),
		fx.Provide(monitoring.DefaultMonitor),
	)
}
