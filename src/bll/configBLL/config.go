/*
项目配置逻辑处理包，初始化所有的配置内容，其它代码需要配置时都从此包内获取
*/
package configBLL

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"time"
)

const (
	// 配置文件名称
	CONFIG_FILE_NAME = "config.ini"
)

var (
	// 服务器组Id
	ServerGroupId int

	// Socket服务器主机地址
	SocketServerHost string

	// 检查客户端过期的时间间隔
	CheckExpiredInterval time.Duration

	// 客户端过期的秒数
	ClientExpiredTime time.Duration

	// Web服务器主机地址
	WebServerHost string

	// 加密Key
	SecretKey string
)

func init() {
	// 由于服务器的运行依赖于init中执行的逻辑，所以如果出现任何的错误都直接panic，让程序启动失败；而不是让它启动成功，但是在运行时出现错误

	// 读取配置文件（一次性读取整个文件，则使用ioutil）
	bytes, err := ioutil.ReadFile(CONFIG_FILE_NAME)
	if err != nil {
		panic(errors.New(fmt.Sprintf("读取配置文件的内容出错，错误信息为：%s", err)))
	}

	// 使用json反序列化
	config := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &config); err != nil {
		panic(errors.New(fmt.Sprintf("反序列化配置文件的内容出错，错误信息为：%s", err)))
	}

	// 解析ServerGroupId
	ServerGroupId = initIntValue(config, "ServerGroupId")

	// 解析SocketServerHost
	SocketServerHost = initStringValue(config, "SocketServerHost")

	// 解析检查客户端过期的时间间隔
	CheckExpiredInterval = time.Duration(initIntValue(config, "CheckExpiredInterval"))

	// 解析客户端过期的秒数
	ClientExpiredTime = time.Duration(initIntValue(config, "ClientExpiredTime"))

	// 解析WebServerHost
	WebServerHost = initStringValue(config, "WebServerHost")

	// 解析SecretKey
	SecretKey = initStringValue(config, "SecretKey")
}

func initIntValue(config map[string]interface{}, configName string) int {
	configValue, ok := config[configName]
	if !ok {
		panic(errors.New(fmt.Sprintf("不存在名为%s的配置或配置为空", configName)))
	}
	configValue_float, ok := configValue.(float64)
	if !ok {
		panic(errors.New(fmt.Sprintf("%s必须为int型", configName)))
	}

	return int(configValue_float)
}

func initStringValue(config map[string]interface{}, configName string) string {
	configValue, ok := config[configName]
	if !ok {
		panic(errors.New(fmt.Sprintf("不存在名为%s的配置或配置为空", configName)))
	}
	configValue_string, ok := configValue.(string)
	if !ok {
		panic(errors.New(fmt.Sprintf("%s必须为string型", configName)))
	}

	return configValue_string
}
