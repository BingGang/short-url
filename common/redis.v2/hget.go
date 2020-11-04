package redis

import (
	"github.com/garyburd/redigo/redis"
)

func (this *Redis) HGET(k string, v interface{}) (interface{}, error) {
	buf, err := this.Do("HGET", k, v)
	switch err {
	case nil:
	case redis.ErrNil:
		return nil, nil
	default:
		return nil, err
	}
	return buf, nil
}
