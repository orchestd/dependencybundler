package bundler

import (
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
		SessionFxOption(),
		TracerFxOption(),
		fx.Invoke(HandlersFunc),
		)

	c := context.Background()
	if err := app.Start(c); err != nil {
		log.Fatal(err)
	}
}

