package redis

import (
	"github.com/garyburd/redigo/redis"
)

func (this *Redis) DEL(k ...interface{}) (int64, error) {
	nCount, err := redis.Int64(this.Do("DEL", k...))
	switch err {
	case nil:
	default:
		return 0, err
	}
	return nCount, nil
}
