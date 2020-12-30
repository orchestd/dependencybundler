package cache

import (
	"bitbucket.org/HeilaSystems/cacheStorage"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/cache"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/credentials"
	"context"
	"go.uber.org/fx"
)

func DefaultCacheStorageClient(lc fx.Lifecycle, credentials credentials.CredentialsGetter, cacheStorage cacheStorage.CacheStorage) (cache.CacheStorageGetter, cache.CacheStorageSetter) {
	creds := credentials.GetCredentials()
	lc.Append(fx.Hook{
		OnStart: func(c context.Context) error {
			return cacheStorage.Connect(c, creds.DbUsername, creds.DbPassword, creds.DbHost, creds.DbName)
		},
		OnStop: func(c context.Context) error {
			return cacheStorage.Close(c)
		},
	})
	return cacheStorage.GetCacheStorageClient()
}
