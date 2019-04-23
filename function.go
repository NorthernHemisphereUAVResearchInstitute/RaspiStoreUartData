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
