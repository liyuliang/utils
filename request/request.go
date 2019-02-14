package request

import (
	"time"
	"github.com/imroc/req"
	"errors"
	"strings"
	"net/url"
	"net/http"
	"github.com/liyuliang/utils/regex"
	"github.com/liyuliang/utils/format"
)

func GetHost(s string) string {
	u, err := url.Parse(s)
	if err == nil {
		return u.Scheme + "://" + u.Host
	}
	return ""
}

func UrlRemoveHost(url string) string {
	host := GetHost(url)
	url = strings.Replace(url, host, "", -1)
	return url
}

func GetReferer(uri string) (string, error) {
	U, err := url.Parse(uri)
	if err != nil {
		return "", err
	}

	referer := regex.Get(U.Host, `([^\.]+\.[^\.]+$)`)
	referer = U.Scheme + "://www." + referer
	return referer, nil
}

func DoReq(uri string, proxy string) (resp *Response) {
	t1 := time.Now()

	resp = new(Response)

	needReferer := IsUrlNeedReferer(uri)

	header := GetHeader()

	if needReferer {
		referer, err := GetReferer(uri)
		if err != nil {
			resp.Err = err
			return resp
		}
		header.Set("Referer", referer)

	}

	r := GetReq()

	if proxy != "" {
		r.SetProxyUrl("http://" + proxy)
	}

	result, err := r.Get(uri, header)
	if err != nil {
		resp.Err = err
		return resp
	}

	if result.Response().StatusCode != 200 {
		// Try again
		result, err = r.Get(uri, header)

		if err != nil {
			resp.Err = err
			return resp
		}

		if result.Response().StatusCode != 200 {
			resp.Err = errors.New("Http get request status code is not 200. ")
			return resp
		}
	}
	t2 := time.Now()

	resp.Err = err
	resp.StatusCode = result.Response().StatusCode
	resp.Data = result.String()
	resp.Speed = t2.Sub(t1).Seconds()

	result.Response().Body.Close()

	return resp
}

func IsUrlNeedReferer(uri string) bool {
	r := req.New()
	resp, err := r.Head(uri)
	if err != nil || (resp != nil && resp.Response().StatusCode == 403) {
		return true
	}
	return false
}

func GetReq() *req.Req {
	r := req.New()
	r.EnableInsecureTLS(true)
	r.SetTimeout(format.IntToTimeSecond(60))
	return r
}

func GetHeader() http.Header {
	header := make(http.Header)
	header.Set("User-Agent", AUserAgents())
	return header
}

func HttpGet(uri string) (httpResponse *Response) {
	resp := DoReq(uri, "")
	return resp
}
