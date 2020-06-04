package main

import (
	"fmt"
)

var (
	IMURawDataStruct = StructIMU{}
)

type StructIMU struct {
	gx int16
	gy int16
	gz int16
	ax int16
	ay int16
	az int16
	mx int16
	my int16
	mz int16
	presure int16
	temperature int16

}

func MavAnalyzeIMU(buf []byte) (imu_data StructIMU) {
	index := 0

	imu_data.gx = Bytes2Int16(buf[index : index+2])
	index += 2
	imu_data.gy = Bytes2Int16(buf[index : index+2])
	index += 2
	imu_data.gz = Bytes2Int16(buf[index : index+2])
	index += 2
	imu_data.ax = Bytes2Int16(buf[index : index+2])
	index += 2
	imu_data.ay = Bytes2Int16(buf[index : index+2])
	index += 2
	imu_data.az = Bytes2Int16(buf[index : index+2])
	index += 2
	imu_data.mx = Bytes2Int16(buf[index : index+2])
	index += 2
	imu_data.my = Bytes2Int16(buf[index : index+2])
	index += 2
	imu_data.mz = Bytes2Int16(buf[index : index+2])
	index += 2
	imu_data.presure = Bytes2Int16(buf[index : index+2])
	index += 2
	imu_data.temperature = Bytes2Int16(buf[index : index+2])
	index += 2

	return imu_data
}

func MavlinkMsgAnalyze(msgid byte, msgdatabuf []byte) {
	switch msgid {
	case 1:
		IMURawDataStruct = MavAnalyzeIMU(msgdatabuf)
		fmt.Printf("%d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d\n",IMURawDataStruct.gx,IMURawDataStruct.gy,IMURawDataStruct.gz,
																IMURawDataStruct.ax,IMURawDataStruct.ay,IMURawDataStruct.az,
																IMURawDataStruct.mx,IMURawDataStruct.my,IMURawDataStruct.mz,
										    					IMURawDataStruct.presure,IMURawDataStruct.temperature)
	default:

	}
	SaveLog([]byte(fmt.Sprintf("%d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d\n",IMURawDataStruct.gx,IMURawDataStruct.gy,IMURawDataStruct.gz,
																IMURawDataStruct.ax,IMURawDataStruct.ay,IMURawDataStruct.az,
																IMURawDataStruct.mx,IMURawDataStruct.my,IMURawDataStruct.mz,
										    					IMURawDataStruct.presure,IMURawDataStruct.temperature)))
}

func DealMavlinkMsgRecv(buf []byte) {
	//解析

	MavlinkMsgAnalyze(buf[MavlinkIndexMsgId], buf[MavlinkPackHeadLen:]) 

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
