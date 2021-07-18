package bundler

import (
	"bitbucket.org/HeilaSystems/debug"
	"context"
	"go.uber.org/fx"
	"log"
)

func CreateApplication(confStruct interface{}, HandlersFunc interface{} , monolithConstructor ...interface{})  {
	app :=  fx.New(
		CredentialsFxOption(),
		CacheFxOption(),
		ConfigFxOption(confStruct),
		LoggerFxOption(),
		TransportFxOption(monolithConstructor...),
		CacheTraceMiddlewareOption(),
		TracerFxOption(),
		SessionFxOption(),
		DebugFxOption(),
		ValidationsFxOption(),
		fx.Invoke(HandlersFunc,debug.InitHandlers),
		)

	c := context.Background()
	if err := app.Start(c); err != nil {
		log.Fatal(err)
	}
}

