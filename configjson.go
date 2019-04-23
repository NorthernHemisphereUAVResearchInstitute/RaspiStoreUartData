package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	ForwarderJson = StructForwarderJson{}
)

type InterfaceUnMarshal interface {
	UnMarshal(bytes []byte) error
}

type StructForwarderJson struct {
	Port    string `json:"Port"`    //连接串口
	Baud    int    `json:"Baud"`    //波特率
	SaveLog bool   `json:"SaveLog"` //是否保存数据
}

func (cfg *StructForwarderJson) UnMarshal(bytes []byte) error {
	return json.Unmarshal(bytes, cfg)
}

func getConfigPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println("GetPath: ", err.Error())
		return ""
	}
	dir += "/cfg/"
	return dir
}

func readConfig(fileName string, config InterfaceUnMarshal) {
	dir := getConfigPath()
	cfgfilename := dir + fileName

	bytes, err := ioutil.ReadFile(cfgfilename)
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
	}
	err = config.UnMarshal(bytes)
	if err != nil {
		fmt.Println("UnMarshal: ", err.Error())
		panic(err)
	}
}

func InitConfig() {
	fmt.Println("read config")
	readConfig("conf.json", &ForwarderJson)
	fmt.Println(ForwarderJson)
}
