package log4go

import (
	"fmt"
)

var (
	log4go = NewLog4go()
)

func LoadConfiguration(filename string) error {
	return log4go.LoadConfiguration(filename)
}

func Close() {
	log4go.Close()
}

func Loggers(l level) []*LogWriter {
	return log4go.Loggers(l)
}

func Error(format string, v ...interface{}) {
	loggers := log4go.Loggers(ERROR)
	if loggers == nil {
		return
	}
	for _, logger := range loggers {
		logger.Logger().Output(2, fmt.Sprintf(format, v...))
	}
}

func Warn(format string, v ...interface{}) {
	loggers := log4go.Loggers(WARNING)
	if loggers == nil {
		return
	}
	for _, logger := range loggers {
		logger.Logger().Output(2, fmt.Sprintf(format, v...))
	}
}

func Info(format string, v ...interface{}) {
	loggers := log4go.Loggers(INFO)
	if loggers == nil {
		return
	}
	for _, logger := range loggers {
		logger.Logger().Output(2, fmt.Sprintf(format, v...))
	}
}

func Trace(format string, v ...interface{}) {
	loggers := log4go.Loggers(TRACE)
	if loggers == nil {
		return
	}
	for _, logger := range loggers {
		logger.Logger().Output(2, fmt.Sprintf(format, v...))
	}
}

func Debug(format string, v ...interface{}) {
	loggers := log4go.Loggers(DEBUG)
	if loggers == nil {
		return
	}
	for _, logger := range loggers {
		logger.Logger().Output(2, fmt.Sprintf(format, v...))
	}
}
