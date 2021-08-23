package depBundler

import "bitbucket.org/HeilaSystems/session/sessionresolver"

const (
	DockerNameEnv = "DOCKER_NAME"
	HeilaEnv = "HEILA_ENV"
	DbHostEnv = "DB_HOST"
	DockerName = "dockerName"
	DbUsernameEnv = "DB_USERNAME"
	LogToFile  = "logToFile"
	LogToConsole = "logToConsole"
	FileJsonFormat = "fileJsonFormat"
	ConsoleJsonFormat = "consoleJsonFormat"
	CompressLogs = "compressLogs"
	SessionCollection = "sessionCollection"
	DisableConsoleColor = "disableConsoleColor"
	MinimumSeverityLevel = "minimumSeverityLevel"
 	FxGroupLoggerContextExtractors = "loggerContextExtractors"
	CacheStorageGetterMiddlewares = "cacheStorageGetterMiddlewares"
	CacheStorageSetterMiddlewares = "cacheStorageSetterMiddlewares"
	MonitorTags = "monitorTags"
	SessionTimeLayoutYYYYMMDD_HHMMSS = sessionresolver.TimeLayoutYYYYMMDD_HHMMSS
)
