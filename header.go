/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package req

import "encoding/json"

// Header represents http request header
type Header map[string]string

var MacChromeHeader = Header{
	"Accept":          "*/*",
	"User-Agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36",
	"Accept-Encoding": "gzip, deflate",
}

var MacFirefoxHeader = Header{
	"Accept":          "*/*",
	"User-Agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/115.0",
	"Accept-Encoding": "gzip, deflate",
}

var CurlHeader = Header{
	"Accept":     "*/*",
	"User-Agent": "curl/7.87.0",
}

var ReqHeader = Header{
	"Accept":     "*/*",
	"User-Agent": "Go Req",
}

func (h Header) Clone() Header {
	if h == nil {
		return nil
	}
	hh := Header{}
	for k, v := range h {
		hh[k] = v
	}
	return hh
}

// ParseStruct parse struct into header
func ParseStruct(h Header, v interface{}) Header {
	data, err := json.Marshal(v)
	if err != nil {
		return h
	}

	_ = json.Unmarshal(data, &h)
	return h
}

// HeaderFromStruct init header from struct
func HeaderFromStruct(v interface{}) Header {

	var header Header
	header = ParseStruct(header, v)
	return header
}

type ReservedHeader map[string]string
