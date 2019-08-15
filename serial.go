package main

import (
	"./serial"
	"fmt"
	"io"
	"time"
)

//------------------------------------------------------------------------------------
//检测串口是否存在
func IsSerialPortExist(comm string, baud int) {
	c := &serial.Config{Name: comm, Baud: baud}

	s, err := serial.OpenPort(c)
	for err != nil {
		fmt.Println("打开串口失败:", comm)
		time.Sleep(3 * time.Second) //三秒尝试一次
		s, err = serial.OpenPort(c)
	}
	fmt.Println("SerailPort:", comm, "存在!")
	s.Close()
}

//------------------------------------------------------------------------------------
//打开串口
func OpoenSerialPort(comm string, baud int) io.ReadWriteCloser {
	c := &serial.Config{Name: comm, Baud: baud}

	s, err := serial.OpenPort(c)
	for err != nil {
		fmt.Println("打开串口失败:", comm, "三秒后重试")
		time.Sleep(3 * time.Second) //三秒尝试一次
		s, err = serial.OpenPort(c)
	}
	fmt.Println("SerailPort:", comm, "打开成功")
	return s
}

////////////////////////////////////////////////////////////////////
//接收数据
func receiveCom(s io.ReadWriteCloser) {
	cachbuf := make([]byte, 512*2048)
	cachIndex := 0
	cachCount := 2000
	cachLen := 512
	//清理缓存
	for i := 0; i < 1; i++ { //清理缓存
		n, err := s.Read(cachbuf)
		if err != nil {
			fmt.Println("read 1", err, n)
		} else {
			fmt.Println("read:", n)
		}
	}

	for {
		buf := cachbuf[cachIndex*cachLen : (cachIndex+1)*cachLen]
		n, err := s.Read(buf)
		if err != nil {
			if n != 0 {
				fmt.Println("read2", err, n)
			} else {
				time.Sleep(3 * time.Millisecond)
			}

			continue
		}
		select {
		case chanCommIn <- buf[:n]:
		default:
			fmt.Println("chanCommIn 溢出")
		}
		cachIndex++
		if cachIndex >= cachCount {
			cachIndex = 0
		}
	}
}

//解析处理
func dealCommInDataByCrc() {
	maxbuflen := 512 * 2048
	buf := make([]byte, maxbuflen*2)
	lenbuf := 0
	fmt.Println("dealCommInDataByCrc")
	for {
		select {
		case data := <-chanCommIn:
			if (lenbuf + len(data)) >= maxbuflen/2 { //越界
				lenbuf = 0
			} else {
				copy(buf[lenbuf:], data)
				lenbuf += len(data)
				//这里应该增加一个判断是否继续处理的函数，看是否是包头，不是包头一直处理,
				headOffset := 0
				for {
					bFind, newIndex, thisPackLen := dealMavlinkCommInData(buf[headOffset:], lenbuf-headOffset)
					if bFind {
						//有数据包，处理
						newbuf := buf[headOffset+newIndex:]
						DealMavlinkMsgRecv(newbuf[:])
					} else { //没有数据包了，退出
						break
					}
					headOffset += (newIndex + thisPackLen) //新的偏移起点
				}

				if headOffset != 0 {
					if headOffset != lenbuf { //需要搬移数据
						copy(buf, buf[headOffset:lenbuf])
						lenbuf = lenbuf - headOffset
					} else {
						lenbuf = 0
					}
				}

			}

		}
	}
}


//------------------------------------------------------------------------------
//串口写数据
func ComSend(s io.ReadWriteCloser) {
	buf := make([]byte, 1)
	bSaveLog := ForwarderJson.SaveLog
	for {
		select {
		case buf = <-chanCommOut:
			s.Write(buf)
		}

		if bSaveLog {
			select {
			case chanLogData <- buf:
			default:
				fmt.Println("chanLogData ComSend 溢出 ")
			}
		}

	}
}

func StartSerial() {
	s := OpoenSerialPort(ForwarderJson.Port, ForwarderJson.Baud)
	go ComSend(s)
	time.Sleep(time.Second)
	go  pack_and_send()
	/*go dealCommInDataByCrc()
	time.Sleep(time.Second)
	go receiveCom(s)*/
}
