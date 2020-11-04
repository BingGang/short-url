package redis

import (
	redisP "github.com/garyburd/redigo/redis"
	"time"
)

type Redis struct {
	Addr      string
	MaxActive int
	MaxIdle   int
	Timeout   int
	Password  string
	Pool *redisP.Pool
}

func (this *Redis) Init() (err error) {
	pool := &redisP.Pool{
		MaxIdle:  this.MaxIdle,
		MaxActive: this.MaxActive,
		Dial: func() (redisP.Conn, error) {
			c, err := redisP.DialTimeout("tcp", this.Addr,time.Duration(this.Timeout) * time.Second,time.Duration(this.Timeout) * time.Second,time.Duration(this.Timeout) * time.Second)
			if err != nil {
				return nil, err
			}
			if this.Password != "" {
				if _, err := c.Do("AUTH", this.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, nil
		},
		// Use the TestOnBorrow function to check the health of an idle connection
		// before the connection is returned to the application.
		TestOnBorrow: func(c redisP.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
		IdleTimeout: time.Duration(this.Timeout) * time.Second,
	}
	/*
	pool := &redisP.Pool{
		MaxIdle:     this.MaxIdle,
		MaxActive:   this.MaxActive,
		IdleTimeout: time.Duration(this.Timeout) * time.Second,
		Dial: func() (redisP.Conn, error) {
			c, err := redisP.Dial("tcp", this.Addr)
			if err != nil {

				return nil, err
			}
			if this.Password != "" {
				if _, err := c.Do("AUTH", this.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, nil
		},
	}*/

	var t_pool *redisP.Pool
	t_pool, this.Pool = this.Pool, pool
	if t_pool != nil {
		t_pool.Close()
	}
	return nil
}

func(this *Redis) Close() {
	if this.Pool != nil {
		this.Pool.Close()
	}
}

func (this *Redis) Do(cmd string, args ...interface{}) (interface{}, error) {
	c := this.Pool.Get()
	if err := c.Err(); err != nil {
		return nil, err
	}
	defer c.Close()
	reply,err :=  c.Do(cmd, args...)
	if err!=nil{
		c := this.Pool.Get()
		if err := c.Err(); err != nil {
			return nil, err
		}
		defer c.Close()
		return c.Do(cmd, args...)
	}
	return reply,err
}
