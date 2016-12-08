package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	_ "github.com/lib/pq"
	"database/sql"
)

type AccessToken struct {
	AppId     string
	AppSecret string
}

func (at *AccessToken) fetch() (string, error) {
	rtn, err := getATResp(fmt.Sprintf("%stoken?grant_type=client_credential&appid=%s&secret=%s", cgibin, at.AppId, at.AppSecret))
	if err != nil {
		return "", err
	}
	return rtn.AccessToken, nil
}

func (at *AccessToken) FetchAtResp() (*At_response, error) {
	return getATResp(fmt.Sprintf("%stoken?grant_type=client_credential&appid=%s&secret=%s", cgibin, at.AppId, at.AppSecret))
}

type At_response struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`

	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// return access token response as a struct
func getATResp(url string) (*At_response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var rtn At_response
	if err := json.Unmarshal(data, &rtn); err != nil {
		return nil, err
	}
	if rtn.ErrCode != 0 {
		return nil, errors.New(fmt.Sprintf("%d %s", rtn.ErrCode, rtn.ErrMsg))
	}
	fmt.Println(&rtn)
	return &rtn, nil
}

func NowAccessToken(db *sql.DB) string {
	var at string
	row := db.QueryRow(`SELECT accesstoken
				FROM accesstoken
				WHERE id = (SELECT max(id)
					    FROM
					      (SELECT *
					       FROM accesstoken
					       WHERE accesstoken IS NOT NULL) a)`)
	row.Scan(&at)
	return at

}