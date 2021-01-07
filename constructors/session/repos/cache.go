package repos

import (
	"bitbucket.org/HeilaSystems/dependencybundler/depBundler"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/cache"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/configuration"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/session"
	cache2 "bitbucket.org/HeilaSystems/session/sessionresolver/repos/cache"
)

func DefaultCacheSessionRepo(config configuration.Config, getter cache.CacheStorageGetter) (session.SessionRepo, error) {
	collectionName, err := config.Get(depBundler.SessionCollection).String()
	if err != nil {
		return nil, err
	}
	sessionCacheRepo := cache2.NewSessionCacheRepo(getter ,collectionName,"1" )
	return sessionCacheRepo,nil
}