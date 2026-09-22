package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

// Init 初始化全局日志
func Init(cfg *Config) (*zap.Logger, error) {
	if cfg.LogDir == "" {
		cfg.LogDir = "logs"
	}
	if cfg.LogFile == "" {
		cfg.LogFile = "app"
	}

	encoderConfig := zap.NewProductionConfig()
	encoderConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logLevel, err := zapcore.ParseLevel(cfg.Level)
	if err != nil {
		return nil, fmt.Errorf("parse log level error: %w", err)
	}
	encoderConfig.Level = zap.NewAtomicLevelAt(logLevel)

	output, err := NewLogWriter(cfg.LogDir, cfg.LogFile, cfg.MaxSize, cfg.MaxAge)
	if err != nil {
		return nil, fmt.Errorf("create log writer error: %w", err)
	}

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig.EncoderConfig),
		zapcore.AddSync(output),
		encoderConfig.Level,
	)

	log = zap.New(core)
	return log, nil
}

// InitWithDefault 使用默认配置初始化
func InitWithDefault(logFile string) (*zap.Logger, error) {
	return Init(DefaultConfig(logFile))
}

// Get 获取全局日志实例
func Get() *zap.Logger {
	return log
}

// Close 关闭日志
func Close() {
	if log != nil {
		log.Sync()
	}
}
