package my_config

import (
	"fmt"
	config "github.com/RainSunshineCloud/go-iniconfig"
	"strings"
)

type myConf struct {
	*config.Config
}

var my_conf = &myConf{}

func init () {
	new();
}
func new() {
	my_conf = &myConf{}
	my_conf.Config = config.New("F:/test/scrapy/my_config/config.ini",false)
	if !my_conf.Load() {
		panic(my_conf.LastErr())
	}
}

type request struct {
	RequestUrl string
	ClientName string
	Params map[string]string
}

func (this *myConf) GetRequest(key string) request {
	res := request{}
	tmp := this.Get(fmt.Sprintf("%s.%s",key,"request_url"))
	if val,ok := tmp.(string); ok {
		res.RequestUrl = val;
	}

	tmp = this.Get(fmt.Sprintf("%s.%s",key,"client_name"))
	if val,ok := tmp.(string); ok {
		res.ClientName = val;
	}

	tmp = this.Get(fmt.Sprintf("%s.%s",key,"request_params"))
	if val,ok := tmp.([]string); ok {
		for _,v := range val {
			tmp := strings.Split(v,">")
			res.Params[tmp[0]] = tmp[1];
		}

	}

	return res;
}

func (this *myConf) GetRequestsName () []string {
	tmp := this.Get("common.requests")
	if val,ok := tmp.([]string); ok {
		return val;
	}
	panic(this.LastErr())
}

func (this *myConf) GetLogDir (key string) string {
	tmp := this.Get("common.log_dir")
	if val,ok := tmp.(string); ok {
		return val;
	}
	panic(this.LastErr())
}

func Get () *myConf {
	return my_conf
}


