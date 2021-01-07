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
	CompressLogs bool `json:"compressLogs"`
	DockerName string `json:"DOCKER_NAME"`
}

