package bundler

import (
	dspConstructor "github.com/orchestd/dependencybundler/constructors/discoveryService"
	transportConstructor "github.com/orchestd/dependencybundler/constructors/transport"
	clientMiddlewares "github.com/orchestd/dependencybundler/constructors/transport/middlewares/client"
	serverMiddlewares "github.com/orchestd/dependencybundler/constructors/transport/middlewares/server"
	"github.com/orchestd/dependencybundler/depBundler/middlewares/context"
	"github.com/orchestd/dependencybundler/depBundler/middlewares/metrics"
	"github.com/orchestd/dependencybundler/depBundler/middlewares/trace"
	"github.com/orchestd/dependencybundler/interfaces/transport"
	"github.com/orchestd/transport/client"
	httpClient "github.com/orchestd/transport/client/http"
	"github.com/orchestd/transport/server"
	"github.com/orchestd/transport/server/http"
	"go.uber.org/fx"
)

const ClientInterceptorsGroup = "clientInterceptors"
const ApiInterceptors = "apiInterceptors"
const RouterInterceptors = "routerInterceptors"
const ServerDebugInterceptors = "serverDebugInterceptors"
const SystemHandlers = "systemHandlers"

func TransportFxOption(monolithConstructor ...interface{}) fx.Option {
	var optionArr []fx.Option
	for _, v := range monolithConstructor {
		if o, ok := v.(fx.Option); ok {
			optionArr = append(optionArr, o)
		} else {
			optionArr = append(optionArr, fx.Provide(v))
		}
	}
	userConstructors := fx.Options(optionArr...)
	return fx.Options(

		//HTTP server bundler

		fx.Provide(dspConstructor.DefaultDiscoveryServiceProvider),
		fx.Provide(fx.Annotated{Group: ClientInterceptorsGroup, Target: clientMiddlewares.DefaultContextValuesToHeaders}),
		fx.Provide(fx.Annotated{Group: ClientInterceptorsGroup, Target: trace.TracerRESTClientInterceptor}),
		fx.Provide(fx.Annotated{Group: ClientInterceptorsGroup, Target: clientMiddlewares.DefaultTokenClientInterceptors}),
		fx.Provide(fx.Annotated{Group: ClientInterceptorsGroup, Target: clientMiddlewares.DefaultServiceNameToHeader}),

		fx.Provide(func() server.HttpBuilder {
			builder := http.Builder()
			return builder
		}),
		fx.Provide(transportConstructor.DefaultTransport),

		fx.Provide(fx.Annotated{Group: SystemHandlers, Target: transport.NewHttpHandler("GET", "metrics", PrometheusHandler())}),
		//HTTP client bundlerDefaultHeadersToContext
		fx.Provide(fx.Annotated{Group: SystemHandlers, Target: transport.NewHttpHandler("GET", "isAlive", http.IsAliveGinHandler)}),

		fx.Provide(fx.Annotated{Group: ApiInterceptors, Target: serverMiddlewares.DefaultHeadersToContext}),
		fx.Provide(fx.Annotated{Group: ApiInterceptors, Target: serverMiddlewares.DefaultBasicRequestId}),
		fx.Provide(fx.Annotated{Group: ApiInterceptors, Target: serverMiddlewares.DefaultJwtToken}),
		fx.Provide(fx.Annotated{Group: ApiInterceptors, Target: serverMiddlewares.DefaultLogHandlerMiddleware}),
		fx.Provide(fx.Annotated{Group: ApiInterceptors, Target: trace.HttpTracingUnaryServerInterceptor}),
		fx.Provide(fx.Annotated{Group: ApiInterceptors, Target: context.CallerToContext}),
		fx.Provide(fx.Annotated{Group: ApiInterceptors, Target: metrics.AverageRequestDurationMetric}),

		fx.Provide(func() client.HTTPClientBuilder {
			builder := httpClient.HTTPClientBuilder()
			return builder
		}),

		userConstructors,
	)
}
