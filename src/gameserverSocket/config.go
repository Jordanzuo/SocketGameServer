package gameserverSocket

import (
	"github.com/Jordanzuo/SocketGameServer/src/bll/configBLL"
)

const (
	// 配置文件名称
	CONFIG_FILE_NAME = "config_gameserver.ini"
)

var (
	// 服务器监听地址
	ListenAddress string
)

func init() {
	// 读取配置文件内容，json类型
	config := configBLL.ReadConfigFile(CONFIG_FILE_NAME)

	ListenAddress = configBLL.GetStringConfig(config, "ListenAddress")
}
