package wechat

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Parse args from the message that post from wechat, so r REQUEST.
func PsText(r *http.Request) *TextMsg {
	resp, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	body := &TextMsg{}
	xml.Unmarshal(resp, body)
	fmt.Println("Text Content:", body.Content)
	return body
}

// make xml from args then post TO wechat, use fmt.Fprintf(w)
func MkText(fromUserName, toUserName, content string) ([]byte, error) {
	textReply := &TextReply{}
	textReply.FromUserName = fromUserName
	textReply.ToUserName = toUserName
	textReply.Content = content
	textReply.MsgType = "text"
	textReply.CreateTime = time.Duration(time.Now().Unix())
	return xml.MarshalIndent(textReply, " ", "  ")
}

type Common struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	//MsgId        int64
}

type Typ struct {
	MsgType string
}

type TextReply struct {
	Common
	Content string
}

type TextMsg struct {
	TextReply
	MsgId int64
}

type ImageMsg struct {
	Common
	PictureUrl string
	MediaId    string
}

type VoiceMsg struct {
	Common
	MediaId     string
	Format      string
	Recognition string
}

type VideoMsg struct { //Also short video
	Common
	MediaId      string
	ThumbMediaId string
}

type PositionMsg struct {
	Common
	Location_X string
	Location_Y string
	Scale      string
	Label      string
}

type LinkMsg struct {
	Title       string
	Description string
	Url         string
}
