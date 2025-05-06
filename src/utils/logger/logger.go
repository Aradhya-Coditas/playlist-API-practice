package logger

import (
	"context"

	genericConstants "omnenest-backend/src/constants"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log Logger

type Logger interface {
	Info(string, ...zap.Field)
	Warn(string, ...zap.Field)
	Error(string, ...zap.Field)
	Debug(string, ...zap.Field)
	Fatal(string, ...zap.Field)
	Panic(string, ...zap.Field)
	With(args ...zap.Field) Logger
	Sync()
}

func GetLogger(ctx context.Context) Logger {
	if ctx.Value(genericConstants.RequestIDHeader) != nil {
		logger := log.With(zap.String(genericConstants.RequestIDHeader, ctx.Value(genericConstants.RequestIDHeader).(string)))
		return logger
	}
	return log
}

func GetLoggerWithoutContext() Logger {
	return log
}

func LogLatency(log Logger, start time.Time, msgLog string) {
	log.With(zap.Int64(genericConstants.LatencyKey, time.Since(start).Milliseconds())).Info(msgLog)
}

var (
	once sync.Once
	loc  *time.Location
)

/*
The LogTimeEncoder function formats the given time value using the time.LoadLocation and time.Format functions.
It appends the formatted time value to the provided encoder.
*/
func LogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	once.Do(func() {
		loc, _ = time.LoadLocation(genericConstants.IST)
	})
	enc.AppendString(t.In(loc).Format(genericConstants.TimeFormat))
}

func LevelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(l.CapitalString())
}
