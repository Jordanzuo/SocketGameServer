package clientSocket

import (
	"encoding/json"
	"fmt"
	"github.com/Jordanzuo/goutil/logUtil"
)

var (
	PlayerModuleName  = "Player"
	LoginMethodName   = "Login"
	ReloginMethodName = "Relogin"
)

func handleRequest(clientObj *Client, id int, content []byte) {
	responseObj := NewResponseObject()

	// 将content进行发序列化
	var requestObj RequestObject
	if err := json.Unmarshal(content, &requestObj); err != nil {
		logUtil.Log(fmt.Sprintf("反序列化客户端数据出错，错误信息为：%s", err), logUtil.Error, true)
		responseObj.SetClientDataError()
		return
	}

	// 设定IP
	requestObj.IP = clientObj.IP

	if requestObj.ModuleName == PlayerModuleName && requestObj.MethodName == LoginMethodName {

	} else if requestObj.ModuleName == PlayerModuleName && requestObj.MethodName == ReloginMethodName {

	} else {
		// 给公共字段赋值
		requestObj.PlayerId = clientObj.PlayerId
		requestObj.PartnerId = clientObj.PartnerId
		requestObj.ServerId = clientObj.ServerId
		requestObj.GameVersionId = clientObj.GameVersionId
		requestObj.ResourceVersionId = clientObj.ResourceVersionId
		requestObj.MAC = clientObj.MAC
		requestObj.IDFA = clientObj.IDFA
	}

	// 发送数据给游戏服务器
	// returnBytes, err := webUtil.PostWebData(GameServerAPIUrl, postDict, nil)
	// if err != nil {
	// 	logUtil.Log(fmt.Sprintf("请求GameServer服务器错误，错误信息为：%s", err), logUtil.Error, true)
	// 	responseObj.SetClientDataError()
	// 	return
	// }
}
