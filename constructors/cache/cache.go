package cache

import (
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/cache"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/credentials"
	"context"
	"go.uber.org/fx"
)


func DefaultCacheStorageClient(lc fx.Lifecycle, credentials credentials.CredentialsGetter,builder cache.CacheStorageBuilder) (cache.CacheStorageGetter, cache.CacheStorageSetter) {
	creds := credentials.GetCredentials()
	cs := builder.SetDatabaseName(creds.DbName).SetHost(creds.DbHost).SetUsername(creds.DbUsername).SetPassword(creds.DbPassword).Build()

	lc.Append(fx.Hook{
		OnStart: func(c context.Context) error {
			return cs.Connect(c)
		},
		OnStop: func(c context.Context) error {
			return cs.Close(c)
		},
	})
	return cs.GetCacheStorageClient()
}

