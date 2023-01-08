package httpcli

import (
	u "cbotcli/src/util"
	"net/http"
)

type Httpcli struct{}

func NewHttpcli() *Httpcli {
	obj := new(Httpcli)
	return obj
}

func (h *Httpcli) Get(url string) *http.Response {

	resp, err := http.Get(url)
	u.Checke(err, "http.GET failed")

	return resp
}

func (h *Httpcli) Post(url string, jsonStr string) *http.Response {

	return u.PostJson(url, jsonStr)
}
