package redis

import (
	"github.com/garyburd/redigo/redis"
)

func (this *Redis) RPOP(k string) (interface{}, error) {
	buf, err := this.Do("RPOP", k)
	switch err {
	case nil:
	case redis.ErrNil:
		return nil, nil
	default:
		return nil, err
	}
	return buf, nil
}
