package main

import (
	"fmt"
)

var (
	Amp_phaStruct = StructAMP_PHA{}
	IMU_Raw_Data  = Struct_IMU_Raw_Data{}
	Attitude_Struct=Struct_Attitude_Data{}
)

type StructAMP_PHA struct {
	ma1_26 float32
	ma1_28 float32
	ma1_30 float32
	ma2_26 float32
	ma2_28 float32
	ma2_30 float32
	ma3_26 float32
	ma3_28 float32
	ma3_30 float32

	mb1_26 float32
	mb1_28 float32
	mb1_30 float32
	mb2_26 float32
	mb2_28 float32
	mb2_30 float32
	mb3_26 float32
	mb3_28 float32
	mb3_30 float32
}

type Struct_IMU_Raw_Data struct {

	acce_x int16
	acce_y int16
	acce_z int16

	gyro_x int16
	gyro_y int16
	gyro_z int16
}

type Struct_Attitude_Data struct {

	roll float32
	pitch float32
	yaw float32

	q0 float32
	q1 float32
	q2 float32
	q3 float32
}

func MavAnalyzeVelPos(buf []byte) (amp_pha StructAMP_PHA) {
	index := 0

	amp_pha.ma1_26 = Bytes2float32(buf[index : index+4])
	index += 4
	amp_pha.ma1_28 = Bytes2float32(buf[index : index+4])
	index += 4
	amp_pha.ma1_30 = Bytes2float32(buf[index : index+4])
	index += 4
	amp_pha.ma2_26 = Bytes2float32(buf[index : index+4])
	index += 4
	amp_pha.ma2_28 = Bytes2float32(buf[index : index+4])
	index += 4
	amp_pha.ma2_30 = Bytes2float32(buf[index : index+4])
	index += 4
	amp_pha.ma3_26 = Bytes2float32(buf[index : index+4])
	index += 4
	amp_pha.ma3_28 = Bytes2float32(buf[index : index+4])
	index += 4
	amp_pha.ma3_30 = Bytes2float32(buf[index : index+4])
	index += 4
	amp_pha.mb1_26 = Bytes2float32(buf[index : index+4])
	index += 4
	amp_pha.mb1_28 = Bytes2float32(buf[index : index+4])
	index += 4
	amp_pha.mb1_30 = Bytes2float32(buf[index : index+4])
	index += 4
	amp_pha.mb2_26 = Bytes2float32(buf[index : index+4])
	index += 4
	amp_pha.mb2_28 = Bytes2float32(buf[index : index+4])
	index += 4
	amp_pha.mb2_30 = Bytes2float32(buf[index : index+4])
	index += 4
	amp_pha.mb3_26 = Bytes2float32(buf[index : index+4])
	index += 4
	amp_pha.mb3_28 = Bytes2float32(buf[index : index+4])
	index += 4
	amp_pha.mb3_30 = Bytes2float32(buf[index : index+4])
	index += 4
	

	return amp_pha
}

func Mav_Analyze_IMU_Raw_Data(buf []byte) (imu_raw_data Struct_IMU_Raw_Data) {
	index := 0
	imu_raw_data.acce_x = Bytes2Int16(buf[index : index+2])
	index += 2
	imu_raw_data.acce_y = Bytes2Int16(buf[index : index+2])
	index += 2
	imu_raw_data.acce_z = Bytes2Int16(buf[index : index+2])
	index += 2
	imu_raw_data.gyro_x = Bytes2Int16(buf[index : index+2])
	index += 2
	imu_raw_data.gyro_y = Bytes2Int16(buf[index : index+2])
	index += 2
	imu_raw_data.gyro_z = Bytes2Int16(buf[index : index+2])
	index += 2


	return imu_raw_data

}

