package main

import (
	"fmt"
	"time"
)

const (
	constCommRawDataChanMaxcount = 512
	constChanMaxcount            = 256
)

var (
	chanCommIn  = make(chan []byte, constCommRawDataChanMaxcount) //串口入队列
	chanCommOut = make(chan []byte, constChanMaxcount)            //串口出队列
	chanLogData = make(chan []byte, constChanMaxcount)            //日志输出
)

func main() {
	//读取配置文件
	InitConfig()

	//开始写日志
	StartLog()

	//打开串口
	StartSerial()

	fmt.Println("Process...")
	for {
		time.Sleep(5 * time.Second)
	}
}
