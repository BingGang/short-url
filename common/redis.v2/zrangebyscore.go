package redis

import (
	"github.com/garyburd/redigo/redis"
)

func (this *Redis) ZRANGEBYSCORE(k string, max, min, limit, offset, count interface{}) ([]interface{}, error) {
	buf, err := redis.Values(this.Do("ZRANGEBYSCORE", k, max, min, limit, offset, count))
	switch err {
	case nil:
	case redis.ErrNil:
		return nil, nil
	default:
		return nil, err
	}
	return buf, nil
}