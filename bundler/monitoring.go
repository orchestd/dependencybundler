package bundler

import (
	"github.com/gin-gonic/gin"
	"github.com/orchestd/dependencybundler/constructors/monitoring"
	"github.com/orchestd/dependencybundler/interfaces/configuration"
	monitoring2 "github.com/orchestd/monitoring"
	"github.com/orchestd/monitoring/bprometheus"
	"github.com/orchestd/sharedlib/consts"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/fx"
)

func MonitoringFxOption() fx.Option {
	return fx.Options(
		fx.Provide(func(config configuration.Config) (monitoring2.Builder, error) {
			name, err := config.Get(consts.ServiceNameEnv).String()
			if err != nil {
				return nil, err
			}
			return bprometheus.Builder().SetNamespace(name), nil
		}),
		fx.Provide(monitoring.DefaultMonitor),
	)
}

func PrometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
