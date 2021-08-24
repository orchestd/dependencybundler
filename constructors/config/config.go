package config

import (
	"bitbucket.org/HeilaSystems/configurations/config"
	"bitbucket.org/HeilaSystems/configurations/config/confgetter/repos/cache"
	"bitbucket.org/HeilaSystems/dependencybundler/depBundler"
	cache2 "bitbucket.org/HeilaSystems/dependencybundler/interfaces/cache"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/configuration"
	"github.com/gin-contrib/cors"
	"os"
)

func DefaultConfiguration(getter cache2.CacheStorageGetter, builder config.Builder) configuration.Config {
	dockerName, isExist := os.LookupEnv(depBundler.DockerNameEnv)
	if !isExist {
		panic("missing DOCKER_NAME environment variable")
	}
	env, isExist := os.LookupEnv(depBundler.HeilaEnv)
	if !isExist {
		panic("missing HEILA_ENV environment variable")
	}
	repo := cache.NewCacheVariablesParamsResolver(dockerName, env, "1", getter)
	cfg, err := builder.SetEnv(env).SetServiceName(dockerName).SetRepo(repo).Build()
	if err != nil {
		panic(err)
	}
	localCfg := configuration.Config(cfg)
	return localCfg
}

func CorsConf() cors.Config {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "token")
	return corsConfig
}
