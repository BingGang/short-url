package main

import (

	log "short_url/common/log4go"
	"short_url/common/perf"
	"short_url/common/process"
	"os"
	"os/signal"
	"runtime"
	"short_url/conf"
	"short_url/controller"
	"short_url/dao"
	"short_url/service"
	"syscall"
)

const default_conf = "./config.ini"

func main() {
	if len(os.Args) < 2 {
		if err := conf.Init(default_conf); err != nil {
			panic(err)
		}
	} else {
		if err := conf.Init(os.Args[1]); err != nil {
			panic(err)
		}
	}
	runtime.GOMAXPROCS(conf.Conf.Base.MaxProc)
	log.LoadConfiguration(conf.Conf.Base.LogPath)
	defer func() {
		log.Close()
	}()

	if err := perf.Init(conf.Conf.Base.PprofAddr); err != nil {
		log.Error("perf.Init() error(%v)", err)
		panic(err)
	}
	defer perf.Close()
	//
	//fmt.Println(conf.Conf.Http)
	if err := controller.Init(conf.Conf.Http); err != nil {
		panic(err)
	}
	defer controller.Close()

	if err := dao.Init(); err != nil {
		log.Error("dao.Init() error(%v)", err)
		panic(err)
	}
	defer dao.Close()

	service.Init()

	if err := process.Init(conf.Conf.Base.PidFile); err != nil {
		log.Error("process.Init(%s, %s, %s) error(%v)", conf.Conf.Base.User, conf.Conf.Base.DirPath, conf.Conf.Base.PidFile, err)
		panic(err)
	}
	log.Info("short_url  start")
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			return
		case syscall.SIGHUP:
			log.Info("im server reload")
			// TODO reload
		default:
			return
		}
	}
	log.Info("short_url back exit")
}
