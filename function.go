package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

//获取当前时间，并格式变换
func GetNowTime() string {
	return time.Now().String()
}

//获取当前时间，并格式变换
func GetNowDateTimeAsYYMMDDHHMISS() string {
	return string([]byte(GetNowTime())[:len("2015-01-01 12:13:14")])
}

//获取当前日期，并格式变换
func GetNowDateAsYYMMDD() string {
	return string([]byte(GetNowTime())[:len("2015-01-01")])
}

//获取当前月份，并格式变换
func GetNowMonthAsYYMM() string {
	return string([]byte(GetNowTime())[:len("2015-01")])
}

func Int162Bytes(value int16, data []byte) {
	//binary.LittleEndian.PutUint16(data, uint16(value))

	b_buf := bytes.NewBuffer([]byte{})
	binary.Write(b_buf, binary.LittleEndian, value)
	copy(data[:2], b_buf.Bytes())

}
func Uint162Bytes(value uint16, data []byte) {
	binary.LittleEndian.PutUint16(data, value)
}
func Int322Bytes(value int32, data []byte) {
	//binary.LittleEndian.PutUint32(data, uint32(value))
	b_buf := bytes.NewBuffer([]byte{})
	binary.Write(b_buf, binary.LittleEndian, value)
	copy(data[:4], b_buf.Bytes())
}
func Uint322Bytes(value uint32, data []byte) {
	binary.LittleEndian.PutUint32(data, value)
}

func Float322Bytes(data float32, value []byte) {
	binary.LittleEndian.PutUint32(value, math.Float32bits(data))
}

func Float642Bytes(data float64, value []byte) {
	binary.LittleEndian.PutUint64(value, math.Float64bits(data))
}

func GetInt32ValueByString(value string) int32 {
	v, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		fmt.Println("GetInt32ValueByString", value, err)
	}
	return int32(v)
}

func Getfloat64ValueByString(value string) float64 {
	v, err := strconv.ParseFloat(value, 10)
	if err != nil {
		fmt.Println("Getfloat64ValueByString", value, err)
	}
	return v
}
func Bytes2Uint16(data []byte) (value uint16) {
	value = binary.LittleEndian.Uint16(data)
	return value
}

func Bytes2Int16(data []byte) (value int16) {
	b_buf := bytes.NewBuffer(data)
	binary.Read(b_buf, binary.LittleEndian, &value)
	return value
}

func Bytes2Uint32(data []byte) (value uint32) {
	value = binary.LittleEndian.Uint32(data)
	return value
}

func Bytes2Int32(data []byte) (value int32) {
	b_buf := bytes.NewBuffer(data)
	binary.Read(b_buf, binary.LittleEndian, &value)
	return value
}

func Bytes2Uint64(data []byte) (value uint64) {
	value = binary.LittleEndian.Uint64(data)
	return value
}

func Bytes2float32(data []byte) (value float32) {
	v := binary.LittleEndian.Uint32(data)
	value = math.Float32frombits(v)
	return value
}

func Bytes2float64(data []byte) (value float64) {
	v := binary.LittleEndian.Uint32(data)
	value = float64(math.Float32frombits(v))
	return value
}

func EightBytes2float64(data []byte) (value float64) {
	v := binary.LittleEndian.Uint64(data)
	value = math.Float64frombits(v)
	return value
}

func GetModulePath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println("GetPath: ", err.Error())
		return ""
	}
	dir += "/"
	return dir
}

func SleepForever() {
	for {
		time.Sleep(5 * time.Second)
	}
}

func Hex2Ascii(hexdata []byte, asciidata []byte) int {
	hexlen := len(hexdata)
	asciilen := len(asciidata)
	if hexlen*3 >= asciilen {
		return 0
	} else {
		for i := 0; i < hexlen; i++ {
			asciidata[i*3] = hex2char(hexdata[i] & 0xf0 >> 4)
			asciidata[i*3+1] = hex2char(hexdata[i] & 0x0f)
			asciidata[i*3+2] = byte(' ')
		}
		return hexlen * 3
	}
}

func hex2char(hex byte) byte {
	if hex >= 0 && hex <= 9 {
		hex += byte('0')
	} else if hex >= 10 && hex <= 15 {
		hex += byte('A') - 10
	} else {
		hex = byte(' ')
	}
	return hex
}

//------------------------------------------------------------------------------

//计算imu校验和
func CheckCrc16(data []byte) uint16 {
	length := len(data)
	if length < 1 {
		return 0xffff
	}

	crcTmp := uint16(0xffff)
	for i := 1; i < length; i++ {
		crcTmp = crc_accumulate(data[i], crcTmp)
	}
	return crcTmp
}

func crc_accumulate(acrc byte, crc16 uint16) uint16 {
	ch := (acrc ^ (byte)(crc16&0x00ff))
	ch = byte(ch ^ (ch << 4))
	ch16 := uint16(ch)
	newcrc16 := (uint16)((crc16 >> 8) ^ (ch16 << 8) ^ (ch16 << 3) ^ (ch16 >> 4))
	return newcrc16
}


