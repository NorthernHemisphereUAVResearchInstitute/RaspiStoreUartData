package main

import (
	"fmt"
)

const (
	MavlinkSyncFirst    = byte(0xf0)              //mavlink packet 第一个标志字节
	MavlinkSyncEnd    	= byte(0xfa)              //mavlink packet 最后个标志字节
	MavlinkIndexDataLen = 1                       //数据长度索引
	MavlinkIndexPkg 	= 2
	MavlinkIndexMSGID 	= 3
	MavlinkIndexData    = 4
	MavlinkIndexCrc16   = -3                      //crc16索引
	MavlinkPackHeadLen  = 4                       //包头长度
	MavlinkPackLenMin   = 7                       //最小数据包长度
	MavlinkLenMax       = MavlinkPackLenMin + 28 //最大数据包长度
	MavlinkLenPayLOADLen=28
	IMU_RAW_PAYLOAD_LEN =28
)

//检查数据长度是否正确
func CheckPackDataLen(datalen byte) (bRet bool) {

	if datalen == MavlinkLenPayLOADLen {
		bRet = true
	} else {
		bRet = false
	}
	return bRet
}

//是否是有效的头部
func IsValidPackHead(packHead []byte) (bOk bool) {
	if len(packHead) >= MavlinkPackHeadLen {
		if packHead[0] == MavlinkSyncFirst && packHead[MavlinkLenMax-1] == MavlinkSyncEnd { //前导标示符相同
			if CheckPackDataLen(packHead[MavlinkIndexDataLen]) { // 数据包长度
				bOk = true
			} else {
				bOk = false
				fmt.Println("----------------->invalid head")
			}
		}
	}
	return bOk
}

func GetMavlinkDataLen(packHead []byte) int {
	datalen := packHead[MavlinkIndexDataLen]
	return int(datalen)
}

//解析处理
func dealMavlinkCommInData(buf []byte, lenbuf int) (bFind bool, newIndex int, thisPackLen int) {
	if lenbuf >= MavlinkPackLenMin {
		endbuflen := lenbuf - MavlinkPackLenMin
		i := 0
		for i < endbuflen {
			if IsValidPackHead(buf[i:]) { //包头正确
				newbuf := buf[i:]
				datalen := GetMavlinkDataLen(newbuf)
				packlen := MavlinkPackLenMin + datalen
				if lenbuf-i >= packlen { //数据包大小够
					newcrc16 := CRC16(newbuf[MavlinkIndexData : packlen+MavlinkIndexCrc16], MavlinkLenPayLOADLen)
					packcrc := Bytes2Uint16(newbuf[packlen+MavlinkIndexCrc16 : packlen+MavlinkIndexCrc16+2])
					if newcrc16 == packcrc { //校验和正确
						bFind = true
						newIndex = i
						thisPackLen = packlen
						break
					} else {
						fmt.Println(fmt.Sprintf("crc err:%x,%x,%x", newcrc16, packcrc, newbuf[:packlen]))
					}
				} else { //数据包长度不够
					break
				}
			}
			i++
		}
	}

	return bFind, newIndex, thisPackLen
}
