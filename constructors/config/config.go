package config

import (
	"bitbucket.org/HeilaSystems/configurations/config"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/configuration"
	"bitbucket.org/HeilaSystems/sharedlib/consts"
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
