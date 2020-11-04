package redis

import (
	"errors"
	"short_url/dao/redis/server_cache1"
	//"short_url/dao/redis/im_server_cache"
	log "short_url/common/log4go"
)

var (
	ServerCache1 *server_cache1.ServerCache1
	//CidServerCache *server_cache1.CidServerCache
	ErrConfig   = errors.New("redis config error")
	ErrNoCreate = errors.New("redis object not create")
)

type Config map[string]*struct {
	Addr      string
	MaxActive int
	MaxIdle   int
	Timeout   int
	Password  string
}

// 初始化
func Init(cacheConf Config) error {
	if c, ok := cacheConf["server_cache1"]; ok {
		ServerCache1 = server_cache1.NewServerCache1()
		if err := ServerCache1.Reload(c.Addr, c.MaxActive, c.MaxIdle, c.Timeout, c.Password); err != nil {
			log.Error(" config err c.Add = %v, c.MaxActive = %v, c.MaxIdle = %v, c.Timeout = %v, c.Password = %v)", c.Addr, c.MaxActive, c.MaxIdle, c.Timeout, c.Password)
			return err
		}
	}
	if ServerCache1 == nil {
		log.Error(" cache err")
		return ErrConfig
	}
	return nil
}

// 关闭
func Close() {
	if ServerCache1 != nil {
		ServerCache1.Close()
	}
	//支持多redis
}
