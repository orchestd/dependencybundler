package bundler

import (
	"bitbucket.org/HeilaSystems/configurations/config"
	"bitbucket.org/HeilaSystems/configurations/config/confgetter"
	confCacheRepo "bitbucket.org/HeilaSystems/configurations/config/confgetter/repos/cache"
	configConstructor "bitbucket.org/HeilaSystems/dependencybundler/constructors/config"
	"go.uber.org/fx"
)

func ConfigFxOption(appconfstruct interface{}) fx.Option {
	return fx.Options(
		fx.Provide(confCacheRepo.NewCacheVariablesParamsResolver),
		fx.Provide(func() config.Builder{
			builder := confgetter.Builder().SetConfStruct(appconfstruct)
			return builder
		}),
		fx.Provide(configConstructor.DefaultConfiguration),
		fx.Provide(configConstructor.CorsConf),
	)
}
