package wechat

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
	"strings"
)

// http handle func for /wx
func Handle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if !checkUrl(w, r) {
		log.Println("Fail")
		log.Println(r.Form["token"], r.Form["timestamp"], r.Form["nonce"])
		return
	}
	if r.Method == "POST" {
		BigContent := PsBig(r)
		if BigContent != nil {
			switch BigContent.MsgType {
			case "text":
				log.Printf("msg:| %s |, user:| %s |", BigContent.Content, BigContent.FromUserName)
				textReply, err := MkText(BigContent.ToUserName, BigContent.FromUserName, reverseStr(BigContent.Content))
				if err != nil {
					log.Println(err)
					return
				}
				fmt.Fprintf(w, string(textReply))
			default:
				fmt.Fprint(w, "success")
			}

		}
	}
	log.Println("Succeed")
}

func sha(timestamp, nonce string) string {
	ls := []string{token, timestamp, nonce}
	sort.Strings(ls)
	s := sha1.New()
	io.WriteString(s, strings.Join(ls, ""))
	return fmt.Sprintf("%x", s.Sum(nil))
}

func checkUrl(w http.ResponseWriter, r *http.Request) bool {
	timestamp := strings.Join(r.Form["timestamp"], "")
	nonce := strings.Join(r.Form["nonce"], "")
	sigSha1ed := sha(timestamp, nonce)
	signature := strings.Join(r.Form["signature"], "")
	if sigSha1ed != signature {
		return false
	}
	echostr := strings.Join(r.Form["echostr"], "")
	fmt.Fprintf(w, echostr)
	return true
}
