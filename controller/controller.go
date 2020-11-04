package controller

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	log "short_url/common/log4go"
	"short_url/ecode"
	"strings"
	"time"
)

const (
	httpReadTimeout  = 15 * time.Second
	httpWriteTimeout = 15 * time.Second
)

var (
	httpListener []net.Listener
	buyurl       = ""

)

func initHttpHandler() *http.ServeMux {
	httpServeMux := http.NewServeMux()
	httpServeMux.HandleFunc("/api/1/set/short/url", setShortUrl)


	return httpServeMux
}

type Config struct {
	Addr []string
	Ping string
	Port string
}

func Init(httpconf Config) error {
	var err error

	if err != nil {
		log.Error("session is error %v", err)
		return err
	}
	httpServeMux := initHttpHandler()
	server := &http.Server{
		Handler:      httpServeMux,
		ReadTimeout:  httpReadTimeout,
		WriteTimeout: httpWriteTimeout,
	}
	for _, addr := range httpconf.Addr {
		l, err := net.Listen("tcp", addr)
		if err != nil {
			log.Error("tls.Listen(\"tcp\", \"%s\") error(%v)", addr, err)
			return err
		}
		httpListener = append(httpListener, l)
		go func() {
			log.Info("start http listen addr: %s", addr)
			if err := server.Serve(l); err != nil {
				log.Error("server.Serve(\"%s\") error(%v)", addr, err)
			}
		}()
	}
	return nil
}

func Close() {
	for _, l := range httpListener {
		l.Close()
	}
}

func retGetWriter(r *http.Request, wr http.ResponseWriter, start time.Time, result map[string]interface{}) {
	if len(result) == 0 {
		log.Info("[%s]get_url:%s(time:%f)", r.RemoteAddr, r.URL.String(), time.Now().Sub(start).Seconds())
		return
	}
	callback := r.FormValue("callback")
	paramJson, _ := json.Marshal(r.Form)
	ret := result["code"]
	retInt, _ := ret.(int64)
	msg := ecode.Ecode_intro[retInt]
	if msg != "" {
		result["msg"] = msg
	}
	if retInt == ecode.NotExist {
		result["error"] = ecode.NotExist
		ret = ecode.OK
	}
	result["code"] = ret
	if len(callback) == 0 {
		byteJson, err := json.Marshal(result)
		if err != nil {
			log.Error("json.Marshal(\"%v\") failed (%v)", result, err)
		}
		if _, err := wr.Write(byteJson); err != nil {
			log.Warn("wr.Write(\"%s\") failed (%v)", string(byteJson), err)
		}
		log.Info("[%s]get_url:%s(time:%f,ret:%s) param(%s)", getClientIp(r), r.URL.String(), time.Now().Sub(start).Seconds(), string(byteJson), string(paramJson))
		return
	}
	byteJson, err := json.Marshal(result)
	if err != nil {
		log.Error("json.Marshal(\"%v\") failed (%v)", result, err)
	}
	d := fmt.Sprintf("%s(%s)", callback, string(byteJson))
	if _, err := wr.Write([]byte(d)); err != nil {
		log.Warn("wr.Write(\"%s\") failed (%v)", d, err)
	}
	log.Info("[%s]get_url:[%s](time:%f,ret:%s) param(%s)", getClientIp(r), r.URL.String(), time.Now().Sub(start).Seconds(), d, string(paramJson))

}

func getClientIp(r *http.Request) string {
	remote := r.Header.Get("X-Forwarded-For")
	if remote == "" {
		return r.Header.Get("X-Real-IP")
	}
	idx := strings.LastIndex(remote, ",")
	if idx > -1 {
		remote = strings.TrimSpace(remote[idx+1:])
	}
	return remote
}


