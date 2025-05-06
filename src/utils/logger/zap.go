package logger

import (
	"os"

	genericConstants "omnenest-backend/src/constants"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logLevelMap = map[string]zapcore.Level{
	"debug": zap.DebugLevel,
	"info":  zap.InfoLevel,
	"error": zap.ErrorLevel,
	"fatal": zap.FatalLevel,
	"panic": zap.PanicLevel,
}

type zapLogger struct {
	Log *zap.Logger
}

// SetupLogging initializes the logger with the specified log level.
func SetupLogging(level string) {
	var logLevel zapcore.Level
	var ok bool
	logLevel, ok = logLevelMap[level]
	if !ok {
		logLevel = zap.InfoLevel
	}

	cfg := zapcore.EncoderConfig{
		MessageKey:  genericConstants.MessageKey,
		LevelKey:    genericConstants.LogLevelKey,
		EncodeLevel: LevelEncoder,
		TimeKey:     genericConstants.TimeLogParam,
		EncodeTime:  LogTimeEncoder,
		FunctionKey: genericConstants.MethodLogParam,
	}

	var cores []zapcore.Core
	consoleErrors := zapcore.Lock(os.Stdout)
	jsonEncoder := NewSensitiveFieldsEncoder(cfg)
	consoleCore := zapcore.NewCore(jsonEncoder, consoleErrors, logLevel)
	cores = append(cores, consoleCore)
	tee := zapcore.NewTee(cores...)
	log = &zapLogger{zap.New(tee, zap.AddCaller(), zap.AddCallerSkip(1))}
}

func (l *zapLogger) With(args ...zap.Field) Logger {
	return &zapLogger{
		Log: l.Log.With(args...),
	}
}

func (l *zapLogger) Info(msg string, args ...zap.Field) {
	l.Log.Info(msg, args...)
}

func (l *zapLogger) Warn(msg string, args ...zap.Field) {
	l.Log.Warn(msg, args...)
}

func (l *zapLogger) Debug(msg string, args ...zap.Field) {
	l.Log.Debug(msg, args...)
}

func (l *zapLogger) Error(msg string, args ...zap.Field) {
	l.Log.Error(msg, args...)
}

func (l *zapLogger) Fatal(msg string, args ...zap.Field) {
	l.Log.Fatal(msg, args...)
}

func (l *zapLogger) Panic(msg string, args ...zap.Field) {
	l.Log.Panic(msg, args...)
}

func (l *zapLogger) Sync() {
	_ = l.Log.Sync()
}
