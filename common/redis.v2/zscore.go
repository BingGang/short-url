package redis

import (
	"github.com/garyburd/redigo/redis"
)

func (this *Redis) ZSCORE(k string, v interface{}) (int64, error) {
	buf, err := redis.Int64(this.Do("ZSCORE", k, v))
	switch err {
	case nil:
	case redis.ErrNil:
		return 0, nil
	default:
		return 0, err
	}
	return buf, nil
}
