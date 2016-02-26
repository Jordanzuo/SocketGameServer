package web

import (
	"github.com/Jordanzuo/SocketGameServer/src/config"
)

const (
	// 配置文件名称
	CONFIG_FILE_NAME = "config_web.ini"
)

var (
	// 服务器监听地址
	ListenAddress string

	// 服务器加密Key
	SecretKey string
)

func init() {
	// 读取配置文件内容，json类型
	configMap := config.ReadConfigFile(CONFIG_FILE_NAME)

	ListenAddress = config.GetStringConfig(configMap, "ListenAddress")
	SecretKey = config.GetStringConfig(configMap, "SecretKey")
}
