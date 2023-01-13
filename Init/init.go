package Init

import (
	"db2struct/config"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
)

func Init() {
	//指定对应的json配置文件
	b, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("读取文件失败,失败原因是:", err)
		return
	}
	err = jsoniter.Unmarshal(b, config.SysConfig)
	if err != nil {
		fmt.Println("转换失败,失败原因是:", err)
		return
	}
}
