package transport

import (
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/configuration"
	transportConstructor "bitbucket.org/HeilaSystems/dependencybundler/interfaces/transport"
	"bitbucket.org/HeilaSystems/transport/client"
	"bitbucket.org/HeilaSystems/transport/server"
	"go.uber.org/fx"
)

func DefaultTransport(lc fx.Lifecycle , builder server.HttpBuilder,clientBuilder client.HTTPClientBuilder,conf configuration.Config)  (transportConstructor.IRouter,transportConstructor.HttpClient) {
	if confPort,err := conf.Get("port").String();err == nil {
		builder = builder.SetPort(confPort)
	}

	if confReadTimeout,err := conf.Get("readTimeOut").Duration();err == nil {
		builder = builder.SetReadTimeout(confReadTimeout)
	}

	if confWriteTimeout,err := conf.Get("writeTimeOut").Duration();err == nil {
		builder = builder.SetWriteTimeout(confWriteTimeout)
	}

	return builder.Build(lc),clientBuilder.Build()
}
