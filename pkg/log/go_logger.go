package log

import (
	"context"
)

type goLogger struct {
}

func NewGoLogger() FieldLogger {
	return &goLogger{}
}

func (m *goLogger) WithError(err error) FieldLogger {
	return m
}

func (m *goLogger) WithContext(ctx context.Context) FieldLogger {
	return m
}

func (m *goLogger) WithField(key string, value interface{}) FieldLogger {
	return m
}

func (m *goLogger) WithFields(fields Fields) FieldLogger {
	return m
}

func (m *goLogger) Debug(args ...interface{}) {
}

func (m *goLogger) Debugf(format string, args ...interface{}) {
}

func (m *goLogger) Info(args ...interface{}) {
}

func (m *goLogger) Infof(format string, args ...interface{}) {
}

func (m *goLogger) Warn(args ...interface{}) {
}

func (m *goLogger) Warnf(format string, args ...interface{}) {
}

func (m *goLogger) Error(args ...interface{}) {
}

func (m *goLogger) Errorf(format string, args ...interface{}) {
}
