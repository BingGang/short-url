package httpclient

import (
	"bytes"
	log "code.google.com/p/log4go"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"
	"io"
	"errors"
	"fmt"
)

const (
	httpMaxRetry = 3
	httpTimeout  = 10 * time.Second
	httpDeadline = 10 * time.Second
)

var (
	client  *http.Client
	bufPool *sync.Pool
)

func init() {
	httpTransport := &http.Transport{
		Dial: func(netw, addr string) (net.Conn, error) {
			c, err := net.DialTimeout(netw, addr, httpTimeout)
			if err != nil {
				log.Error("net.DialTimeout(\"%s\", \"%s\", %d", netw, addr, httpTimeout)
				return nil, err
			}
			deadline := time.Now().Add(httpDeadline)
			c.SetDeadline(deadline)
			return c, nil
		},
		DisableKeepAlives: false,
	}
	client = &http.Client{
		Transport: httpTransport,
	}
	bufPool = &sync.Pool{
		New: func() interface{} {
			return &bytes.Buffer{}
		},
	}
}

func HttpPost(url string, data url.Values) ([]byte, error) {
	var (
		resp *http.Response
		err  error
	)
	for i := 1; i <= httpMaxRetry; i++ {
		resp, err = client.PostForm(url, data)
		if err != nil {
			log.Error("cmbHttpClient.Post(%s, %s) try(%d) resp(%v) error(%v)", url, data, i, resp, err)
			continue
		}
		break
	}
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("ioutil.ReadAll(resp.Body) error(%v)", err)
		return nil, err
	}
	return body, nil
}

func HttpPostBody(url string, bodyType string, b io.Reader) ([]byte, error) {
	var (
		res *http.Response
		err error
	)
	for i := 1; i <= httpMaxRetry; i++ {
		res, err = client.Post(url, bodyType, b)
		if err != nil {
			log.Error("http.Post(%s, %v) error(%v)", url, bodyType, err)
			continue
		}
		break
	}
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Error("http status:%d", res.StatusCode)
		return nil, errors.New(fmt.Sprintf("http status:%d", res.StatusCode))
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error("url(%s) ioutil.ReadAll() error(%v)", url, err)
		return nil, err
	}

	return body, nil
}


// httpGet http request using get method
func HttpGet(url string) ([]byte, error) {
	var (
		resp *http.Response
		err  error
	)
	for i := 1; i <= httpMaxRetry; i++ {
		resp, err = client.Get(url)
		if err != nil {
			log.Error("cmbHttpClient.Get(\"%s\") try(%d) error(%v)", url, i, err)
			continue
		}
		break
	}
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("ioutil.ReadAll(resp.Body) error(%v)", err)
		return nil, err
	}
	return body, nil
}

func HttpGetWithCookie(url string, cookie *http.Cookie) ([]byte, error) {
	var (
		resp *http.Response
		err  error
	)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.AddCookie(cookie)
	for i := 1; i <= httpMaxRetry; i++ {
		resp, err = client.Do(req)
		if err != nil {
			log.Error("client.Do(\"%s\") try(%d) error(%v)", url, i, err)
			continue
		}
		break
	}
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("ioutil.ReadAll(resp.Body) error(%v)", err)
		return nil, err
	}
	return body, nil
}