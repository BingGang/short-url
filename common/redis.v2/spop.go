package redis

import (
	"github.com/garyburd/redigo/redis"
)

func (this *Redis) SPOP(k string) (interface{}, error) {
	buf, err := this.Do("SPOP", k)
	switch err {
	case nil:
	case redis.ErrNil:
		return nil, nil
	default:
		return nil, err
	}
	return buf, nil
}
