package log

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
}

func NewDefaultLogger() Logger {
	l := logrus.New()
	l.SetLevel(logrus.DebugLevel)
	return l
}

func NewFileLogger(path string) Logger {
	l, err := NewLogger(path, zapcore.DebugLevel)
	if err != nil {
		return NewDefaultLogger()
	}
	return l
}
