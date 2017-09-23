/*
	Copyright 2017 by chspidor author.
*/
package spider

import (
	"chspider/base"
	//	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

func NewJar() *cookiejar.Jar {
	cookieJar, _ := cookiejar.New(nil)
	return cookieJar
}

var (
	//mylog
	//	mylog = base.Mylog()

	//要管理HTTP客户端的头域、重定向策略和其他设置，创建一个Client
	Client = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			//			mylog.Debugf("-----------Redirect:", req.URL)
			//	mylog.Debuf("111", req.URL)
			return nil
		},
		Jar: NewJar(),
	}

	// // 没有cookie的客户端
	NoCookieClient = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			//			mylog.Debugf("-----------Redirect:", req.URL)
			return nil
		},
	}
	//不设置超时时间
	DefaultTimeOut = 0
)

// a proxy client
// 带代理客户端，全部有带cookie
func NewProxyClient(proxystring string) (*http.Client, error) {
	proxy, err := url.Parse(proxystring)
	if err != nil {
		return nil, err
	}
	client := &http.Client{
		// allow redirect
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			//			mylog.Debugf("-----------Redirect:", req.URL)
			return nil
		},
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
		Jar:     NewJar(),
		Timeout: base.Second(DefaultTimeOut),
	}
	return client, nil
}

// a client
// 不带代理客户端
func NewClient() (*http.Client, error) {
	client := &http.Client{
		// allow redirect
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			//			mylog.Debugf("-----------Redirect:", req.URL)
			return nil
		},
		Jar:     NewJar(),
		Timeout: base.Second(DefaultTimeOut),
	}
	return client, nil
}
