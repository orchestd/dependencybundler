package middlewares

import (
	"bitbucket.org/HeilaSystems/cacheStorage"
	"bitbucket.org/HeilaSystems/cacheStorage/mongodb/middlewares"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/cache"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/configuration"
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

func DefaultCacheWrapper(deps CacheWrapperDeps) (cache.CacheStroageGetterWrapper, cache.CacheStroageSetterWrapper) {
	return middlewares.CreateCacheGetterWrapper(deps.CacheStorageGetter,deps.GetterMiddlewares...), middlewares.CreateCacheSetterWrapper(deps.CacheStorageSetter, deps.SetterMiddlewares...)
}
