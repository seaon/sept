package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

var logPath string

func openLog(fileName string) io.Writer {
	fileInfo, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		log.Fatal("文件不存在")
	}
	if err != nil {
		log.Fatal(fmt.Sprintf("%v", err))
	}

	if fileInfo.Size() > 1024*1024*5 {
		unixTime := strconv.FormatInt(time.Now().Unix(), 10)
		_ = os.Rename(logPath+fileName, logPath+unixTime+fileName)
	}
	file, err := os.OpenFile("file.log", os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
	return file
}
