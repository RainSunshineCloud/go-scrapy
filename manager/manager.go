package manager

import (
	"errors"
	"log"
	"net/http"
	"sync"
)

type manager struct {
	clients *sync.Map
	log LogInterface
}

type LogInterface interface {
	Error() *log.Logger
	Info() *log.Logger
}

type RequestInterface interface {
	Get() (*http.Request,error)
}

type ResponseInterface interface {
	ToDo() error
}

type ClientInterface interface {
	Init() (ClientInterface,error)
	Lock()
	Unlock()
	Request(requestInterface RequestInterface) (ResponseInterface,error)
}

//获取
func (this *manager) get (key string) (ClientInterface,error) {
	client,ok := this.clients.Load(key)
	if !ok {
		return nil,errors.New("未有该客户端")
	}
	val,ok := client.(ClientInterface);
	if  !ok {
		return nil,errors.New("未有该客户端")
	}
	val.Lock()
	defer func () {
		this.Set(key,val)
		val.Unlock()
	}();
	return val.Init()
}


//设置
func (this *manager) Set (key string,client ClientInterface)  {
	this.clients.Store(key,client)
}

func New () *manager {
	return &manager{
		clients:&sync.Map{},
	}
}

func (this *manager) Run (key string,request RequestInterface) {
	for {
		this.runFlow(key,request)
	}
}

func (this *manager) runFlow (key string,request RequestInterface) {
	//记录错误
	defer func () {
		if r := recover();r != nil {
			this.log.Error().Println(r)
		}
	}()

	//获取客户端
	data_client,err := this.get(key);
	if err != nil {
		this.log.Error().Println(err)
		return;
	}

	//跑请求
	res,err := data_client.Request(request);
	//设置请求体
	defer func (key string, data_client ClientInterface) {
		this.Set(key,data_client)
	}(key ,data_client)

	if err != nil {
		this.log.Error().Println(err)
		return;
	}

	//响应后要做的事
	if err := res.ToDo(); err != nil {
		this.log.Error().Println(err)
		return;
	}

}

func (this *manager) SetLog (my_log LogInterface) {
	this.log = my_log;
}
