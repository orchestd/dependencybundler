package depBundler

type DependencyBundlerConfiguration struct {
	Port string `json:"port"`
	ContextHeaders []string `json:"contextHeaders"`
	SessionCollection string `json:"sessionCollection"`
	LogToFile bool `json:"logToFile"`
	MinimumSeverityLevel string `json:"minimumSeverityLevel"`
	LogToConsole bool `json:"logToConsole"`
	FileJsonFormat bool `json:"fileJsonFormat"`
	ConsoleJsonFormat bool `json:"consoleJsonFormat"`
	CompressLogs interface{} `json:"compressLogs"`
	DisableConsoleColor bool `json:"disableConsoleColor"`

	DockerName string `json:"DOCKER_NAME"`
	ProjectId *string `json:"PROJECT_ID,omitempty"`
	SecretManager bool `json:"ENABLE_SECRET_MANAGER,omitempty"`
	SecretManagerVersion *string `json:"SECRET_MANAGER_VERSION,omitempty"`
}

