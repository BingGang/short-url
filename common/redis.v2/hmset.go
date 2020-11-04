package redis

import (
	"github.com/garyburd/redigo/redis"
)

func (this *Redis) HMSET(k string, v map[interface{}]interface{}) (bool, error) {
	args := []interface{}{k}
	for k,value:= range v{
		args = append(args,k)
		args = append(args,value)
	}
	result, err := redis.String(this.Do("HMSET", args...))
	switch err {
	case nil:
	case redis.ErrPoolExhausted:
		return false, nil
	default:
		return false, err
	}

	return result == "OK", nil
}
