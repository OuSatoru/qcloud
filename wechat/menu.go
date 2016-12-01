package wechat

import "encoding/json"

type menu struct {
	Button []buttons `json:"button"`
}

type buttons struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Key string `json:"key"`
	Url string `json:"url"`
	MediaId string `json:"media_id"`
	SubButton []buttons `json:"sub_button"`
}

type subButtons struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Key string `json:"key"`
	Url string `json:"url"`
	MediaId string `json:"media_id"`
}

func MenuToJson(m menu) string {
	return string(json.MarshalIndent(m, "", "  "))
}

// test: var m = menu{Button: []buttons{{Type: "ttype", Name: "nname", Key: "kkey",
// SubButton: []buttons{{Type: "tttype", Name: "nnname", Url: "http://www.baidu.com", SubButton: []buttons{}}}}}}