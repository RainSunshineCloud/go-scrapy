package baidu_client

import (
	"errors"
	"io/ioutil"
	"net/http"
	"scrapy/manager"
	"sync"
	"time"
)

type BaiduClient struct {
	sync.Mutex
	isLogin bool;
	http.Client
}



func  NewClient () *BaiduClient {
	return &BaiduClient{
		isLogin: true,
		Client:  http.Client{},
	}
}

func (this *BaiduClient) Init () (manager.ClientInterface,error) {

	if this.isLogin {
		return this,nil;
	}
	//重新登录的时间点
	defer time.Sleep(4 * time.Second)
	return this,errors.New("登录失败");
}

func (this *BaiduClient) Request (requests manager.RequestInterface) (manager.ResponseInterface,error) {
	req,err := requests.Get()
	if err != nil  {
		return nil,err;
	}
	this.isLogin = false;
	data,err := this.Do(req)
	if err != nil  {
		return nil,err;
	}

	if  data.StatusCode != 200 {
		return nil,errors.New(data.Status)
	}

	data_bytes,err := ioutil.ReadAll(data.Body)
	if err != nil {
		return nil,err;
	}

	if  data_bytes == nil || len(data_bytes) == 0 {
		return nil,errors.New("数据为空");
	}
	tmp :=  NewResponse(data_bytes);
	return tmp,nil;
}






