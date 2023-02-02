package bundler

import (
	"github.com/orchestd/configurations/config"
	"github.com/orchestd/configurations/config/confgetter"
	configConstructor "github.com/orchestd/dependencybundler/constructors/config"
	"go.uber.org/fx"
)

func ConfigFxOption(appconfstruct interface{}) fx.Option {
	return fx.Options(
		/*fx.Provide(confCacheRepo.NewCacheVariablesParamsResolver),*/
		fx.Provide(func() config.Builder {
			builder := confgetter.Builder().SetConfStruct(appconfstruct)
			return builder
		}),
		fx.Provide(configConstructor.DefaultConfiguration),
	)
}
