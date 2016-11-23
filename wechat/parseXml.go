package wechat

type Common struct {
	ToUserName   string
	FromUserName string
	CreateTime   int
	MsgType      string
	MsgId        int64
}

type Typ struct {
	MsgType string
}

type TextMsg struct {
	Common
	Content string
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
