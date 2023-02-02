package config

import (
	"github.com/orchestd/configurations/config"
	"github.com/orchestd/dependencybundler/interfaces/configuration"
	"github.com/orchestd/sharedlib/consts"
	"os"
)

/*func DefaultConfiguration(getter cache2.CacheStorageGetter, builder config.Builder) configuration.Config {*/
func DefaultConfiguration(builder config.Builder) configuration.Config {
	serviceName, isExist := os.LookupEnv(consts.ServiceNameEnv)
	if !isExist {
		panic("missing " + consts.ServiceNameEnv + " environment variable")
	}
	env, isExist := os.LookupEnv(consts.HeilaEnv)
	if !isExist {
		panic("missing HEILA_ENV environment variable")
	}
	//repo := cache.NewCacheVariablesParamsResolver(dockerName, env, "1", getter)
	//cfg, err := builder.SetEnv(env).SetServiceName(dockerName).SetRepo(repo).Build()
	cfg, err := builder.SetEnv(env).SetServiceName(serviceName).Build()
	if err != nil {
		panic(err)
	}
	localCfg := configuration.Config(cfg)
	return localCfg
}
