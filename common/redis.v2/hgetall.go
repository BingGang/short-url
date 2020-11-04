package redis

import (
"github.com/garyburd/redigo/redis"
)

func (this *Redis) HGETALL(k string) (interface{}, error) {
	buf, err := this.Do("HGETALL", k)
	switch err {
	case nil:
	case redis.ErrNil:
		return nil, nil
	default:
		return nil, err
	}
	return buf, nil
}
