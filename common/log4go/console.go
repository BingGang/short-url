package log4go

import (
	"os"
)

type Console struct {
	file *os.File
}

func NewConsole() *Console {
	return &Console{os.Stdout}
}

func (this *Console) Write(p []byte) (n int, err error) {
	return this.file.Write(p)
}

func (this *Console) Close() {
	this.file = nil
}
