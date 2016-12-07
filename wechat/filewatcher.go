package wechat

import (
	"io/ioutil"
	"log"
	"crypto/md5"
	"encoding/hex"
	"time"
)

//only watch a list of small files
type FileModified struct {
	fileName string
	modified bool
}

func FileWatch(fileName string, modified chan FileModified)  {
	for  {
		fileSum := checkSum(fileName)
		time.Sleep(2 * time.Second)
		isModified := fileModified(fileName, fileSum)
		modifiedStruct := FileModified{fileName:fileName, modified:isModified}
		modified <- modifiedStruct
	}
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
