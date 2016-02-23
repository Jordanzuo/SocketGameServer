/*
项目配置的逻辑处理包，初始化所有的配置内容，其它代码需要配置时都从此包内来获取
包括数据库配置和文件配置
*/
package configBLL

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Jordanzuo/SocketGameServer/src/dal"
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

	// 客户端过期的秒数
	ClientExpireSeconds time.Duration

	// Socket服务器主机地址
	SocketServerHost string

	// Web服务器主机地址
	WebServerHost string

	// 加密Key
	SecretKey string

	// 模型数据库连接字符串
	ModelDBConnection string

	// 模型数据库的最大连接数
	ModelDBMaxOpenConns int

	// 模型数据库的最大空闲数
	ModelDBMaxIdleConns int
)

func init() {
	// 由于服务器的运行依赖于init中执行的逻辑，所以如果出现任何的错误都直接panic，让程序启动失败；而不是让它启动成功，但是在运行时出现错误

	// 读取配置文件（一次性读取整个文件，则使用ioutil）
	bytes, err := ioutil.ReadFile(CONFIG_FILE_NAME)
	if err != nil {
		panic(errors.New("读取配置文件的内容出错"))
	}

	// 使用json反序列化
	config := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &config); err != nil {
		panic(errors.New("反序列化配置文件的内容出错"))
	}

	// 解析ServerGroupId
	ServerGroupId = initServerGroup(config)

	// 解析客户端过期的秒数
	ClientExpireSeconds = initClientExpireSeconds(config)

	// 解析SocketServerHost
	SocketServerHost = initSocketServerHost(config)

	// 解析WebServerHost
	WebServerHost = initWebServerHost(config)

	// 解析SecretKey
	SecretKey = initSecretKey(config)

	// 解析ModelDBConnection配置
	ModelDBConnection, ModelDBMaxOpenConns, ModelDBMaxIdleConns = initDBConnection(config, "ModelDBConnection", "ModelDBMaxOpenConns", "ModelDBMaxIdleConns")

	// 初始化数据库连接相关的配置
	dal.InitDB(ModelDBConnection, ModelDBMaxOpenConns, ModelDBMaxIdleConns)
}

func initServerGroup(config map[string]interface{}) int {
	serverGroupId, ok := config["ServerGroupId"]
	if !ok {
		panic(errors.New("不存在名为ServerGroupId的配置或配置为空"))
	}
	serverGroupId_float, ok := serverGroupId.(float64)
	if !ok {
		panic(errors.New("ServerGroupId必须为int型"))
	}

	return int(serverGroupId_float)
}

func initClientExpireSeconds(config map[string]interface{}) time.Duration {
	clientExpireSeconds, ok := config["ClientExpireSeconds"]
	if !ok {
		panic(errors.New("不存在名为ClientExpireSeconds的配置或配置为空"))
	}
	clientExpireSeconds_float, ok := clientExpireSeconds.(float64)
	if !ok {
		panic(errors.New("ClientExpireSeconds必须为int型"))
	}

	return time.Duration(int(clientExpireSeconds_float))
}

func initSocketServerHost(config map[string]interface{}) string {
	socketServerHost, ok := config["SocketServerHost"]
	if !ok {
		panic(errors.New("不存在名为SocketServerHost的配置或配置为空"))
	}
	socketServerHost_string, ok := socketServerHost.(string)
	if !ok {
		panic(errors.New("SocketServerHost必须为string型"))
	}

	return socketServerHost_string
}

func initWebServerHost(config map[string]interface{}) string {
	webServerHost, ok := config["WebServerHost"]
	if !ok {
		panic(errors.New("不存在名为WebServerHost的配置或配置为空"))
	}
	webServerHost_string, ok := webServerHost.(string)
	if !ok {
		panic(errors.New("WebServerHost必须为string型"))
	}

	return webServerHost_string
}

func initSecretKey(config map[string]interface{}) string {
	secretKey, ok := config["SecretKey"]
	if !ok {
		panic(errors.New("不存在名为SecretKey的配置或配置为空"))
	}
	secretKey_string, ok := secretKey.(string)
	if !ok {
		panic(errors.New("SecretKey必须为string型"))
	}

	return secretKey_string
}

func initDBConnection(config map[string]interface{}, dbConnectionName, maxOpenConnsName, maxIdleConnsName string) (string, int, int) {
	// 解析DBConnection
	dbConnection, ok := config[dbConnectionName]
	if !ok {
		panic(errors.New(fmt.Sprintf("不存在名为%s的配置或配置为空", dbConnectionName)))
	}
	dbConnection_string, ok := dbConnection.(string)
	if !ok {
		panic(errors.New(fmt.Sprintf("%s必须为字符串类型", dbConnectionName)))
	}

	// 解析MaxOpenConns
	maxOpenConns, ok := config[maxOpenConnsName]
	if !ok {
		panic(errors.New(fmt.Sprintf("不存在名为%s的配置或配置为空", maxOpenConnsName)))
	}
	maxOpenConns_float, ok := maxOpenConns.(float64)
	if !ok {
		panic(errors.New(fmt.Sprintf("%s必须是int型", maxOpenConnsName)))
	}

	// 解析MaxIdleConns
	maxIdleConns, ok := config[maxIdleConnsName]
	if !ok {
		panic(errors.New(fmt.Sprintf("不存在名为%s的配置或配置为空", maxIdleConnsName)))
	}
	maxIdleConns_float, ok := maxIdleConns.(float64)
	if !ok {
		panic(errors.New(fmt.Sprintf("%s必须是int型", maxIdleConnsName)))
	}

	return dbConnection_string, int(maxOpenConns_float), int(maxIdleConns_float)
}
