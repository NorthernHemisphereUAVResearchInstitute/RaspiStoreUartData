package main

import (
	"github.com/golang/protobuf/proto"
    "indem_ar"
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

	/*msg_test:=&indem_ar.Pose{}

	msg_test.Time=123456787890
	msg_test.X=0.1
	msg_test.Y=0.2
	msg_test.Z=0.3
	msg_test.QuartX=0.001
	msg_test.QuartY=0.002
	msg_test.QuartZ=0.003
	msg_test.QuartW=0.004

	in_data, err := proto.Marshal(msg_test)
    if err != nil {
        fmt.Println("Marshaling error: ", err)
        //s.Exit(1)
    }
    fmt.Println(fmt.Sprintf("%x\r\n",in_data))*/
	msg_encoding:= &indem_ar.Pose{}
	//MavlinkMsgAnalyze(buf[2], buf[3:])
	//err = proto.Unmarshal(in_data, msg_encoding)
	err := proto.Unmarshal(buf[3:51], msg_encoding)
	//err := proto.Unmarshal(buf[2:len(buf)-2], msg_encoding)
    if err != nil {
        fmt.Println("Unmarshaling error: ", err)
    }
    //fmt.Println(len(buf))
	//fmt.Println(fmt.Sprintf("%x\r\n", buf[6:len(buf)-2]))
	//fmt.Println(msg_encoding.GetTime())
	fmt.Printf("%.3f,%.3f,%.3f,%.3f,%.3f,%.3f,%.3f\r\n",msg_encoding.GetX(),msg_encoding.GetY(),msg_encoding.GetZ(),msg_encoding.GetQuartX(),msg_encoding.GetQuartY(),msg_encoding.GetQuartZ(),msg_encoding.GetQuartW())
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
