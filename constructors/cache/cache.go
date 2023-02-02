package cache

import (
	"context"
	"github.com/orchestd/cacheStorage"
	"github.com/orchestd/dependencybundler/interfaces/cache"
	"github.com/orchestd/dependencybundler/interfaces/configuration"
	"github.com/orchestd/dependencybundler/interfaces/credentials"
	"go.uber.org/fx"
)

func DefaultCacheStorageClient(lc fx.Lifecycle, credentials credentials.CredentialsGetter, config configuration.Config,
	cacheStorage cacheStorage.CacheStorage) (cache.CacheStorageGetter, cache.CacheStorageSetter) {
	creds := credentials.GetCredentials()
	dbName, err := config.Get("CACHE_DB_NAME").String()
	if err != nil {
		panic("env variable CACHE_DB_NAME must be defined")
	}
	host, err := config.Get("CACHE_HOST").String()
	if err != nil {
		panic("env variable CACHE_HOST must be defined")
	}

	lc.Append(fx.Hook{
		OnStart: func(c context.Context) error {
			return cacheStorage.Connect(c, host, creds.CacheUserName, creds.CacheUserPw, dbName)
		},
		OnStop: func(c context.Context) error {
			return cacheStorage.Close(c)
		},
	})
	return cacheStorage.GetCacheStorageClient()
}
