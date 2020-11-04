package controller

import (
	"net/http"
	"short_url/ecode"
	"short_url/service"
	"strconv"
	"time"
)


func setShortUrl(wr http.ResponseWriter, r *http.Request) {
	res := make(map[string]interface{})
	res["data"] = make(map[string]interface{})
	res["msg"] = ""
	res["toast"] = ""
	var ret int64 = ecode.OK
	res["code"] = ret
	defer retGetWriter(r, wr, time.Now(), res)
	reqUrl := r.FormValue("req_url")
	ex := r.FormValue("ex")
	if reqUrl == "" || ex==""{
		res["code"] = ecode.ParamError
		return
	}
	exInt,err :=strconv.ParseInt(ex,10,64)
	if err != nil {
		res["code"] = ecode.ParamError
		return
	}
	shortUrlManagerService := service.ShortUrlManager{}
	detail, ret := shortUrlManagerService.SetUrl(reqUrl, exInt)
	if ret == ecode.NotExist {
		ret = ecode.OK
		var cc = make([]interface{}, 0)
		res["data"] = cc
	} else {
		res["data"] = detail
	}
	res["code"] = ret
	return
}
