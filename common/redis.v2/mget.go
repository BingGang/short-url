package redis

import (
	"github.com/garyburd/redigo/redis"
)

func (this *Redis) MGET(v []interface{}) (interface{}, error) {
	if len(v) == 0 {
		return nil,nil
	}
	arg := []interface{}{}
	for _,item := range v   {
		arg = append(arg,item)
	}
	buf, err := this.Do("MGET",arg...)
	switch err {
	case nil:
	case redis.ErrNil:
		return nil, nil
	default:
		return nil, err
	}
	return buf, nil
}
