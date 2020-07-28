package utils

import (
	"github.com/sirupsen/logrus"
)

// LogLevel of quic-go
type LogLevel uint8

const (
	// LogLevelNothing disables
	LogLevelNothing LogLevel = iota
	// LogLevelError enables err logs
	LogLevelError
	// LogLevelInfo enables info logs (e.g. packets)
	LogLevelInfo
	// LogLevelDebug enables debug logs (e.g. packet contents)
	LogLevelDebug
)

// A Logger logs.
type Logger interface {
	SetLogLevel(LogLevel)
	SetLogTimeFormat(format string)
	WithPrefix(prefix string) Logger
	Debug() bool

	Errorf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Debugf(format string, args ...interface{})
}

// NewDefaultLogger constructs a new Logger.
func NewDefaultLogger(le *logrus.Entry) Logger {
	if le == nil {
		log := logrus.New()
		log.SetLevel(logrus.InfoLevel)
		le = logrus.NewEntry(log)
	}
	return &defaultLogger{Entry: le}
}

type defaultLogger struct {
	*logrus.Entry
	prefix string
}

var _ Logger = &defaultLogger{}

// SetLogLevel sets the log level
func (l *defaultLogger) SetLogLevel(level LogLevel) {
	// noop
}

// SetLogTimeFormat sets the format of the timestamp
// an empty string disables the logging of timestamps
func (l *defaultLogger) SetLogTimeFormat(format string) {
	// noop
}

func (l *defaultLogger) WithPrefix(prefix string) Logger {
	if len(l.prefix) > 0 {
		prefix = l.prefix + " " + prefix
	}
	return &defaultLogger{
		Entry:  l.Entry,
		prefix: prefix,
	}
}

func (l *defaultLogger) Debugf(fmt string, args ...interface{}) {
	if l.Debug() {
		l.Entry.Debugf(fmt, args)
	}
}

// Debug returns true if the log level is LogLevelDebug
func (l *defaultLogger) Debug() bool {
	return l.Entry.Level >= logrus.DebugLevel
}
