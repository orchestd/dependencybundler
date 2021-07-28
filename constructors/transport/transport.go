package transport

import (
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
	Lc                        fx.Lifecycle
	ServerBuilder             server.HttpBuilder
	ClientBuilder             client.HTTPClientBuilder
	Conf                      configuration.Config
	Logger                    log.Logger
	ClientInterceptors        []client.HTTPClientInterceptor `group:"clientInterceptors"`
	ServerContextInterceptors []gin.HandlerFunc              `group:"serverContextInterceptors"`
	ServerInterceptors        []gin.HandlerFunc              `group:"serverInterceptors"`
	SystemHandlers        []server.IHandler              `group:"systemHandlers"`
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
		deps.ServerBuilder = deps.ServerBuilder.SetReadTimeout(confReadTimeout*time.Millisecond)
	}

	if confWriteTimeout, err := deps.Conf.Get("writeTimeOutMs").Duration(); err != nil {
		deps.Logger.WithError(err).Debug(context.Background(), "Cannot get writeTimeOutMs from configuration, setting write time out to 30 seconds")
	} else {
		deps.ServerBuilder = deps.ServerBuilder.SetWriteTimeout(confWriteTimeout*time.Millisecond)
	}
	deps.ServerBuilder = deps.ServerBuilder.SetLogger(deps.Logger)
	if len(deps.ClientInterceptors) > 0 {
		deps.ClientBuilder = deps.ClientBuilder.AddInterceptors(deps.ClientInterceptors...)
	}
	deps.ClientBuilder = deps.ClientBuilder.SetConfig(deps.Conf)

	if len(deps.ServerContextInterceptors) > 0 {
		deps.ServerBuilder = deps.ServerBuilder.AddContextInterceptors(deps.ServerContextInterceptors...)
	}

	if len(deps.ServerInterceptors) > 0 {
		deps.ServerBuilder = deps.ServerBuilder.AddInterceptors(deps.ServerInterceptors...)
	}

	if len(deps.SystemHandlers) > 0 {
		deps.ServerBuilder = deps.ServerBuilder.AddSystemHandlers(deps.SystemHandlers...)
	}

	client, err := deps.ClientBuilder.Build()
	if err != nil {
		panic(err)
	}
	return deps.ServerBuilder.Build(deps.Lc), client
}
