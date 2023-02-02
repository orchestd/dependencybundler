package depBundler

import (
	cacheStorageConfiguration "github.com/orchestd/cacheStorage/configuration"
	"github.com/orchestd/configurations/config/envConfiguration"
	"github.com/orchestd/configurations/credentials/credentialsConfiguration"
	logConfiguration "github.com/orchestd/log/configuration"
	monitoringConfiguration "github.com/orchestd/monitoring/configuration"
	sessionConfiguration "github.com/orchestd/session/configuration"
	transportConfiguration "github.com/orchestd/transport/configuration"
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
