package runner

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"github.com/OuSatoru/qcloud/wechat"
	"io/ioutil"
	"log"
	"time"
)

//only watch a list of small files
type FileModified struct {
	FileName string
	Modified bool
}

func fileWatch(fileName string, modified chan FileModified) {
	for {
		fileSum := checkSum(fileName)
		time.Sleep(2 * time.Second)
		isModified := fileModified(fileName, fileSum)
		modifiedStruct := FileModified{FileName: fileName, Modified: isModified}
		modified <- modifiedStruct
	}
}

func FileExecute(db *sql.DB, fileName string, modified chan FileModified) {
	go fileWatch(fileName, modified)
	for {
		mod := <-modified
		switch {
		case mod.Modified == true:
			switch mod.FileName {
			case wechat.CREATEMENU:
				wechat.CreateMenu(db)
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
