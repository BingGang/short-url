package service

import (
	"short_url/common/snowflake"
	"short_url/conf"
	"time"
)

var (
	idWorker     snowflake.IdWorker
	httpMaxRetry = 3
	readTimeout  = 10 * time.Second
)

func Init() error {
	idWorker = snowflake.NewIdWorker(conf.Conf.Snowflake)
	return nil
}






