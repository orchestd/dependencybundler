package bundler

import (
	"bitbucket.org/HeilaSystems/debug"
	"bitbucket.org/HeilaSystems/dependencybundler/constructors/session"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/cache"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/configuration"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/validations"
	"bitbucket.org/HeilaSystems/validations/cacheValidator"
	"go.uber.org/fx"
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
		fx.Provide(func(storage cache.CacheStorageGetter, conf configuration.Config) validations.ValidatorRunner {
			return cacheValidator.NewCacheValidatorRunner(storage, conf)
		}),
		fx.Invoke(HandlersFunc, debug.InitHandlers),
	)

	app.Run()
}
