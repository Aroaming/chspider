package spider

import (
	"fmt"
	"net/http"
	"os"
)

const (
	// 暂停时间 default wait time
	WaitTime = 5

	// HTTP方法
	POST     = "POST"
	POSTJSON = "POSTJSON"
	POSTXML  = "POSTXML"
	POSTFILE = "POSTFILE"

	// 实现了!
	PUT     = "PUT"
	PUTJSON = "PUTJSON"
	PUTXML  = "PUTXML"
	PUTFILE = "PUTFILE"

	DELETE = "DELETE"
	GET    = "GET"
	OTHER  = "OTHER"

	CRITICAL = "CRITICAL"
	ERROR    = "ERROR"
	WARNING  = "WARNING"
	NOTICE   = "NOTICE"
	INFO     = "INFO"
	DEBUG    = "DEBUG"

	HTTPFORMContentType = "application/x-www-form-urlencoded"
	HTTPJSONContentType = "application/json"
	HTTPXMLContentType  = "text/xml"
	HTTPFILEContentType = "multipart/form-data"
)

//clone a header
func CloneHeader(h map[string][]string) map[string][]string {
	if h == nil || len(h) == 0 {
		return map[string][]string{}
	}
	return CopyM(h)
}

func CopyM(h http.Header) http.Header {
	if h == nil || len(h) == 0 {
		return h
	}
	h2 := make(http.Header, len(h))
	for k, vv := range h {
		v2 := make([]string, len(vv))
		copy(v2, vv)
		h2[k] = v2
	}
	return h2
}

func MkdirVendor(pagnum int) error {
	for i := 1; i <= pagnum; i++ {
		dir := fmt.Sprintf("./pic/%d", i)
		err := os.MkdirAll(dir, 0777)
		if err != nil {
			return err
		}
	}
	return nil
}
