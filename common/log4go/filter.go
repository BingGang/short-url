package log4go

import (
	"log"
)

type FileLogWriter interface {
	Write([]byte) (int, error)
	Close()
}

type LogWriter struct {
	in     *Buffer
	out    FileLogWriter
	writer *Writer
	logger *log.Logger
}

func (this *LogWriter) Wait() {
	this.writer.Wait()
	this.out.Close()
}

func (this *LogWriter) Close() error {
	return this.in.Close()
}

func (this *LogWriter) Logger() *log.Logger {
	return this.logger
}
