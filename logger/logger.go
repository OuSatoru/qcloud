package logger

import (
	"fmt"
	"time"
)

func Log(format string, a interface{}) {
	fmt.Printf("%s "+format, time.Now().Format("2006-01-02 15:04:05"), a)
}
