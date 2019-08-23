package main

import (
	"fmt"
	"time"
	"math/rand"
)

var (
	InCommCachbuf   = make([]byte, 64*20480)
	InCommCachIndex = 0
	InCommCachCount = 20000
	InCommCachLen   = 64

	MavLinkSeq = byte(0)
)

func GetMavLinkSeq() byte {
	MavLinkSeq++
	return MavLinkSeq
}

//-------------------------------------------------------------------------------------------------------
func BufInchanCommOut(buf []byte, datalen int) {
	lenth := datalen
	outbuf := InCommCachbuf[InCommCachIndex*InCommCachLen : InCommCachIndex*InCommCachLen+lenth]
	copy(outbuf, buf[:lenth])
	fmt.Printf("data:%x\r\n",outbuf[:] )
	//入串口队列，准备串口发送
	select {
	case chanCommOut <- outbuf:
	default:
		fmt.Println("chanUdpIn:数据溢出")
	}

	InCommCachIndex++
	if InCommCachIndex >= InCommCachCount {
		InCommCachIndex = 0
	}
}

func pack_and_send() {

	buf := make([]byte, 20)
	
	for {
		GetMavLinkSeq()
		buf[19]=MavLinkSeq
		buf[0]=0xd0
		for i := 1; i < 19; i++ {
			buf[i]=byte(rand.Intn(10))
			//buf[i]=0x22
		}

		BufInchanCommOut(buf[:], 20)

		time.Sleep(20 * time.Millisecond)

	}
}