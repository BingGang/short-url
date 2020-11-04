package service

import (
	"short_url/dao/redis/server_cache1"
	"short_url/ecode"
	"short_url/conf"

)

type ShortUrlManager struct{}

func (t *ShortUrlManager) SetUrl(sourceUrl string, exTime int64) (shortUrl string, ret int64) {

	shortArr, err := Transform(sourceUrl)
	if err != nil {
		ret = ecode.ServerError
		return
	}

	t.setShortAddr(shortArr[1], sourceUrl, int64(exTime))
	shortUrl=conf.Conf.Base.Domain+"/"+shortArr[1]
	return
}
func (t *ShortUrlManager) setShortAddr(key, addr string, ex int64) (ret int64) {
	shortAddrDao := server_cache1.ShortAddrDao{}
	ret = shortAddrDao.Add(key, addr, ex)
	return
}


