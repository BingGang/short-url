package server_cache1

import (
	"short_url/ecode"
	log "short_url/common/log4go"
)

type ShortAddrDao struct {
}

func (this *ShortAddrDao) Add(key, addr string, ex int64) int64 {
	_, err := redisCache.SETEX(key, addr, ex)
	if err != nil {
		log.Error(".Add() error(%v)", err)
		return ecode.ServerError
	}
	return ecode.OK
}
