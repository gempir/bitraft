package log

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

var logrusLogger = logrus.New()

// Fields wraps logrus.Fields, which is a map[string]interface{}
type Fields logrus.Fields

func SetLogLevel(level logrus.Level) {
	logrusLogger.Level = level
}

func SetLogFormatter(formatter logrus.Formatter) {
	logrusLogger.Formatter = formatter
}

// Debug logs a message at level Debug on the standard logrusLogger.
func Debug(args ...interface{}) {
	if logrusLogger.Level >= logrus.DebugLevel {
		entry := logrusLogger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Debug(args...)
	}
}

// Debug logs a message at level Debug on the standard logrusLogger.
func Debugf(args ...interface{}) {
	if logrusLogger.Level >= logrus.DebugLevel {
		entry := logrusLogger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Debug(args...)
	}
}

// Info logs a message at level Info on the standard logrusLogger.
func Info(args ...interface{}) {
	if logrusLogger.Level >= logrus.InfoLevel {
		entry := logrusLogger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Info(args...)
	}
}

// Info logs a message at level Info on the standard logrusLogger.
func Infof(format string, args ...interface{}) {
	if logrusLogger.Level >= logrus.InfoLevel {
		entry := logrusLogger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Infof(format, args...)
	}
}

// Warn logs a message at level Warn on the standard logrusLogger.
func Warn(args ...interface{}) {
	if logrusLogger.Level >= logrus.WarnLevel {
		entry := logrusLogger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Warn(args...)
	}
}

// Warn logs a message at level Warn on the standard logrusLogger.
func Warnf(args ...interface{}) {
	if logrusLogger.Level >= logrus.WarnLevel {
		entry := logrusLogger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Warn(args...)
	}
}

// Error logs a message at level Error on the standard logrusLogger.
func Error(args ...interface{}) {
	if logrusLogger.Level >= logrus.ErrorLevel {
		entry := logrusLogger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Error(args...)
	}
}

// Error logs a message at level Error on the standard logrusLogger.
func Errorf(args ...interface{}) {
	if logrusLogger.Level >= logrus.ErrorLevel {
		entry := logrusLogger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Error(args...)
	}
}

// Fatal logs a message at level Fatal on the standard logrusLogger.
func Fatal(args ...interface{}) {
	if logrusLogger.Level >= logrus.FatalLevel {
		entry := logrusLogger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Fatal(args...)
	}
}

// Fatal logs a message at level Fatal on the standard logrusLogger.
func Fatalf(args ...interface{}) {
	if logrusLogger.Level >= logrus.FatalLevel {
		entry := logrusLogger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Fatal(args...)
	}
}

// Panic logs a message at level Panic on the standard logrusLogger.
func Panic(args ...interface{}) {
	if logrusLogger.Level >= logrus.PanicLevel {
		entry := logrusLogger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Panic(args...)
	}
}

// Panic logs a message at level Panic on the standard logrusLogger.
func Panicf(args ...interface{}) {
	if logrusLogger.Level >= logrus.PanicLevel {
		entry := logrusLogger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Panic(args...)
	}
}

func fileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}
