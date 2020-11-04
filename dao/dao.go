package dao

import (
	log "short_url/common/log4go"
	"short_url/conf"
	"short_url/dao/redis"
)

func Init() (err error) {

	if err = redis.Init(conf.Conf.Redis); err != nil {
		log.Error("redis.Init(%v) error(%v)", conf.Conf.Redis, err)
		return
	}
	return nil
}
func Close() {
	redis.Close()
}
