package bundler

import (
	transportConstructor "bitbucket.org/HeilaSystems/dependencybundler/constructors/transport"
	"bitbucket.org/HeilaSystems/transport/client"
	httpClient "bitbucket.org/HeilaSystems/transport/client/http"
	"bitbucket.org/HeilaSystems/transport/server"
	"bitbucket.org/HeilaSystems/transport/server/http"
	"fmt"
	"go.uber.org/fx"
	"reflect"
)

func TransportFxOption(monolithConstructor ...interface{} ) fx.Option {
	var optionArr []fx.Option
	for _, v := range monolithConstructor {
		fmt.Println(reflect.TypeOf(v).String())
		optionArr = append(optionArr, fx.Provide(v))
	}
	userConstructors := fx.Options(optionArr...)
	return fx.Options(
		fx.Provide(func() server.HttpBuilder {
			builder := http.Builder()
			return builder
		}),
		fx.Provide(func() client.HTTPClientBuilder {
			builder := httpClient.HTTPClientBuilder()
			return builder
		}),
		fx.Provide(transportConstructor.DefaultTransport),
		userConstructors,
		)
}
