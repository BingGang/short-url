package log4go

import (
	"io"
	"sync"
)

type Writer struct {
	in  io.Reader
	out io.Writer
	wg  sync.WaitGroup
}

func NewWriter(in io.Reader, out io.Writer) *Writer {
	w := &Writer{in: in, out: out}
	w.wg.Add(1)
	go w.flush()
	return w
}

func (this *Writer) Wait() {
	this.wg.Wait()
}

func (this *Writer) flush() {
	defer this.wg.Done()
	var (
		data []byte
		nlen int
		err  error
	)
	data = make([]byte, 2048)
	for {
		if nlen, err = this.in.Read(data); err != nil {
			panic(err)
		}
		if nlen == 0 {
			break
		}
		this.out.Write(data[:nlen])
		data = data[:]
	}
}
