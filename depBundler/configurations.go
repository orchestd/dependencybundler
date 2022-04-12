package depBundler

import (
	cacheStorageConfiguration "bitbucket.org/HeilaSystems/cacheStorage/configuration"
	"bitbucket.org/HeilaSystems/configurations/config/envConfiguration"
	"bitbucket.org/HeilaSystems/configurations/credentials/credentialsConfiguration"
	logConfiguration "bitbucket.org/HeilaSystems/log/configuration"
	monitoringConfiguration "bitbucket.org/HeilaSystems/monitoring/configuration"
	sessionConfiguration "bitbucket.org/HeilaSystems/session/configuration"
	transportConfiguration "bitbucket.org/HeilaSystems/transport/configuration"
)

type DependencyBundlerConfiguration struct {
	credentialsConfiguration.CredentialsConfiguration
	envConfiguration.EnvConfiguration
	cacheStorageConfiguration.CacheStorageConfiguration
	sessionConfiguration.SessionConfiguration
	monitoringConfiguration.MonitoringConfiguration
	logConfiguration.LogConfiguration
	transportConfiguration.TransportConfiguration
}
