package wechat

import "fmt"

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
