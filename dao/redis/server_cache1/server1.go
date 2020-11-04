package server_cache1

import (
	"errors"

	"short_url/common/redis.v2"
)

var (
	ErrAddrIsEmpty   = errors.New("addr is empty")
	ErrMaxOpenIsZero = errors.New("maxopen is 0")
	ErrMaxIdleIsZero = errors.New("maxidle is 0")
	redisCache       *redis.Redis
)

type ServerCache1 struct {
	redisCache *redis.Redis
}

// 创建对象
func NewServerCache1() *ServerCache1 {
	return &ServerCache1{}
}

// 关闭
func (this *ServerCache1) Close() {
	if this.redisCache != nil {
		this.redisCache.Close()
	}
}

// 加载
func (this *ServerCache1) Reload(addr string, maxopen, maxidle, timeout int, password string) error {
	if len(addr) == 0 {
		return ErrAddrIsEmpty
	}
	if maxopen == 0 {
		return ErrMaxOpenIsZero
	}
	if maxidle == 0 {
		return ErrMaxIdleIsZero
	}
	redisCache = &redis.Redis{Addr: addr, MaxActive: maxopen, MaxIdle: maxidle, Timeout: timeout, Password: password}
	err := redisCache.Init()
	if err != nil {
		return err
	}
	var tdb *redis.Redis
	tdb, this.redisCache = this.redisCache, redisCache

	if tdb != nil {
		tdb.Close()
	}
	return nil
}
