package zap

import (
	"github.com/gordonmleigh/byolog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLogger struct {
	zap *zap.Logger
}

// New creates a byolog instance from a zap instance
func New(logger *zap.Logger) byolog.Logger {
	return &zapLogger{
		zap: logger.WithOptions(zap.AddCallerSkip(1)),
	}
}

func (l *zapLogger) Error(msg string, fields ...byolog.Field) {
	l.zap.Error(msg, convertFields(fields)...)
}

func (l *zapLogger) Warn(msg string, fields ...byolog.Field) {
	l.zap.Warn(msg, convertFields(fields)...)
}

func (l *zapLogger) Info(msg string, fields ...byolog.Field) {
	l.zap.Info(msg, convertFields(fields)...)
}

func (l *zapLogger) Debug(msg string, fields ...byolog.Field) {
	l.zap.Debug(msg, convertFields(fields)...)
}

func (l *zapLogger) With(fields ...byolog.Field) byolog.Logger {
	return New(l.zap.With(convertFields(fields)...))
}

func (l *zapLogger) Named(name string) byolog.Logger {
	return New(l.zap.Named(name))
}

func convertFields(fields []byolog.Field) []zapcore.Field {
	ret := make([]zapcore.Field, len(fields))
	for i, f := range fields {
		ret[i] = zap.Any(f.Name, f.Value)
	}
	return ret
}
