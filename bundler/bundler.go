package bundler

import (
	"bitbucket.org/HeilaSystems/debug"
	"bitbucket.org/HeilaSystems/dependencybundler/constructors/session"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/cache"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/validations"
	"bitbucket.org/HeilaSystems/validations/cacheValidator"
	"context"
	"go.uber.org/fx"
	"log"
)

func CreateApplication(confStruct interface{}, HandlersFunc interface{}, monolithConstructor ...interface{}) {
	app := fx.New(
		CredentialsFxOption(),
		CacheFxOption(),
		ConfigFxOption(confStruct),
		LoggerFxOption(),
		TransportFxOption(monolithConstructor...),
		CacheTraceMiddlewareOption(),
		TracerFxOption(),
		SessionFxOption(),
		fx.Provide(session.NewContextData),
		DebugFxOption(),
		ValidationsFxOption(),
		MonitoringFxOption(),
		fx.Invoke(HandlersFunc, debug.InitHandlers),
		fx.Provide(func(storage cache.CacheStorageGetter) validations.ValidatorRunner {
			return cacheValidator.NewCacheValidatorRunner(storage)
		}),
	)

	c := context.Background()
	if err := app.Start(c); err != nil {
		log.Fatal(err)
	}
}
