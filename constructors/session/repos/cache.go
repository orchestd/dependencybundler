package repos

import (
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/cache"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/configuration"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/session"
	cache2 "bitbucket.org/HeilaSystems/session/sessionresolver/repos/cache"
	"bitbucket.org/HeilaSystems/sharedlib/consts"
)

func DefaultCacheSessionRepo(config configuration.Config, getter cache.CacheStorageGetterWrapper, setter cache.CacheStorageSetterWrapper) (session.SessionRepo, error) {
	collectionName, err := config.Get(consts.SessionCollection).String()
	if err != nil {
		return nil, err
	}
	sessionCacheRepo := cache2.NewSessionCacheRepo(getter ,setter,collectionName,"1" )
	return sessionCacheRepo,nil
}
