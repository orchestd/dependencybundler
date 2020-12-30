package bundler

import (
	"bitbucket.org/HeilaSystems/cacheStorage/mongodb"
	"bitbucket.org/HeilaSystems/dependencybundler/constructors/cache"
	"go.uber.org/fx"
)

func CacheFxOption() fx.Option {
	return fx.Options(
		fx.Provide(mongodb.NewMongoDbCacheStorage),
		fx.Provide(cache.DefaultCacheStorageClient),
	)
}
