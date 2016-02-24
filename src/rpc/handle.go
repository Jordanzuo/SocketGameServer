package rpc

import (
	"encoding/json"
	"fmt"
	"github.com/Jordanzuo/SocketGameServer/src/model/requestDataObject"
	"github.com/Jordanzuo/SocketGameServer/src/model/responseDataObject"
	"github.com/Jordanzuo/goutil/logUtil"
)

var (
	PlayerModuleName  = "Player"
	LoginMethodName   = "Login"
	ReloginMethodName = "Relogin"
)

func handleRequest(clientObj *Client, id int, content []byte) {
	responseObj := responseDataObject.NewSocketResponseObject()

	// 将content进行发序列化
	var requestDataObj requestDataObject.RequestDataObject
	if err = json.Unmarshal(content, &requestDataObj); err != nil {
		logUtil.Log(fmt.Sprintf("反序列化客户端数据出错，错误信息为：%s", err), logUtil.Error, true)
		responseObj.SetClientDataError()
		return
	}

	if requestDataObj.ModuleName == PlayerModuleName && requestDataObj.MethodName == LoginMethodName {

	} else if requestDataObj.ModuleName == PlayerModuleName && requestDataObj.MethodName == ReloginMethodName {

	} else {
		// 给公共字段赋值
		requestDataObj.PlayerId = clientObj.PlayerId
		requestDataObj.PartnerId = clientObj.PartnerId
		requestDataObj.ServerId = clientObj.ServerId
		requestDataObj.GameVersionId = clientObj.GameVersionId
		requestDataObj.ResourceVersionId = clientObj.ResourceVersionId
		requestDataObj.IP = clientObj.IP
		requestDataObj.MAC = clientObj.MAC
		requestDataObj.IDFA = clientObj.IDFA
	}

	// 发送数据给游戏服务器
	returnBytes, err := webUtil.PostWebData(GameServerUrl(), postDict, nil)
	if err != nil {
		logUtil.Log(fmt.Sprintf("请求GameServer服务器错误，错误信息为：%s", err), logUtil.Error, true)
		responseObj.SetClientDataError()
		return
	}
}
