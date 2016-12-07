package wechat

import (
	"io/ioutil"
	"log"
	"crypto/md5"
	"encoding/hex"
)

//only watch a list of small files
func FileWatcher(fileName string)  {
	
}

func fileModified(fileName string, oriSum string) bool {
	return oriSum == checkSum(fileName)
}

func checkSum(fileName string) string {
	//md5
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Println(err)
	}
	h := md5.New()
	h.Write(content)
	return hex.EncodeToString(h.Sum(nil))
}
