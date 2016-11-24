package wechat

import (
	"crypto/sha1"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
	"log"
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
		textContent := PsText(r)
		if textContent != nil {
			log.Printf("msg:| %s |, user:| %s |", textContent.Content, textContent.FromUserName)
			textReply, err := MkText(textContent.ToUserName, textContent.FromUserName, reverseStr(textContent.Content))
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Fprintf(w, string(textReply))
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
