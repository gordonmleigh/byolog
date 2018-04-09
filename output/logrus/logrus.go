package logrus

import (
	"github.com/gordonmleigh/byolog"
	"github.com/sirupsen/logrus"
)

type logrusLogger struct {
	logrus logrusSubset
	name   string
}

// logrusSubset is the common methods used from logrus.Logger and logrus.Entry.
type logrusSubset interface {
	Debugln(args ...interface{})
	Infoln(args ...interface{})
	Warnln(args ...interface{})
	Errorln(args ...interface{})
	WithField(name string, value interface{}) *logrus.Entry
}

// New creates a byolog instance from a logrus instance.
func New(logger *logrus.Logger) byolog.Logger {
	return &logrusLogger{
		logrus: logger,
	}
}

// Named creates a byolog instance from a logrus instance.
func Named(logger *logrus.Logger, name string) byolog.Logger {
	return &logrusLogger{
		logrus: logger,
		name:   name,
	}
}

// New creates a byolog instance from a logrus instance.
func newInternal(logger logrusSubset, name string) byolog.Logger {
	return &logrusLogger{
		logrus: logger,
		name:   name,
	}
}

func (l *logrusLogger) Error(msg string, fields ...byolog.Field) {
	withFields(l.logrus, fields).Errorln(msg)
}

func (l *logrusLogger) Warn(msg string, fields ...byolog.Field) {
	withFields(l.logrus, fields).Warnln(msg)
}

func (l *logrusLogger) Info(msg string, fields ...byolog.Field) {
	withFields(l.logrus, fields).Infoln(msg)
}

func (l *logrusLogger) Debug(msg string, fields ...byolog.Field) {
	withFields(l.logrus, fields).Debugln(msg)
}

func (l *logrusLogger) With(fields ...byolog.Field) byolog.Logger {
	return newInternal(withFields(l.logrus, fields), l.name)
}

func (l *logrusLogger) Named(name string) byolog.Logger {
	if l.name != "" {
		name = l.name + "/" + name
	}
	return newInternal(l.logrus.WithField("name", name), name)
}

func withFields(logger logrusSubset, fields []byolog.Field) logrusSubset {
	for _, f := range fields {
		logger = logger.WithField(f.Name, f.Value)
	}
	return logger
}
