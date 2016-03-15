package clientSocket

import (
	"encoding/json"
	"fmt"
	"github.com/Jordanzuo/goutil/logUtil"
)

func handleRequest(clientObj *Client, id int, content []byte) {
	responseObj := NewResponseObject()

	// 将content进行发序列化
	var requestObj *RequestObject
	if err := json.Unmarshal(content, &requestObj); err != nil {
		logUtil.Log(fmt.Sprintf("反序列化客户端数据出错，错误信息为：%s", err), logUtil.Error, true)
		responseObj.SetClientDataError()
		return
	}

	// 设定IP
	requestObj.IP = clientObj.IP

	// // 重新登录Socket服务器，不需要去访问游戏服务器
	// if requestObj.ModuleName == "Player" && requestObj.MethodName == "ReLogin" {

	// } else {

	// 	// 发送数据给游戏服务器
	// 	// returnBytes, err := webUtil.PostWebData(GameServerAPIUrl, postDict, nil)
	// 	// if err != nil {
	// 	// 	logUtil.Log(fmt.Sprintf("请求GameServer服务器错误，错误信息为：%s", err), logUtil.Error, true)
	// 	// 	responseObj.SetClientDataError()
	// 	// 	return
	// 	// }

	// 	// 如果是登录方法，需要解析返回值中的PlayerId
	// 	if requestObj.ModuleName == "Player" && (requestObj.MethodName == "Login" || requestObj.MethodName == "Login_ForTest") {

	// 	}
	// }
}
