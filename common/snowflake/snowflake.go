package snowflake

import (
	"sync"
	"time"
)

type Config struct {
	Machine    int64
	Datacenter int64
	Epoch      int64
}

type IdWorker interface {
	Generate() int64
}

var (
	mutex sync.Mutex
)

type idWorker struct {
	machine       int64
	datacenter    int64
	epoch         int64
	sequence      int64
	lasttimestamp int64
}

func NewIdWorker(c Config) IdWorker {
	return &idWorker{c.Machine & 0x1F, c.Datacenter & 0x1F, c.Epoch, 0, -1}
}

func (this *idWorker) Generate() (value int64) {
	mutex.Lock()
	defer mutex.Unlock()
	timeGen := func() int64 {
		return time.Now().UnixNano() / int64(time.Millisecond)
	}

	t := timeGen()
	if t != this.lasttimestamp {
		this.sequence = 0
		goto Generate
	}

	this.sequence = (this.sequence + 1) & 0xFFF
	if this.sequence == 0 {
		for {
			t = timeGen()
			if t > this.lasttimestamp {
				break
			}
		}
	}
Generate:
	this.lasttimestamp = t - this.epoch
	// 时间左移 12+5+5
	value = this.lasttimestamp << 22
	// 数据中心ID左移17位
	value |= this.datacenter << 17
	// 机器码ID左移12位
	value |= this.machine << 12
	// 最后12位
	value |= this.sequence
	return
}
