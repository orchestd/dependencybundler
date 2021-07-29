package depBundler

type DependencyBundlerConfiguration struct {
	Port string `json:"port"`
	ReadTimeOutMs string `json:"readTimeOutMs,omitempty"`
	WriteTimeOutMs string `json:"writeTimeOutMs,omitempty"`
	ContextHeaders []string `json:"contextHeaders,omitempty"`
	MonitorTags map[string]string `json:"monitorTags,omitempty"`
	SessionCollection string `json:"sessionCollection"`
	LogToFile bool `json:"logToFile"`
	MinimumSeverityLevel string `json:"minimumSeverityLevel"`
	LogToConsole bool `json:"logToConsole"`
	FileJsonFormat bool `json:"fileJsonFormat"`
	ConsoleJsonFormat bool `json:"consoleJsonFormat"`
	CompressLogs interface{} `json:"compressLogs"`
	DisableConsoleColor bool `json:"disableConsoleColor,omitempty"`
	DebugMode bool `json:"debugMode,omitempty"`

	HeilaEnv string `json:"HEILA_ENV"`
	DockerName string `json:"DOCKER_NAME"`
	ProjectId *string `json:"PROJECT_ID,omitempty"`
	SecretManager bool `json:"ENABLE_SECRET_MANAGER,omitempty"`
	SecretManagerVersion *string `json:"SECRET_MANAGER_VERSION,omitempty"`
	DbUsername *string `json:"DB_USERNAME,omitempty"`
	DbHost *string `json:"DB_HOST,omitempty"`

}

