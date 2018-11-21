package http

import (
	"io/ioutil"
	goHttp "net/http"
)

type Request struct {
	Request *goHttp.Request
}

func (request *Request) GetBodyString() string  {
	content, _ := ioutil.ReadAll(request.Request.Body);
	return string(content);
}
