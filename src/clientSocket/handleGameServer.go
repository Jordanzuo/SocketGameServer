package clientSocket

import (
	"github.com/Jordanzuo/SocketGameServer/src/gameserverSocket"
)

func init() {
	gameserverSocket.RegisterHandleDataFunc(handleGameServerData)
}

// 处理游戏服务端数据的方法
// requestObj：游戏服务端发送过来的数据
func handleGameServerData(requestObj *gameserverSocket.RequestObject) {
	clientList := make([]*Client, 0, 1024)

	// 判断是否发送给所有人
	if len(requestObj.To) == 0 {
		clientList = GetAllClient()
	} else {
		clientList = GetClientByPlayerIdList(requestObj.To)
	}

	// 将请求数据发给每个客户端进行处理
	responseObj := NewResponseObject()
	responseObj.SetData(requestObj.Data)
	for _, clientObj := range clientList {
		clientObj.AddPushData(responseObj)
	}
}
