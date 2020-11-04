package log4go

import (
	"bytes"
	"errors"
	"fmt"
	"runtime"
	"text/template"
	"time"
)

type Buffer struct {
	tmpl  *template.Template
	level string
	store chan []byte
}

func NewBuffer(format, level string) *Buffer {
	tmplText := adapter.Replace(string(format))
	if len(format) == 0 || format[len(format)-1] != '\n' {
		tmplText += "\n"
	}
	tmpl, err := template.New("Log").Parse(tmplText)
	if err != nil {
		panic(err)
	}
	return &Buffer{tmpl, level, make(chan []byte, 10240)}
}

func (this *Buffer) Close() (err error) {
	defer func() {
		if e := recover(); e != nil {
			strErr, _ := e.(string)
			err = errors.New(strErr)
		}
	}()
	close(this.store)
	return
}

func (this *Buffer) Write(p []byte) (n int, err error) {
	defer func() {
		if e := recover(); e != nil {
			strErr, _ := e.(string)
			n = 0
			err = errors.New(strErr)
		}
	}()
	util := make(map[string]interface{})
	util["LEVEL"] = this.level
	util["TIME"] = time.Now()
	util["MESSAGE"] = string(p[:len(p)-1])
	if _, file, line, ok := runtime.Caller(3); ok {
		util["SOURCE"] = fmt.Sprintf("%s:%d", file, line)
	}
	buf := bytes.NewBufferString("")
	this.tmpl.Execute(buf, util)
	this.store <- buf.Bytes()
	return len(p), nil
}

func (this *Buffer) Read(p []byte) (n int, err error) {
	b, ok := <-this.store
	if !ok {
		return 0, nil
	}
	return copy(p, b), nil
}
