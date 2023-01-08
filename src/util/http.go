package util

import (
	"bytes"

	"io/ioutil"
	"net/http"
)

func PostJson(url string, obj string) *http.Response {

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(obj)))
	Checke(err, "http.NewRequest failed")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	Checke(err, "client.Do failed")

	return resp
}

func ReadResp(res *http.Response) string {

	body, err := ioutil.ReadAll(res.Body)
	Checke(err, "ioutil.ReadAll failed")

	return string(body)
}
