package cache

import (
	"bitbucket.org/HeilaSystems/cacheStorage"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/cache"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/configuration"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/credentials"
	"context"
	"go.uber.org/fx"
)

func DefaultCacheStorageClient(lc fx.Lifecycle, credentials credentials.CredentialsGetter, config configuration.Config,
	cacheStorage cacheStorage.CacheStorage) (cache.CacheStorageGetter, cache.CacheStorageSetter) {
	creds := credentials.GetCredentials()
	dbName, err := config.Get("CACHE_DB_NAME").String()
	if err != nil {
		panic("env variable CACHE_DB_NAME must be defined")
	}
	lc.Append(fx.Hook{
		OnStart: func(c context.Context) error {
			return cacheStorage.Connect(c,creds.CacheConnectionString, dbName)
		},
		OnStop: func(c context.Context) error {
			return cacheStorage.Close(c)
		},
	})
	return cacheStorage.GetCacheStorageClient()
}
