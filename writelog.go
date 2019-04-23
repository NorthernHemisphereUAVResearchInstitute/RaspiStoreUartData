package main

import (
	"fmt"
	"os"
	"strings"
)

/////////////////////////////////////////////////////////////////////////////////////////////////////
//

func CreateTodayDir() (bRet bool) {
	dir := GetModulePath() + "log/"
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		bRet = true
	}
	return bRet
}

func CreateLogDataFile(suffix string) *os.File {
	dir := GetModulePath() + "log/"
	fileTime := GetNowDateTimeAsYYMMDDHHMISS()
	fileTime = strings.Replace(fileTime, "-", "", -1)
	fileTime = strings.Replace(fileTime, ":", "", -1)
	fileTime = strings.Replace(fileTime, " ", "", -1)
	fileName := dir + fileTime + suffix + ".log"
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(fileName)
		panic(err)
	}
	fmt.Println(fileName)
	return f
}

//把mavlinklog数据写入文件
func HandleWriteMavlinkLogData() {
	CreateTodayDir()
	f := CreateLogDataFile("fc")
	defer f.Close()
	for {
		select {
		case logData := <-chanLogData:
			f.WriteString(string(logData))
		}
	}
}

func StartLog() {
	if ForwarderJson.SaveLog {
		go HandleWriteMavlinkLogData()
	}
}
