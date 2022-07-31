package trace

import (
	"bitbucket.org/HeilaSystems/cacheStorage"
	"bitbucket.org/HeilaSystems/cacheStorage/mongodb/middlewares/trace"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/configuration"
	"bitbucket.org/HeilaSystems/sharedlib/consts"
	"github.com/opentracing/opentracing-go"
	"strings"
)


func DefaultCacheGetterTraceMiddleware( tracer opentracing.Tracer,config configuration.Config)cacheStorage.CacheStorageGetterMiddleware {
	wrapperConf := getWrapperConf(config)
	return trace.NewMongoCacheStorageGetterWrapper(tracer,wrapperConf)
}
func DefaultCacheSetterMiddleware( tracer opentracing.Tracer,config configuration.Config)cacheStorage.CacheStorageSetterMiddleware {
	wrapperConf := getWrapperConf(config)
	return trace.NewMongoCacheStorageSetterWrapper(tracer,wrapperConf)
}

func getWrapperConf(config configuration.Config)trace.CacheWrapperConfiguration {
	var wrapperConf trace.CacheWrapperConfiguration
	if 	dockerName,err := config.Get(consts.ServiceNameEnv).String();err == nil {
		wrapperConf.ServiceName=dockerName
	}
	if dbHost ,err := config.Get(consts.DbHostEnv).String();err == nil {
		if dbHostArr:=strings.Split(dbHost,"@"); len(dbHostArr) > 1 {
			dbHost = dbHostArr[1]
		}
		wrapperConf.DbHost=dbHost
	}

	if dbUser,err := config.Get(consts.DbUsernameEnv).String();err == nil {
		wrapperConf.DbUser=dbUser
	}
	return wrapperConf
}

