package request

import (
	"time"
	"github.com/imroc/req"
	"utils/format"
	"errors"
	"log"
	"strings"
	"utils/regex"
	UrlPkg "net/url"
)

func getHost(s string) string {
	u, err := UrlPkg.Parse(s)
	if err == nil {
		return u.Scheme + "://" + u.Host
	}
	return ""
}

func urlRemoveHost(url string) string {
	host := getHost(url)
	url = strings.Replace(url, host, "", -1)
	return url
}

func GetReferer(uri string) (string, error) {
	Url, err := UrlPkg.Parse(uri)
	if err != nil {
		return "", err
	}

	referer := regex.Get(Url.Host, `([^\.]+\.[^\.]+$)`)
	referer = Url.Scheme + "://www." + referer
	return referer, nil
}

func doReq(uri string, proxy string) (resp *Response) {
	t1 := time.Now()

	resp = new(Response)

	referer, err := GetReferer(uri)
	if err != nil {
		resp.Err = err
		return resp
	}

	agent := AUserAgents()

	r := req.New()

	r.EnableInsecureTLS(true)
	r.SetTimeout(format.IntToTimeSecond(30))
	if proxy != "" {
		r.SetProxyUrl("http://" + proxy)
		r.SetTimeout(format.IntToTimeSecond(25))
	}

	result, err := r.Get(uri, req.Header{
		"User-Agent": agent,
		"Referer":    referer,
	})

	if err != nil {
		resp.Err = err
		return resp
	}

	if result.Response().StatusCode != 200 {
		// Try again
		result, err = r.Get(uri, req.Header{
			"User-Agent": agent,
			"Referer":    referer,
		})

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
	if proxy != "" {
		log.Printf("Using proxy ip: %s, use: %f", proxy, resp.Speed)
	}
	return resp
}

func HttpGet(uri string) (httpResponse *Response) {
	resp := doReq(uri, "")
	return resp
}
