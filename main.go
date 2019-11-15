package main

import (
	"scrapy/baidu_client"
	"scrapy/manager"
	"scrapy/my_config"
	"scrapy/my_log"
)

var stop <-chan string;
var client_manager = manager.New()

func init () {
	client_manager.SetLog(my_log.Get())
}

func main () {
	registerClient()
	for _,request := range getRequests() {
		go client_manager.Run(request.clientName, request.req)
	}

	<-stop;
}

type Requests struct {
	req manager.RequestInterface
	clientName string
}

func getRequests () []Requests {
	res := []Requests{};
	for _,name := range my_config.Get().GetRequestsName() {
		if name == "" {
			continue;
		}
		obj := my_config.Get().GetRequest(name)
		req := getRequestsByName(obj.ClientName,obj.RequestUrl,obj.Params);
		if req != nil {
			res = append(res,Requests{
				req:  req,
				clientName: obj.ClientName,
			})
		}
	}
	if len(res) == 0 {
		panic("未有请求");
	}
	return res;
}

//获取request
func getRequestsByName (client_name string ,request_url string , request_params map[string]string) manager.RequestInterface {
	switch client_name {
		case "baidu":
			return baidu_client.NewRequest(request_url,request_params)
	}
	return nil;
}
//client
func registerClient () {
	tmp := baidu_client.NewClient();
	client_manager.Set("baidu",tmp)
}