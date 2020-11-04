package log4go

import (
	"fmt"
	"sync"
)

type level int

const (
	DEBUG level = iota
	TRACE
	INFO
	WARNING
	ERROR
)

var (
	levelStrings = [...]string{"DEBG", "TRAC", "INFO", "WARN", "EROR"}
)

type Log4go struct {
	filters map[level][]*LogWriter
}

func NewLog4go() *Log4go {
	return &Log4go{make(map[level][]*LogWriter)}
}

func (this *Log4go) Close() {
	var wg sync.WaitGroup
	for level, filters := range this.filters {
		for _, filter := range filters {
			filter.Close()
			wg.Add(1)
			go func(f *LogWriter) {
				f.Wait()
				wg.Done()
			}(filter)
		}
		delete(this.filters, level)
	}
	wg.Wait()
}

func (this *Log4go) Loggers(l level) []*LogWriter {
	if filter, ok := this.filters[l]; ok {
		return filter
	}
	return nil
}

func (this *Log4go) Error(format string, v ...interface{}) {
	loggers := this.Loggers(ERROR)
	if loggers == nil {
		return
	}
	for _, logger := range loggers {
		logger.Logger().Output(2, fmt.Sprintf(format, v...))
	}
}

func (this *Log4go) Warn(format string, v ...interface{}) {
	loggers := this.Loggers(WARNING)
	if loggers == nil {
		return
	}
	for _, logger := range loggers {
		logger.Logger().Output(2, fmt.Sprintf(format, v...))
	}
}

func (this *Log4go) Info(format string, v ...interface{}) {
	loggers := this.Loggers(INFO)
	if loggers == nil {
		return
	}
	for _, logger := range loggers {
		logger.Logger().Output(2, fmt.Sprintf(format, v...))
	}
}

func (this *Log4go) Trace(format string, v ...interface{}) {
	loggers := this.Loggers(TRACE)
	if loggers == nil {
		return
	}
	for _, logger := range loggers {
		logger.Logger().Output(2, fmt.Sprintf(format, v...))
	}
}

func (this *Log4go) Debug(format string, v ...interface{}) {
	loggers := this.Loggers(DEBUG)
	if loggers == nil {
		return
	}
	for _, logger := range loggers {
		logger.Logger().Output(2, fmt.Sprintf(format, v...))
	}
}
