package transport

import (
	discoveryServiceProviders "bitbucket.org/HeilaSystems/dependencybundler/constructors/discoveryService/providers"
	"bitbucket.org/HeilaSystems/dependencybundler/depBundler"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/configuration"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/log"
	transportConstructor "bitbucket.org/HeilaSystems/dependencybundler/interfaces/transport"
	"bitbucket.org/HeilaSystems/transport/client"
	"bitbucket.org/HeilaSystems/transport/server"
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"time"
)

type transportDeps struct {
	fx.In
	Lc                       fx.Lifecycle
	ServerBuilder            server.HttpBuilder
	ClientBuilder            client.HTTPClientBuilder
	Conf                     configuration.Config
	Logger                   log.Logger
	ClientInterceptors       []client.HTTPClientInterceptor `group:"clientInterceptors"`
	ServerDebugInterceptors  []gin.HandlerFunc              `group:"serverDebugInterceptors"`
	ApiInterceptors          []gin.HandlerFunc              `group:"apiInterceptors"`
	RouterInterceptors       []gin.HandlerFunc              `group:"routerInterceptors"`
	SystemHandlers           []server.IHandler              `group:"systemHandlers"`
	DiscoveryServiceProvider transportConstructor.DiscoveryServiceProvider
}

type AssetRoot struct {
	UrlPath     string
	FolderPath  string
	AllowOnProd bool
}

func DefaultTransport(deps transportDeps) (transportConstructor.IRouter, transportConstructor.HttpClient) {
	if confPort, err := deps.Conf.Get("port").String(); err != nil {
		deps.Logger.WithError(err).Info(context.Background(), "Cannot get port from configuration, setting port to 8080")
	} else {
		deps.ServerBuilder = deps.ServerBuilder.SetPort(confPort)
	}

	if confReadTimeout, err := deps.Conf.Get("readTimeOutMs").Duration(); err != nil {
		deps.Logger.WithError(err).Debug(context.Background(), "Cannot get readTimeOutMs from configuration, setting read time out to 30 seconds")
	} else {
		deps.ServerBuilder = deps.ServerBuilder.SetReadTimeout(confReadTimeout * time.Millisecond)
	}

	if confWriteTimeout, err := deps.Conf.Get("writeTimeOutMs").Duration(); err != nil {
		deps.Logger.WithError(err).Debug(context.Background(), "Cannot get writeTimeOutMs from configuration, setting write time out to 30 seconds")
	} else {
		deps.ServerBuilder = deps.ServerBuilder.SetWriteTimeout(confWriteTimeout * time.Millisecond)
	}
	deps.ServerBuilder = deps.ServerBuilder.SetLogger(deps.Logger)
	if len(deps.ClientInterceptors) > 0 {
		deps.ClientBuilder = deps.ClientBuilder.AddInterceptors(deps.ClientInterceptors...)
	}
	deps.ClientBuilder = deps.ClientBuilder.SetConfig(deps.Conf)

	var staticHandlers = make(map[string]string)

	var assetRoots []AssetRoot

	if deps.Conf.Get("assetRoots").IsSet() {
		if err := deps.Conf.Get("assetRoots").Unmarshal(&assetRoots); err != nil {
			deps.Logger.WithError(err).Info(context.Background(), "Cannot Unmarshal assetRoots from configuration")
			panic("Cannot read statics from configuration")
		} else {
			for _, a := range assetRoots {
				if a.AllowOnProd {
					staticHandlers[a.UrlPath] = a.FolderPath
				}
			}
		}
	}

	if debug, err := deps.Conf.Get("debugMode").Bool(); err != nil {
		deps.Logger.WithError(err).Debug(context.Background(), "Cannot get debug mode from configurations, setting mode to false")
	} else if debug {
		if len(deps.ServerDebugInterceptors) > 0 {
			deps.ServerBuilder = deps.ServerBuilder.AddApiInterceptors(deps.ServerDebugInterceptors...)
		}
		for _, a := range assetRoots {
			if !a.AllowOnProd {
				staticHandlers[a.UrlPath] = a.FolderPath
			}
		}
		deps.ServerBuilder = deps.ServerBuilder.SetStatics(staticHandlers)
	}

	if len(deps.ApiInterceptors) > 0 {
		deps.ServerBuilder = deps.ServerBuilder.AddApiInterceptors(deps.ApiInterceptors...)
	}

	if len(deps.RouterInterceptors) > 0 {
		deps.ServerBuilder = deps.ServerBuilder.AddRouterInterceptors(deps.RouterInterceptors...)
	}

	if len(deps.SystemHandlers) > 0 {
		deps.ServerBuilder = deps.ServerBuilder.AddSystemHandlers(deps.SystemHandlers...)
	}
	client, err := deps.ClientBuilder.Build()

	if err != nil {
		panic(err)
	}

	var dsp transportConstructor.DiscoveryServiceProvider
	if dspType, err := deps.Conf.Get(depBundler.DiscoveryServiceProvider).String(); err != nil {
		deps.Logger.WithError(err).Debug(context.Background(), "Cannot get discoveryServiceProvider from configurations, setting to none")
	} else {
		switch dspType {
		case "none":
			dsp = discoveryServiceProviders.NewNoDSP()
		case "conf":
			dsp = discoveryServiceProviders.NewConfDSP()
		case "simple":
			dsp = discoveryServiceProviders.NewSimpleHttpDSP(client, deps.Conf, deps.Logger)
		}
	}
	//ugly
	deps.ServerBuilder = deps.ServerBuilder.SetDiscoveryServiceProvider(dsp)
	deps.DiscoveryServiceProvider = dsp
	client.SetDiscoveryServiceProvider(dsp)
	//end ugly

	return deps.ServerBuilder.Build(deps.Lc), client
}
