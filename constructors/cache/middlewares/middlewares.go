package middlewares

import (
	"github.com/orchestd/cacheStorage"
	"github.com/orchestd/cacheStorage/mongodb/middlewares"
	"github.com/orchestd/dependencybundler/interfaces/cache"
	"github.com/orchestd/dependencybundler/interfaces/configuration"
	"go.uber.org/fx"
)

type CacheWrapperDeps struct {
	fx.In
	Config             configuration.Config
	CacheStorageGetter cache.CacheStorageGetter
	CacheStorageSetter cache.CacheStorageSetter
	GetterMiddlewares  []cacheStorage.CacheStorageGetterMiddleware `group:"cacheStorageGetterMiddlewares"`
	SetterMiddlewares  []cacheStorage.CacheStorageSetterMiddleware `group:"cacheStorageSetterMiddlewares"`
}

func DefaultCacheWrapper(deps CacheWrapperDeps) (cache.CacheStorageGetterWrapper, cache.CacheStorageSetterWrapper) {
	return middlewares.CreateCacheGetterWrapper(deps.CacheStorageGetter, deps.GetterMiddlewares...), middlewares.CreateCacheSetterWrapper(deps.CacheStorageSetter, deps.SetterMiddlewares...)
}
