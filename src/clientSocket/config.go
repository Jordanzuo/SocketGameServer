package clientSocket

import (
	"github.com/Jordanzuo/SocketGameServer/src/bll/configBLL"
	"time"
)

const (
	// 配置文件名称
	CONFIG_FILE_NAME = "config_client.ini"
)

var (
	// 服务器监听地址
	ListenAddress string

	// 检测客户端过期的时间间隔（单位：秒）
	CheckExpiredInterval time.Duration

	// 客户端过期的时间（单位：秒）
	ClientExpiredTime time.Duration

	// 游戏服务器地址
	GameServerAPIUrl string
)

func init() {
	// 读取配置文件内容，json类型
	config := configBLL.ReadConfigFile(CONFIG_FILE_NAME)

	ListenAddress = configBLL.GetStringConfig(config, "ListenAddress")
	CheckExpiredInterval = time.Duration(configBLL.GetIntConfig(config, "CheckExpiredInterval"))
	ClientExpiredTime = time.Duration(configBLL.GetIntConfig(config, "ClientExpiredTime"))
	GameServerAPIUrl = configBLL.GetStringConfig(config, "GameServerAPIUrl")
}
