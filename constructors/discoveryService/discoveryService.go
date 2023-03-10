package discoveryService

import (
	transport "github.com/orchestd/dependencybundler/interfaces/transport"
	"github.com/orchestd/transport/discoveryService"
	"go.uber.org/fx"
)

type dspDeps struct {
	fx.In
	Lc fx.Lifecycle
	//Client     transport.HttpClient
	//Conf       configuration.Config
	//Logger     log.Logger
}

func DefaultDiscoveryServiceProvider(deps dspDeps) transport.DiscoveryServiceProvider {
	//TODO - true implementation would need to be here !
	//But httpServer is not a struct, so we manually push ds into in + httpClient
	var dsp discoveryService.DiscoveryServiceProvider
	return dsp

}
