package bundler

import (
	"bitbucket.org/HeilaSystems/cacheStorage/mongodb"
	"bitbucket.org/HeilaSystems/dependencybundler/constructors/cache"
	"bitbucket.org/HeilaSystems/dependencybundler/constructors/cache/middlewares"
	"bitbucket.org/HeilaSystems/dependencybundler/constructors/cache/middlewares/trace"
	"bitbucket.org/HeilaSystems/dependencybundler/depBundler"
	"go.uber.org/fx"
)

func CacheFxOption() fx.Option {
	return fx.Options(
		fx.Provide(mongodb.NewMongoDbCacheStorage),
		fx.Provide(cache.DefaultCacheStorageClient),
	)
}

func CacheTraceMiddlewareOption() fx.Option {
	return fx.Options(
		fx.Provide(fx.Annotated{Group: depBundler.CacheStorageGetterMiddlewares, Target: trace.DefaultCacheGetterTraceMiddleware}),
		fx.Provide(fx.Annotated{Group: depBundler.CacheStorageSetterMiddlewares, Target: trace.DefaultCacheSetterMiddleware}),
		fx.Provide(middlewares.DefaultCacheWrapper),
	)
}
