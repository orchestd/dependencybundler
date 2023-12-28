package bundler

import (
	"github.com/orchestd/debug"
	"github.com/orchestd/dependencybundler/constructors/session"
	"github.com/orchestd/dependencybundler/interfaces/cache"
	"github.com/orchestd/dependencybundler/interfaces/configuration"
	"github.com/orchestd/dependencybundler/interfaces/validations"
	"github.com/orchestd/tokenauth"
	"github.com/orchestd/validations/cacheValidator"
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
		fx.Provide(tokenauth.NewJwtToken),
	)

	app.Run()
}
