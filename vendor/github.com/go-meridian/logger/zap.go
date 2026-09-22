package logger

import "go.uber.org/zap"

// Logger 封装 zap.Logger，提供常用的日志方法
type Logger struct {
	zl *zap.Logger
}

// Wrap 将 *zap.Logger 包装为 Logger
func Wrap(zl *zap.Logger) *Logger {
	return &Logger{zl: zl}
}

// L 获取全局 Logger 实例
func L() *Logger {
	return &Logger{zl: log}
}

// With 添加字段，返回新的 Logger
func (l *Logger) With(fields ...Field) *Logger {
	return &Logger{zl: l.zl.With(fields...)}
}

// Debug 记录 debug 级别日志
func (l *Logger) Debug(msg string, fields ...Field) {
	l.zl.Debug(msg, fields...)
}

// Info 记录 info 级别日志
func (l *Logger) Info(msg string, fields ...Field) {
	l.zl.Info(msg, fields...)
}

// Warn 记录 warn 级别日志
func (l *Logger) Warn(msg string, fields ...Field) {
	l.zl.Warn(msg, fields...)
}

// Error 记录 error 级别日志
func (l *Logger) Error(msg string, fields ...Field) {
	l.zl.Error(msg, fields...)
}

// Fatal 记录 fatal 级别日志并退出程序
func (l *Logger) Fatal(msg string, fields ...Field) {
	l.zl.Fatal(msg, fields...)
}

// DPanic 在开发环境下 panic，生产环境下记录 error
func (l *Logger) DPanic(msg string, fields ...Field) {
	l.zl.DPanic(msg, fields...)
}

// Sync 刷新日志缓冲区
func (l *Logger) Sync() error {
	return l.zl.Sync()
}

// Named 为 Logger 添加名称前缀
func (l *Logger) Named(s string) *Logger {
	return &Logger{zl: l.zl.Named(s)}
}

// Z 返回底层 *zap.Logger，用于需要直接操作 zap 的场景
func (l *Logger) Z() *zap.Logger {
	return l.zl
}
