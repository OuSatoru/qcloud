package runner

import (
	"io/ioutil"
	"log"
	"crypto/md5"
	"encoding/hex"
	"time"
	"github.com/OuSatoru/qcloud/wechat"
)

//only watch a list of small files
type FileModified struct {
	FileName string
	Modified bool
}

func fileWatch(fileName string, modified chan FileModified)  {
	for  {
		fileSum := checkSum(fileName)
		time.Sleep(2 * time.Second)
		isModified := fileModified(fileName, fileSum)
		modifiedStruct := FileModified{FileName:fileName, Modified:isModified}
		modified <- modifiedStruct
	}
}

func FileExecute(fileName string, modified chan FileModified)  {
	go fileWatch(fileName, modified)
	for {
		switch mod := <- modified {
		case mod.Modified == true:
			switch mod.FileName {
			case wechat.CREATEMENU:

			}
		case mod.Modified == false:

		}
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
