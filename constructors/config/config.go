package config

import (
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/config"
	"os"
)

func DefaultCacheStorageClient( builder config.Builder,repo config.ConfParamsResolver) config.Config {
	dockerName , isExist :=os.LookupEnv("DOCKER_NAME")
	if !isExist{
		panic("missing DOCKER_NAME environment variable")
	}
	env , isExist :=os.LookupEnv("HEILA_ENV")
	if !isExist{
		panic("missing HEILA_ENV environment variable")
	}

	cfg ,err:= builder.SetEnv(env).SetServiceName(dockerName).SetRepo(repo).Build()
	if err != nil {
		panic(err)
	}
	return cfg
}

