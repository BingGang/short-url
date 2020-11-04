package redis

import (
	"github.com/garyburd/redigo/redis"
)

func (this *Redis) ZREVRANGE(k string, s interface{}, e...interface{}) ([]interface{}, error) {
	arg := []interface{}{k}
	arg = append(arg,s)
	arg = append(arg,e...)
	buf, err := redis.Values(this.Do("ZREVRANGE",arg...))
	switch err {
	case nil:
	case redis.ErrNil:
		return nil, nil
	default:
		return nil, err
	}
	return buf, nil
}
