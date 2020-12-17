package config

import (
	"bitbucket.org/HeilaSystems/configurations/config"
	"bitbucket.org/HeilaSystems/configurations/config/confgetter/repos/cache"
	cache2 "bitbucket.org/HeilaSystems/dependencybundler/interfaces/cache"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/configuration"
	"os"
)

func DefaultConfiguration( getter cache2.CacheStorageGetter,builder config.Builder) configuration.Config {
	dockerName , isExist :=os.LookupEnv("DOCKER_NAME")
	if !isExist{
		panic("missing DOCKER_NAME environment variable")
	}
	env , isExist :=os.LookupEnv("HEILA_ENV")
	if !isExist{
		panic("missing HEILA_ENV environment variable")
	}
	repo := cache.NewCacheVariablesParamsResolver(dockerName,env,getter)
	cfg ,err:= builder.SetEnv(env).SetServiceName(dockerName).SetRepo(repo).Build()
	if err != nil {
		panic(err)
	}
	localCfg := configuration.Config(cfg)
	return localCfg
}

