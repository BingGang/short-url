package redis

import (
	"github.com/garyburd/redigo/redis"
)

func (this *Redis) LREM(k string,count int64,v interface{}) (int64, error) {

	value, err := redis.Int64(this.Do("LREM", k,count, v ))
	switch err {
	case nil:
	case redis.ErrNil:
		return 0, nil
	default:
		return 0, err
	}
	return value, nil
}
