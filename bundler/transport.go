package bundler

import (
	transportConstructor "bitbucket.org/HeilaSystems/dependencybundler/constructors/transport"
	clientMiddlewares "bitbucket.org/HeilaSystems/dependencybundler/constructors/transport/middlewares/client"
	serverMiddlewares "bitbucket.org/HeilaSystems/dependencybundler/constructors/transport/middlewares/server"
	"bitbucket.org/HeilaSystems/dependencybundler/depBundler/middlewares/context"
	"bitbucket.org/HeilaSystems/dependencybundler/depBundler/middlewares/metrics"
	"bitbucket.org/HeilaSystems/dependencybundler/depBundler/middlewares/trace"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/transport"
	"bitbucket.org/HeilaSystems/transport/client"
	httpClient "bitbucket.org/HeilaSystems/transport/client/http"
	"bitbucket.org/HeilaSystems/transport/server"
	"bitbucket.org/HeilaSystems/transport/server/http"
	"go.uber.org/fx"
)

const ClientInterceptorsGroup = "clientInterceptors"
const ServerInterceptors = "serverInterceptors"
const ServerDebugInterceptors = "serverDebugInterceptors"
const SystemHandlers = "systemHandlers"

func TransportFxOption(monolithConstructor ...interface{}) fx.Option {
	var optionArr []fx.Option
	for _, v := range monolithConstructor {
		optionArr = append(optionArr, fx.Provide(v))
	}
	userConstructors := fx.Options(optionArr...)
	return fx.Options(

		//HTTP server bundler

		fx.Provide(fx.Annotated{Group: ClientInterceptorsGroup, Target: clientMiddlewares.DefaultContextValuesToHeaders}),
		fx.Provide(fx.Annotated{Group: ClientInterceptorsGroup, Target: trace.TracerRESTClientInterceptor}),
		fx.Provide(fx.Annotated{Group: ClientInterceptorsGroup, Target: clientMiddlewares.DefaultTokenClientInterceptors}),
		fx.Provide(fx.Annotated{Group: ClientInterceptorsGroup, Target: clientMiddlewares.DefaultServiceNameToHeader}),

		fx.Provide(func() server.HttpBuilder {
			builder := http.Builder()
			return builder
		}),
		fx.Provide(transportConstructor.DefaultTransport),

		fx.Provide(fx.Annotated{Group: SystemHandlers, Target: transport.NewHttpHandler("GET" , "metrics" , PrometheusHandler())}),
		//HTTP client bundlerDefaultHeadersToContext
		fx.Provide(fx.Annotated{Group: ServerInterceptors, Target: serverMiddlewares.DefaultHeadersToContext}),
		fx.Provide(fx.Annotated{Group: ServerInterceptors, Target: serverMiddlewares.DefaultBasicRequestId}),
		fx.Provide(fx.Annotated{Group: ServerInterceptors, Target: serverMiddlewares.DefaultJwtToken}),
		fx.Provide(fx.Annotated{Group: ServerInterceptors, Target: serverMiddlewares.DefaultLogHandlerMiddleware}),
		fx.Provide(fx.Annotated{Group: ServerInterceptors, Target: trace.HttpTracingUnaryServerInterceptor}),
		fx.Provide(fx.Annotated{Group: ServerInterceptors, Target: context.CallerToContext}),
		fx.Provide(fx.Annotated{Group: ServerInterceptors, Target: metrics.AverageRequestDurationMetric}),



		fx.Provide(func() client.HTTPClientBuilder {
			builder := httpClient.HTTPClientBuilder()
			return builder
		}),

		userConstructors,
	)
}
