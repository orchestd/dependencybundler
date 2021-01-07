package depBundler

type DependencyBundlerConfiguration struct {
	Port string `json:"port"`
	ContextHeaders []string `json:"contextHeaders"`
	SessionCollection string `json:"sessionCollection"`
	LogToFile string `json:"logToFile"`
	MinimumSeverityLevel string `json:"minimumSeverityLevel"`
	LogToConsole string `json:"logToConsole"`
	FileJsonFormat string `json:"fileJsonFormat"`
	ConsoleJsonFormat string `json:"consoleJsonFormat"`
	CompressLogs string `json:"compressLogs"`
	DockerName string `json:"DOCKER_NAME"`
}

