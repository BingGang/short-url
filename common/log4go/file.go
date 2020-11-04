package log4go

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type File struct {
	ext        string
	name       string
	path       string
	size       int
	delay      int
	day        string
	num        int
	maxsize    int
	file       *os.File
	wg         sync.WaitGroup
	closeEvent chan bool
}

func NewFile(name string, size, delay int) *File {
	dir := filepath.Dir(name)
	ext := filepath.Ext(name)
	base := filepath.Base(name)
	name = strings.TrimSuffix(base, ext)
	f := &File{
		path:       dir,
		ext:        ext,
		name:       name,
		maxsize:    size,
		delay:      delay,
		closeEvent: make(chan bool, 1),
	}
	if err := f.Open(); err != nil {
		panic(err)
	}

	if f.delay > 0 {
		f.wg.Add(1)
		go f.flush()
	}
	return f
}

func (this *File) fixSize(nlen int) error {
	if this.maxsize == 0 || this.size+nlen < this.maxsize {
		return nil
	}
	old := this.file.Name()
	new := fmt.Sprintf("%s/%s-%s-%d%s", this.path, this.name, this.day, this.num,
		this.ext)
	this.num++
	this.file.Close()
	this.file = nil
	this.size = 0
	return os.Rename(old, new)
}

func (this *File) Open() (err error) {
	now := time.Now().Format("2006-01-02")
	if now != this.day {
		if this.file != nil {
			this.file.Close()
			this.file = nil
		}
		this.day = now
		this.size = 0
		this.num = 0
	}
	if this.file != nil {
		return
	}
	name := fmt.Sprintf("%s/%s-%s%s", this.path, this.name, this.day, this.ext)
	this.file, err = os.OpenFile(name, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0660)
	return
}

func (this *File) flush() {
	defer this.wg.Done()
	delay := time.Second * time.Duration(this.delay)
	t := time.NewTimer(delay)
	for {
		select {
		case <-this.closeEvent:
			this.file.Sync()
			return
		case <-t.C:
			this.file.Sync()
			t.Reset(delay)
		}
	}
}

func (this *File) Write(p []byte) (n int, err error) {
	this.Open()
	this.fixSize(len(p))
	n, err = this.file.Write(p)
	if err != nil {
		return
	}
	this.size += n
	if this.delay == 0 {
		this.file.Sync()
	}
	return
}

func (this *File) Close() {
	close(this.closeEvent)
	this.wg.Wait()
	this.file.Close()
	this.file = nil
}
