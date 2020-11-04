package redis

import (
	"github.com/garyburd/redigo/redis"
	"reflect"
)

func (this *Redis) MSET(v map[interface{}]interface{}) (bool, error) {
	if len(v) == 0 {
		return false,nil
	}
	args := []interface{}{}
	for k,value:= range v{
		vTt := reflect.ValueOf(value)
		if vTt.IsNil() {
			return false,nil
		}
		args = append(args,k)
		args = append(args,value)
	}
	result, err := redis.String(this.Do("MSET", args...))
	switch err {
	case nil:
	case redis.ErrPoolExhausted:
		return false, nil
	default:
		return false, err
	}

	return result == "OK", nil
}
