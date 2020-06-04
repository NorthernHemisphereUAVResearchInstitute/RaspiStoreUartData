package main

import (
	"fmt"
)

const (
	MavlinkSyncFirst    = byte(0xF0)              //mavlink packet 第一个标志字节
	MavlinkIndexDataLen = 7                       //数据长度索引
	MavlinkIndexMsgId   = 6                       //消息id索引
	MavlinkIndexCrc16   = 2                      //crc16索引
	MavlinkPackHeadLen  = 8                       //包头长度
	MavlinkPackLenMin   = 10                      //最小数据包长度
	MavlinkLenMax       = 32 					  //最大数据包长度
	IMUMsgID 			= 0x01
	IMUMsgLEN        	= 22
)

//检查数据长度是否正确
//检查数据长度是否正确
func CheckPackDataLen(msgid byte,datalen byte) (bRet bool) {

	if msgid == IMUMsgID {
		if datalen == IMUMsgLEN {
			bRet = true
		} else {
			bRet = false
		}
	}

	return bRet
}

//是否是有效的头部
func IsValidPackHead(packHead []byte) (bOk bool) {
	if len(packHead) >= MavlinkPackHeadLen {
		if (packHead[0] == MavlinkSyncFirst) && (packHead[1] == MavlinkSyncFirst){ //前导标示符相同
			if CheckPackDataLen(packHead[MavlinkIndexMsgId],packHead[MavlinkIndexDataLen]) { // 数据包长度
				bOk = true
			} else {
				bOk = false
				fmt.Println("----------------->invalid header")
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
					newcrc16 :=CRC16(newbuf[2:packlen-MavlinkIndexCrc16],packlen-MavlinkIndexCrc16-2)
					packcrc := Bytes2Uint16(newbuf[packlen-MavlinkIndexCrc16 : packlen])
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
