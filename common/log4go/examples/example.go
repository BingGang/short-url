package main

import (
	log "common/log4go"
	"time"
)

func main() {
	if err := log.LoadConfiguration("logformat.xml"); err != nil {
		return
	}
	defer log.Close()
	ll := log.Logger(log.ERROR)
	log.Debug("debug")
	ll.Println("xxx")
	log.Debug("debug2")
	log.Trace("trace")

	//ll.Println("xxx")
	log.Warn("warn")
	//ll.Println("xxx")
	log.Error("error")

	for {
		log.Info("info asdadadlkjadadlkjalkjdalkjdlkjadlkjalkjdalkjdlkjadlkjada")
		time.Sleep(time.Second)
	}
}
