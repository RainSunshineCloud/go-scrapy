package baidu_client

type Response struct {
	data []byte
}

func (this *Response) ToDo () error {
	//time.Sleep(10*time.Second)
	return nil;
}

func NewResponse (data_bypes []byte) *Response {
	return &Response{
		data_bypes,
	}
}