package logger

import (
	"os"
	"runtime"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Dir        string `toml:"dir" yaml:"dir" json:"dir"`
	Level      string `toml:"level" yaml:"level" json:"level"`
	Color      bool   `toml:"color" yaml:"color" json:"color"`
	Terminal   bool   `toml:"terminal" yaml:"terminal" json:"terminal"`
	ShowIp     bool   `toml:"show_ip" yaml:"show_ip" json:"show_ip"`
	TimeFormat string `toml:"time_format" yaml:"time_format" json:"time_format"`
}

type Logger struct {
	l *zap.Logger
}

func DefaultConfig() *Config {
	return &Config{
		Dir:        "./logs",
		Level:      "debug",
		Color:      true,
		Terminal:   true,
		ShowIp:     true,
		TimeFormat: "2006-01-02 15:04:05",
	}
}

func NewLogger(c *Config) *Logger {
	var level zapcore.Level
	switch c.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	caller := zap.AddCaller()
	development := zap.Development()
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.Lock(os.Stdout),
		level,
	)

	l := &Logger{
		l: zap.New(core, caller, development),
	}

	l.Info("Logger init success")
	return l
}

func (l *Logger) Zap() *zap.Logger {
	return l.l
}

func (l *Logger) Sugar() *zap.SugaredLogger {
	return l.l.Sugar()
}

// ======== zap sugar logger =========

func (l *Logger) Debug(args ...interface{}) {
	l.l.Sugar().Debug(args...)
}

func (l *Logger) Debugf(template string, args ...interface{}) {
	l.l.Sugar().Debugf(template, args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.l.Sugar().Info(args...)
}

func (l *Logger) Infof(template string, args ...interface{}) {
	l.l.Sugar().Infof(template, args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.l.Sugar().Warn(args...)
}

func (l *Logger) Warnf(template string, args ...interface{}) {
	l.l.Sugar().Warnf(template, args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.l.Sugar().Error(args...)
}

func (l *Logger) Errorf(template string, args ...interface{}) {
	l.l.Sugar().Errorf(template, args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.l.Sugar().Fatal(args...)
}

func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.l.Sugar().Fatalf(template, args...)
}

// ======= zap logger ========

func (l *Logger) With(fields ...zap.Field) *zap.Logger {
	return l.l.With(fields...)
}

func (l *Logger) ZapDebug(msg string, fields ...zap.Field) {
	l.l.Debug(msg, fields...)
}

func (l *Logger) ZapInfo(msg string, fields ...zap.Field) {
	l.l.Info(msg, fields...)
}

func (l *Logger) ZapWarn(msg string, fields ...zap.Field) {
	l.l.Warn(msg, fields...)
}

func (l *Logger) ZapError(msg string, fields ...zap.Field) {
	l.l.Error(msg, fields...)
}

func (l *Logger) getFileInfo() (file string, line int) {
	_, file, line, ok := runtime.Caller(3)

	if !ok {
		return "???", 1
	}

	if dirs := strings.Split(file, "/"); len(dirs) >= 2 {
		return dirs[len(dirs)-2] + "/" + dirs[len(dirs)-1], line
	}

	return
}
