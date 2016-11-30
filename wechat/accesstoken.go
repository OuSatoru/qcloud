package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AccessToken struct {
	AppId     string
	AppSecret string
}

func (at *AccessToken) fetch() (string, error) {
	rtn, err := get(fmt.Sprintf("%stoken?grant_type=client_credential&appid=%s&secret=%s", cgibin, at.AppId, at.AppSecret))
	if err != nil {
		return "", err
	}
	return rtn.AccessToken, nil
}

type at_response struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`

	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func get(url string) (*at_response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var rtn at_response
	if err := json.Unmarshal(data, &rtn); err != nil {
		return nil, err
	}
	if rtn.ErrCode != 0 {
		return nil, errors.New(fmt.Sprintf("%d %s", rtn.ErrCode, rtn.ErrMsg))
	}

	return &rtn, nil
}
