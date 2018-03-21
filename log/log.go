package log

import (
	"os"

	"github.com/sirupsen/logrus"
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
	l := logrus.New()
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		l.Out = file
	} else {
		l.Errorf("Failed to log to file %s, using default stderr %v", path, err)
	}
	return l
}
