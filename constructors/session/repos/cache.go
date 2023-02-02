package repos

import (
	"github.com/orchestd/dependencybundler/interfaces/cache"
	"github.com/orchestd/dependencybundler/interfaces/configuration"
	"github.com/orchestd/dependencybundler/interfaces/session"
	cache2 "github.com/orchestd/session/sessionresolver/repos/cache"
	"github.com/orchestd/sharedlib/consts"
)

func DefaultCacheSessionRepo(config configuration.Config, getter cache.CacheStorageGetterWrapper, setter cache.CacheStorageSetterWrapper) (session.SessionRepo, error) {
	collectionName, err := config.Get(consts.SessionCollection).String()
	if err != nil {
		return nil, err
	}
	sessionCacheRepo := cache2.NewSessionCacheRepo(getter, setter, collectionName, "1")
	return sessionCacheRepo, nil
}
