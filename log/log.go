package log

import (
	l4g "cement/log/log4go"
)

type Logger interface {
	Debug(fmt string, args ...interface{})
	Info(fmt string, args ...interface{})
	Warn(fmt string, args ...interface{}) error
	Error(fmt string, args ...interface{}) error
	Close()
}

type LogLevel string

const (
	Debug LogLevel = "debug"
	Info  LogLevel = "info"
	Warn  LogLevel = "warn"
	Error LogLevel = "error"
)

func NewLog4jLogger(filename string, level LogLevel, maxSize, maxFileCount int) (Logger, error) {
	return NewLog4jLoggerWithFmt(filename, level, maxSize, maxFileCount, "")
}

func NewLog4jLoggerWithFmt(filename string, level LogLevel, maxSize, maxFileCount int, fmt string) (Logger, error) {
	logger := make(l4g.Logger)

	flw, err := l4g.NewFileLogWriter(filename, fmt, maxSize, maxFileCount)
	if err != nil {
		return nil, err
	}

	switch level {
	case Debug:
		logger.AddFilter("file", l4g.DEBUG, flw)
	case Info:
		logger.AddFilter("file", l4g.INFO, flw)
	case Warn:
		logger.AddFilter("file", l4g.WARNING, flw)
	case Error:
		logger.AddFilter("file", l4g.ERROR, flw)
	}

	return &logger, nil
}

func NewLog4jConsoleLogger(level LogLevel) Logger {
	switch level {
	case Debug:
		return l4g.NewDefaultLogger(l4g.DEBUG)
	case Info:
		return l4g.NewDefaultLogger(l4g.INFO)
	case Warn:
		return l4g.NewDefaultLogger(l4g.WARNING)
	case Error:
		return l4g.NewDefaultLogger(l4g.ERROR)
	default:
		panic("unkown level" + string(level))
	}
}
