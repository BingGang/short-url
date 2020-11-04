package redis

import (
	"github.com/garyburd/redigo/redis"
)

func (this *Redis) HMGET(k string, v...interface{}) (interface{}, error) {
	arg := []interface{}{k}
	arg = append(arg,v...)
	buf, err := this.Do("HMGET",arg...)
	switch err {
	case nil:
	case redis.ErrNil:
		return nil, nil
	default:
		return nil, err
	}
	return buf, nil
}
