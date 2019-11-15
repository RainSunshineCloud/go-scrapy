package baidu_client

import (
	"net/http"
	"scrapy/manager"
)

type Request struct {
	url string
	method string
	params map[string]string
}

func (this *Request) Get() (*http.Request,error) {
	req,err :=  http.NewRequest(this.method,this.url,nil)
	if err != nil {
		return nil,err;
	}

	return req,nil;
}

func NewRequest (url string,params map[string]string) manager.RequestInterface {
	return &Request{url,"get",params}
}
