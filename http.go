package util

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"
)

func HttpGetJson(url string, recvData interface{}, timeout time.Duration) (resData interface{}, err error) {
	defaultTimeout := http.DefaultClient.Timeout
	defer func() {
		http.DefaultClient.Timeout = defaultTimeout
	}()
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	if timeout != 0 {
		http.DefaultClient.Timeout = timeout
	} else {
		http.DefaultClient.Timeout = time.Second * 15
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal([]byte(resBody), recvData)
	return recvData, nil
}

func HttpPostJson(url string, sendData, recvData interface{}, timeout time.Duration) (interface{}, error) {
	defaultTimeout := http.DefaultClient.Timeout
	defer func() {
		http.DefaultClient.Timeout = defaultTimeout
	}()
	if timeout != 0 {
		http.DefaultClient.Timeout = timeout
	} else {
		http.DefaultClient.Timeout = time.Second * 15
	}
	reqData, _ := json.Marshal(sendData)
	body := strings.NewReader(string(reqData))
	res, err := http.DefaultClient.Post(url, "application/json", body)
	if err != nil {
		return nil, err
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal([]byte(resBody), recvData)
	return recvData, nil
}
