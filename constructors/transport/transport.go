package transport

import (
	"context"
	"github.com/gin-gonic/gin"
	discoveryServiceProviders "github.com/orchestd/dependencybundler/constructors/discoveryService/providers"
	"github.com/orchestd/dependencybundler/interfaces/configuration"
	"github.com/orchestd/dependencybundler/interfaces/credentials"
	"github.com/orchestd/dependencybundler/interfaces/log"
	"github.com/orchestd/dependencybundler/interfaces/transport"
	transportConstructor "github.com/orchestd/dependencybundler/interfaces/transport"
	"github.com/orchestd/sharedlib/consts"
	"github.com/orchestd/transport/client"
	"github.com/orchestd/transport/server"
	"go.uber.org/fx"
	"net/http"
	"net/http/pprof"
	"time"
)

type transportDeps struct {
	fx.In
	Lc                       fx.Lifecycle
	ServerBuilder            server.HttpBuilder
	ClientBuilder            client.HTTPClientBuilder
	Conf                     configuration.Config
	Cred                     credentials.CredentialsGetter
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

	var enabledPprofHandlersUntil time.Time

	deps.SystemHandlers = append(deps.SystemHandlers, transport.NewHttpHandler("POST", "/enable_pprof/", func(c *gin.Context) {
		type EnablePprof struct {
			EnableUntil time.Time
			Pw          string
		}

		enablePprofObj := EnablePprof{}
		err := c.BindJSON(&enablePprofObj)
		if err != nil {
			c.AbortWithError(500, err)
		}
		if deps.Cred.GetCredentials().EnablePprofPw != enablePprofObj.Pw {
			c.AbortWithStatus(501)
		}
		enabledPprofHandlersUntil = enablePprofObj.EnableUntil
		c.AbortWithStatus(200)
	})())
	checkCanRunPprofHandler := func(c *gin.Context, f func(w http.ResponseWriter, r *http.Request)) {
		if enabledPprofHandlersUntil.After(time.Now()) {
			f(c.Writer, c.Request)
		} else {
			c.AbortWithStatus(404)
		}
	}
	deps.SystemHandlers = append(deps.SystemHandlers, transport.NewHttpHandler("GET", "/pprof/", func(c *gin.Context) {
		checkCanRunPprofHandler(c, pprof.Index)
	})())
	deps.SystemHandlers = append(deps.SystemHandlers, transport.NewHttpHandler("GET", "/pprof/cmdline", func(c *gin.Context) {
		checkCanRunPprofHandler(c, pprof.Cmdline)
	})())
	deps.SystemHandlers = append(deps.SystemHandlers, transport.NewHttpHandler("GET", "/pprof/profile", func(c *gin.Context) {
		checkCanRunPprofHandler(c, pprof.Profile)
	})())
	deps.SystemHandlers = append(deps.SystemHandlers, transport.NewHttpHandler("POST", "/pprof/symbol", func(c *gin.Context) {
		checkCanRunPprofHandler(c, pprof.Symbol)
	})())
	deps.SystemHandlers = append(deps.SystemHandlers, transport.NewHttpHandler("GET", "/pprof/symbol", func(c *gin.Context) {
		checkCanRunPprofHandler(c, pprof.Symbol)
	})())
	deps.SystemHandlers = append(deps.SystemHandlers, transport.NewHttpHandler("GET", "/pprof/trace", func(c *gin.Context) {
		checkCanRunPprofHandler(c, pprof.Trace)
	})())
	deps.SystemHandlers = append(deps.SystemHandlers, transport.NewHttpHandler("GET", "/pprof/allocs", func(c *gin.Context) {
		checkCanRunPprofHandler(c, pprof.Handler("allocs").ServeHTTP)
	})())
	deps.SystemHandlers = append(deps.SystemHandlers, transport.NewHttpHandler("GET", "/pprof/block", func(c *gin.Context) {
		checkCanRunPprofHandler(c, pprof.Handler("block").ServeHTTP)
	})())
	deps.SystemHandlers = append(deps.SystemHandlers, transport.NewHttpHandler("GET", "/pprof/goroutine", func(c *gin.Context) {
		checkCanRunPprofHandler(c, pprof.Handler("goroutine").ServeHTTP)
	})())
	deps.SystemHandlers = append(deps.SystemHandlers, transport.NewHttpHandler("GET", "/pprof/heap", func(c *gin.Context) {
		checkCanRunPprofHandler(c, pprof.Handler("heap").ServeHTTP)
	})())
	deps.SystemHandlers = append(deps.SystemHandlers, transport.NewHttpHandler("GET", "/pprof/mutex", func(c *gin.Context) {
		checkCanRunPprofHandler(c, pprof.Handler("mutex").ServeHTTP)
	})())
	deps.SystemHandlers = append(deps.SystemHandlers, transport.NewHttpHandler("GET", "/pprof/threadcreate", func(c *gin.Context) {
		checkCanRunPprofHandler(c, pprof.Handler("threadcreate").ServeHTTP)
	})())

	if deps.Conf.Get("assetRoots").IsSet() {
		if err := deps.Conf.Get("assetRoots").Unmarshal(&assetRoots); err != nil {
			deps.Logger.WithError(err).Info(context.Background(), "Cannot Unmarshal assetRoots from configuration")
			panic("Cannot read assetRoots from configuration")
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
	}
	deps.ServerBuilder = deps.ServerBuilder.SetStatics(staticHandlers)

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
	if dspType, err := deps.Conf.Get(consts.DiscoveryServiceProvider).String(); err != nil {
		deps.Logger.WithError(err).Debug(context.Background(), "Cannot get discoveryServiceProvider from configurations, setting to none")
	} else {
		switch dspType {
		case "none":
			dsp = discoveryServiceProviders.NewNoDSP()
		case "conf":
			dsp = discoveryServiceProviders.NewConfDSP()
		case "simple":
			dsp = discoveryServiceProviders.NewSimpleHttpDSP(client, deps.Conf, deps.Logger)
		case "templated":
			dsp = discoveryServiceProviders.NewTemplatedDSP(deps.Conf, deps.Logger)
		}
	}
	//ugly
	deps.ServerBuilder = deps.ServerBuilder.SetDiscoveryServiceProvider(dsp)
	deps.DiscoveryServiceProvider = dsp
	client.SetDiscoveryServiceProvider(dsp)
	//end ugly

	return deps.ServerBuilder.Build(deps.Lc), client
}
