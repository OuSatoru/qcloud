package grab

import (
	"net/http"
	//"fmt"
	"log"
	"regexp"
	"io/ioutil"
)

const YandeHead  = "https://yande.re/post/show/"

func Grab(ur string) string {
	htm, err := http.Get(ur)
	if err != nil {
		log.Fatalln(err)
		return ""
	}
	defer htm.Body.Close()
	iou, _ := ioutil.ReadAll(htm.Body)
	//from 5 to newest
	reg := regexp.MustCompile(`<img.*?id="image".*?src="(?P<first>.*?)".*?/>`)
	src := reg.FindStringSubmatch(string(iou))
	if len(src) != 0 {
		return src[1]
	} else {
		return ""
	}
}