const (
	 TRANS_HEAD= 0x5a
	 HEAD_SIZE= 1
	 LEN_SIZE=2
	 CRC_SIZE= 2
	 HEAD_LENGTH_LEN =    (HEAD_SIZE + LEN_SIZE)
	 HEAD_LENGTH_CRC_LEN =    (HEAD_LENGTH_LEN + CRC_SIZE)
)

 var CRC16Table = [256]uint16{  0x0000, 0xC0C1, 0xC181, 0x0140, 0xC301, 0x03C0, 0x0280, 0xC241, 0xC601, 0x06C0, 0x0780, 0xC741, 0x0500, 0xC5C1, 0xC481, 0x0440, 0xCC01, 0x0CC0, 0x0D80,
        0xCD41, 0x0F00, 0xCFC1, 0xCE81, 0x0E40, 0x0A00, 0xCAC1, 0xCB81, 0x0B40, 0xC901, 0x09C0, 0x0880, 0xC841, 0xD801, 0x18C0, 0x1980, 0xD941, 0x1B00, 0xDBC1, 0xDA81, 0x1A40,
        0x1E00, 0xDEC1, 0xDF81, 0x1F40, 0xDD01, 0x1DC0, 0x1C80, 0xDC41, 0x1400, 0xD4C1, 0xD581, 0x1540, 0xD701, 0x17C0, 0x1680, 0xD641, 0xD201, 0x12C0, 0x1380, 0xD341, 0x1100,
        0xD1C1, 0xD081, 0x1040, 0xF001, 0x30C0, 0x3180, 0xF141, 0x3300, 0xF3C1, 0xF281, 0x3240, 0x3600, 0xF6C1, 0xF781, 0x3740, 0xF501, 0x35C0, 0x3480, 0xF441, 0x3C00, 0xFCC1,
        0xFD81, 0x3D40, 0xFF01, 0x3FC0, 0x3E80, 0xFE41, 0xFA01, 0x3AC0, 0x3B80, 0xFB41, 0x3900, 0xF9C1, 0xF881, 0x3840, 0x2800, 0xE8C1, 0xE981, 0x2940, 0xEB01, 0x2BC0, 0x2A80,
        0xEA41, 0xEE01, 0x2EC0, 0x2F80, 0xEF41, 0x2D00, 0xEDC1, 0xEC81, 0x2C40, 0xE401, 0x24C0, 0x2580, 0xE541, 0x2700, 0xE7C1, 0xE681, 0x2640, 0x2200, 0xE2C1, 0xE381, 0x2340,
        0xE101, 0x21C0, 0x2080, 0xE041, 0xA001, 0x60C0, 0x6180, 0xA141, 0x6300, 0xA3C1, 0xA281, 0x6240, 0x6600, 0xA6C1, 0xA781, 0x6740, 0xA501, 0x65C0, 0x6480, 0xA441, 0x6C00,
        0xACC1, 0xAD81, 0x6D40, 0xAF01, 0x6FC0, 0x6E80, 0xAE41, 0xAA01, 0x6AC0, 0x6B80, 0xAB41, 0x6900, 0xA9C1, 0xA881, 0x6840, 0x7800, 0xB8C1, 0xB981, 0x7940, 0xBB01, 0x7BC0,
        0x7A80, 0xBA41, 0xBE01, 0x7EC0, 0x7F80, 0xBF41, 0x7D00, 0xBDC1, 0xBC81, 0x7C40, 0xB401, 0x74C0, 0x7580, 0xB541, 0x7700, 0xB7C1, 0xB681, 0x7640, 0x7200, 0xB2C1, 0xB381,
        0x7340, 0xB101, 0x71C0, 0x7080, 0xB041, 0x5000, 0x90C1, 0x9181, 0x5140, 0x9301, 0x53C0, 0x5280, 0x9241, 0x9601, 0x56C0, 0x5780, 0x9741, 0x5500, 0x95C1, 0x9481, 0x5440,
        0x9C01, 0x5CC0, 0x5D80, 0x9D41, 0x5F00, 0x9FC1, 0x9E81, 0x5E40, 0x5A00, 0x9AC1, 0x9B81, 0x5B40, 0x9901, 0x59C0, 0x5880, 0x9841, 0x8801, 0x48C0, 0x4980, 0x8941, 0x4B00,
        0x8BC1, 0x8A81, 0x4A40, 0x4E00, 0x8EC1, 0x8F81, 0x4F40, 0x8D01, 0x4DC0, 0x4C80, 0x8C41, 0x4400, 0x84C1, 0x8581, 0x4540, 0x8701, 0x47C0, 0x4680, 0x8641, 0x8201, 0x42C0,
        0x4380, 0x8341, 0x4100, 0x81C1, 0x8081, 0x4040, }

func CRC16(buf []byte, len int)uint16 {
     crc_result := uint16(0x00)
     table_num := uint16(0x00)
     
    for i := 0;i < len;i++  {
        table_num = ((crc_result & 0xff) ^ (uint16)(buf[i] & 0xff))
        crc_result = ((crc_result >> 8) & 0xff) ^ CRC16Table[table_num]
    }   
    return crc_result
}
/*
	函数crc_check： 检查数据是否完整
	return:
		-1 校验失败， 0 检验成功
*/
/*func crc_check(protocol_buf []byte)uint16 {
	/*if(protocol_buf == nil){
		return -1
	}*/
    //payload_len := (protocol_buf[1] << 8 & 0xff00) | (protocol_buf[2] & 0xff)
    //total_size := int(HEAD_SIZE + LEN_SIZE + payload_len + CRC_SIZE)
    //src_crc_sum := (protocol_buf[total_size - 2] << 8 & 0xff00) | (protocol_buf[total_size - 1] & 0xff)
    //p_data := protocol_buf + HEAD_SIZE + LEN_SIZE
    //rc_sum_calc := CRC16(protocol_buf[:], int(payload_len))
    /*if crc_sum_calc != src_crc_sum {
        return -1
    }
    return 0

    return crc_sum_calc
}*/
