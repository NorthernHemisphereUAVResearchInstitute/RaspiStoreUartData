package main

import (
	"fmt"
)

var (
	VelPosStruct = StructVelPos{}
)

type StructVelPos struct {
	vx float32
	vy float32
	vz float32
	px float32
	py float32
	pz float32
}

func MavAnalyzeVelPos(buf []byte) (velpos StructVelPos) {
	index := 0

	velpos.vx = Bytes2float32(buf[index : index+4])
	index += 4
	velpos.vy = Bytes2float32(buf[index : index+4])
	index += 4
	velpos.vz = Bytes2float32(buf[index : index+4])
	index += 4
	velpos.px = Bytes2float32(buf[index : index+4])
	index += 4
	velpos.py = Bytes2float32(buf[index : index+4])
	index += 4
	velpos.pz = Bytes2float32(buf[index : index+4])
	index += 4

	return velpos
}

func MavlinkMsgAnalyze(msgid byte, msgdatabuf []byte) {
	switch msgid {
	case 1:
		VelPosStruct = MavAnalyzeVelPos(msgdatabuf)
		fmt.Printf("%.3f,%.3f,%.3f，%.3f，%.3f，%.3f\n", VelPosStruct.vx, VelPosStruct.vy, VelPosStruct.vz, VelPosStruct.px, VelPosStruct.py, VelPosStruct.pz)
	default:

	}
	SaveLog([]byte(fmt.Sprintf("vx:%.6f,vy:%.6f,vz:%.6f,px:%.6f,py:%.6f,pz:%.6f\r\n",
		VelPosStruct.vx, VelPosStruct.vy, VelPosStruct.vz, VelPosStruct.px, VelPosStruct.py, VelPosStruct.pz)))
}

func DealMavlinkMsgRecv(buf []byte) {
	//解析
	MavlinkMsgAnalyze(buf[2], buf[3:])
}

func SaveLog(logbuf []byte) {
	if ForwarderJson.SaveLog {
		select {
		case chanLogData <- logbuf:
		default:
			fmt.Println("chanLogData 溢出 ")
		}
	}
}
