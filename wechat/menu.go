package wechat

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"github.com/OuSatoru/qcloud/runner"
	"net/http"
	"strings"
	"log"
)

type menu struct {
	Button []buttons `json:"button" yaml:"button"`
}

type buttons struct {
	Type      string    `json:"type,omitempty" yaml:"type,omitempty"` //omit only if have sub_button
	Name      string    `json:"name" yaml:"name"`
	Key       string    `json:"key,omitempty" yaml:"key,omitempty"`           //click must
	Url       string    `json:"url,omitempty" yaml:"url,omitempty"`           //view must
	MediaId   string    `json:"media_id,omitempty" yaml:"media_id,omitempty"` //media_id view_limited must
	SubButton []buttons `json:"sub_button,omitempty" yaml:"sub_button,omitempty"`
}

type subButtons struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Key     string `json:"key"`
	Url     string `json:"url"`
	MediaId string `json:"media_id"`
}

func MenuToJson(m menu) string {
	v, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return ""
	}
	return string(v)
}

func JsonToMenu(j []byte) (m menu) {
	if err := json.Unmarshal(j, &m); err != nil {
		return nil
	}
	return m
}

func YamlToMenu(s []byte) (m menu) {
	if err := yaml.Unmarshal(s, &m); err != nil {
		return nil
	}
	return m
}

func MenuToYaml(m menu) string {
	v, err := yaml.Marshal(m)
	if err != nil {
		return ""
	}
	return string(v)
}

func CreateMenu(db runner.DbLogin)  {
	content, err := ioutil.ReadFile(CREATEMENU)
	if err != nil {
		log.Println(err)
		return
	}
	menuJson := MenuToJson(YamlToMenu(content))
	resp, err := http.Post(cgibin+"menu/create?access_token="+ NowAccessToken(db),"application/json",strings.NewReader(menuJson))
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(body))
}

func SearchMenu(db runner.DbLogin) string {
	resp, err := http.Get(cgibin+"menu/get?access_token="+ NowAccessToken(db))
	if err != nil {
		log.Println(err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return ""
	}
	return MenuToYaml(JsonToMenu(body))
}

func DeleteMenu(db runner.DbLogin)  {
	resp, err := http.Get(cgibin+"menu/delete?access_token="+ NowAccessToken(db))
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(body))
}

// test: var m = menu{Button: []buttons{{Type: "ttype", Name: "nname", Key: "kkey",
// SubButton: []buttons{{Type: "tttype", Name: "nnname", Url: "http://www.baidu.com", SubButton: []buttons{}}}}}}
