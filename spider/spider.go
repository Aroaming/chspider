package spider

import (
	"chspider/base"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

//	DefaultSpider *Spider
//var logger = base.Mylog()

func init() {
	//
	//UnInit()

	//默认爬虫
	chsp := new(Spider)
	chsp.SpiderConfig = new(SpiderConfig)
	chsp.Header = http.Header{}
	chsp.Data = url.Values{}
	chsp.BData = []byte{}
	chsp.Client = Client

	//全局爬虫使用曲剧客户端
	//	DefaultSpider = chsp

}

type SpiderConfig struct {
	Url    string      // now fetch url 这次要抓取的Url
	Method string      // Get Post 请求方法
	Header http.Header // 请求头部
	Data   url.Values  // post form data 表单字段
	BData  []byte      // binary data 文件上传二进制流
	Wait   int         // sleep time 等待时间
}

type Spider struct {
	*SpiderConfig
	Preurl        string         //pre url 上一次访问的URL
	Raw           []byte         //抓取的二进制流
	UrlStatusCode int            //响应状态吗 如 404
	Client        *http.Client   //客户端
	Fetchtimes    int            //抓取次数
	Errortimes    int            //失败次数
	Ipstring      string         //代理IP地址，没有代理默认localhost
	Request       *http.Request  //方便外部调试
	Response      *http.Response //
	mux           sync.RWMutex   //锁，并发抓取
}

func NewSpider(ipstring interface{}) (*Spider, error) {
	fmt.Println("Creat a new spider")
	chsp := new(Spider) //创建一个对象
	chsp.SpiderConfig = new(SpiderConfig)
	chsp.Header = http.Header{}
	chsp.Data = url.Values{}
	chsp.BData = []byte{}
	if ipstring != nil {
		client, err := NewProxyClient(ipstring.(string))
		chsp.Client = client
		chsp.Ipstring = ipstring.(string)
		return chsp, err
	} else {
		client, err := NewClient()
		chsp.Client = client
		chsp.Ipstring = "localhost"
		return chsp, err
	}

}

//设置URL
func (config *SpiderConfig) SetUrl(Url string) *SpiderConfig {
	config.Url = Url
	temp := strings.Split(Url, "//")
	if len(temp) >= 2 {
		config.SetHost(strings.Split(temp[1], "/")[0])
	}
	return config
}

func (config *SpiderConfig) SetRefer(refer string) *SpiderConfig{
	config.Header.Set("Referer", refer)
	return config
}

//设置host
func (config *SpiderConfig) SetHost(Host string) *SpiderConfig {
	config.Header.Set("Host", Host)
	return config
}

func (sp *Spider) Go() (data []byte, e error) {
	switch strings.ToUpper(sp.Method) {
	
		//		return sp.Post()
		//	case POSTJSON:
		//		return sp.PostJSON()
		//	case POSTXML:
		//		return sp.PostXML()
		//	case POSTFILE:
		//		return sp.PostFILE()
		//	case PUT:
		//		return sp.Put()
		//	case PUTJSON:
		//		return sp.PutJSON()
		//	case PUTXML:
		//		return sp.PutXML()
		//	case PUTFILE:
		//		return sp.PutFILE()
		//	case DELETE:
		//		return sp.Delete()
		//	case OTHER:
	case POST:		
		return []byte(""), errors.New("return other")
	default:
		return sp.Get()
	}
}

func (sp *Spider) Get() (body []byte, e error) {
	sp.mux.Lock()
	defer sp.mux.Unlock()

	//wait second not
	base.Wait(sp.Wait)

	//	logger.Debug("Get URL : " + sp.Url)

	//a new request
	request, _ := http.NewRequest("GET", sp.Url, nil)

	request.Header = CloneHeader(sp.Header)

	sp.Request = request
	//	logger.Debugf("----request header----", request)

	//start request
	if sp.Client == nil {
		// default client
		sp.Client = Client
	}
	response, err := sp.Client.Do(request)
	if err != nil {
		sp.Errortimes++
		return nil, err
	}

	if response != nil {
		defer response.Body.Close()
	}

	//debug
	sp.UrlStatusCode = response.StatusCode
	//设置新Cookie
	//Cookieb = MergeCookie(Cookieb, response.Cookies())

	//返回内容 return bytes
	body, e = ioutil.ReadAll(response.Body)
	sp.Raw = body

	sp.Fetchtimes++

	sp.Preurl = sp.Url

	sp.Response = response
	return

}
