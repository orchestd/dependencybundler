package bundler

import (
	"bitbucket.org/HeilaSystems/cacheStorage"
	"bitbucket.org/HeilaSystems/cacheStorage/mock"
	"bitbucket.org/HeilaSystems/dependencybundler/constructors/cache"
	"go.uber.org/fx"
)

func CacheFxOption() fx.Option {
	return fx.Options(
		fx.Provide(func() cacheStorage.CacheStorageBuilder{
			return mock.Builder()
		}),
		fx.Provide(cache.DefaultCacheStorageClient),
	)
}
