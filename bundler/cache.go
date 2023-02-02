package bundler

import (
	"github.com/orchestd/cacheStorage/mongodb"
	"github.com/orchestd/dependencybundler/constructors/cache"
	"github.com/orchestd/dependencybundler/constructors/cache/middlewares"
	"github.com/orchestd/dependencybundler/constructors/cache/middlewares/trace"
	"github.com/orchestd/sharedlib/consts"
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
		fx.Provide(fx.Annotated{Group: consts.CacheStorageGetterMiddlewares, Target: trace.DefaultCacheGetterTraceMiddleware}),
		fx.Provide(fx.Annotated{Group: consts.CacheStorageSetterMiddlewares, Target: trace.DefaultCacheSetterMiddleware}),
		fx.Provide(middlewares.DefaultCacheWrapper),
	)
}
