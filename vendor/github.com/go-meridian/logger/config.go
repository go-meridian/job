package logger

// Config 日志配置
type Config struct {
	Level   string `json:"level" yaml:"level"`     // 日志级别: debug/info/warn/error
	LogFile string `json:"logFile" yaml:"logFile"` // 日志文件名前缀
	LogDir  string `json:"logDir" yaml:"logDir"`   // 日志目录，默认 "logs"
	MaxSize int    `json:"maxSize" yaml:"maxSize"` // 单文件最大 MB，默认 500
	MaxAge  int    `json:"maxAge" yaml:"maxAge"`   // 旧日志保留天数，默认 30
}

// DefaultConfig 返回默认配置
func DefaultConfig(logFile string) *Config {
	return &Config{
		Level:   "info",
		LogFile: logFile,
		LogDir:  "logs",
		MaxSize: 500,
		MaxAge:  30,
	}
}
