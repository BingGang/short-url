package redis

import (
	"common/redis/conn"
	"log"
	"testing"
)

func TestRegisterPlan(t *testing.T) {
	xx := conn.Config{}
	xx.Addr = "10.33.21.150:6380"
	xx.MaxActive = 10
	xx.MaxIdle = 10
	xx.Timeout = 1
	log.Println(conn.Init(xx))
	r, e := command.SET("test", 1)
	log.Println(r, e)
	log.Println(command.GET("test"))
}
