package transport

import (
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/configuration"
	transportConstructor "bitbucket.org/HeilaSystems/dependencybundler/interfaces/transport"
	"bitbucket.org/HeilaSystems/transport/client"
	"bitbucket.org/HeilaSystems/transport/server"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type transportDeps struct {
	fx.In
	Lc                 fx.Lifecycle
	ServerBuilder      server.HttpBuilder
	ClientBuilder      client.HTTPClientBuilder
	Conf               configuration.Config
	ClientInterceptors []client.HTTPClientInterceptor`group:"clientInterceptors"`
	ServerInterceptors []gin.HandlerFunc `group:"serverInterceptors"`

}
func DefaultTransport(deps transportDeps)  (transportConstructor.IRouter,transportConstructor.HttpClient) {
	if confPort,err := deps.Conf.Get("port").String();err == nil {
		deps.ServerBuilder = deps.ServerBuilder.SetPort(confPort)
	}

	if confReadTimeout,err := deps.Conf.Get("readTimeOut").Duration();err == nil {
		deps.ServerBuilder = deps.ServerBuilder.SetReadTimeout(confReadTimeout)
	}

	if confWriteTimeout,err := deps.Conf.Get("writeTimeOut").Duration();err == nil {
		deps.ServerBuilder = deps.ServerBuilder.SetWriteTimeout(confWriteTimeout)
	}

	if len(deps.ClientInterceptors) > 0 {
		deps.ClientBuilder = deps.ClientBuilder.AddInterceptors(deps.ClientInterceptors...)
	}

	if len(deps.ServerInterceptors) > 0 {
		deps.ServerBuilder = deps.ServerBuilder.AddInterceptors(deps.ServerInterceptors...)
	}

	return deps.ServerBuilder.Build(deps.Lc),deps.ClientBuilder.Build()
}