func Mav_Analyze_Attitude_Data(buf []byte)(attitude_data Struct_Attitude_Data) {
	index :=0
	attitude_data.roll=Bytes2float32(buf[index : index+4])
	index += 4
	attitude_data.pitch=Bytes2float32(buf[index : index+4])
	index += 4
	attitude_data.yaw=Bytes2float32(buf[index : index+4])
	index += 4
	attitude_data.q0=Bytes2float32(buf[index : index+4])
	index += 4
	attitude_data.q1=Bytes2float32(buf[index : index+4])
	index += 4
	attitude_data.q2=Bytes2float32(buf[index : index+4])
	index += 4
	attitude_data.q3=Bytes2float32(buf[index : index+4])
	index += 4

	return attitude_data
	
}

func MavlinkMsgAnalyze(msgdatabuf []byte) {

	/*Amp_phaStruct = MavAnalyzeVelPos(msgdatabuf)
	
	fmt.Printf("%.3f, %.3f, %.3f, %.3f, %.3f,%.3f, %.3f, %.3f, %.3f, %.3f, %.3f, %.3f, %.3f, %.3f，%.3f, %.3f, %.3f, %.3f\n", 
		Amp_phaStruct.ma1_26,Amp_phaStruct.ma1_28,Amp_phaStruct.ma1_30,
		Amp_phaStruct.ma2_26,Amp_phaStruct.ma2_28,Amp_phaStruct.ma2_30,
		Amp_phaStruct.ma3_26,Amp_phaStruct.ma3_28,Amp_phaStruct.ma3_30,
		Amp_phaStruct.mb1_26,Amp_phaStruct.mb1_28,Amp_phaStruct.mb1_30,
		Amp_phaStruct.mb2_26,Amp_phaStruct.mb2_28,Amp_phaStruct.mb2_30,
		Amp_phaStruct.mb3_26,Amp_phaStruct.mb3_28,Amp_phaStruct.mb3_30,
	)

	SaveLog([]byte(fmt.Sprintf("%.3f,%.3f,%.3f,%.3f,%.3f,%.3f,%.3f,%.3f,%.3f,%.3f,%.3f,%.3f,%.3f,%.3f,%.3f,%.3f,%.3f,%.3f\r\n", 
		Amp_phaStruct.ma1_26,Amp_phaStruct.ma1_28,Amp_phaStruct.ma1_30,
		Amp_phaStruct.ma2_26,Amp_phaStruct.ma2_28,Amp_phaStruct.ma2_30,
		Amp_phaStruct.ma3_26,Amp_phaStruct.ma3_28,Amp_phaStruct.ma3_30,
		Amp_phaStruct.mb1_26,Amp_phaStruct.mb1_28,Amp_phaStruct.mb1_30,
		Amp_phaStruct.mb2_26,Amp_phaStruct.mb2_28,Amp_phaStruct.mb2_30,
		Amp_phaStruct.mb3_26,Amp_phaStruct.mb3_28,Amp_phaStruct.mb3_30,)))*/


		/*IMU_Raw_Data=Mav_Analyze_IMU_Raw_Data(msgdatabuf)

		fmt.Printf("%d,%d,%d,%d,%d,%d\r\n", IMU_Raw_Data.acce_x,IMU_Raw_Data.acce_y,IMU_Raw_Data.acce_z,
			IMU_Raw_Data.gyro_x,IMU_Raw_Data.gyro_y,IMU_Raw_Data.gyro_z)

		SaveLog([]byte(fmt.Sprintf("%d,%d,%d,%d,%d,%d\r\n", IMU_Raw_Data.acce_x,IMU_Raw_Data.acce_y,IMU_Raw_Data.acce_z,
			IMU_Raw_Data.gyro_x,IMU_Raw_Data.gyro_y,IMU_Raw_Data.gyro_z)))*/

			Attitude_Struct=Mav_Analyze_Attitude_Data(msgdatabuf)
			fmt.Printf("%.3f,%.3f,%.3f\r\n", Attitude_Struct.roll,Attitude_Struct.pitch,Attitude_Struct.yaw)
}

func DealMavlinkMsgRecv(buf []byte) {
	//解析
	MavlinkMsgAnalyze(buf[4:])
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
