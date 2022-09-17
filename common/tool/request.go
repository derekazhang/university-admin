package tool

import "github.com/parnurzeal/gorequest"

var Request *gorequest.SuperAgent

func init() {
	Request = gorequest.New()
}

func GetRequest() *gorequest.SuperAgent {
	return gorequest.New()
}
