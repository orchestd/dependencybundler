package bundler

import (
	"bitbucket.org/HeilaSystems/dependencybundler/constructors/monitoring"
	"bitbucket.org/HeilaSystems/dependencybundler/depBundler"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/configuration"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/transport"
	monitoring2 "bitbucket.org/HeilaSystems/monitoring"
	"bitbucket.org/HeilaSystems/monitoring/bprometheus"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

func PrometheusHandler()  gin.HandlerFunc {
	h:= promhttp.Handler()
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}