package gameserverSocket

import (
	"encoding/json"
	"fmt"
	"github.com/Jordanzuo/goutil/logUtil"
)

func handleRequest(content []byte) {
	// 将content反序列化成RequestObject
	var requestObj *RequestObject
	if err := json.Unmarshal(content, &requestObj); err != nil {
		logUtil.Log(fmt.Sprintf("反序列化游戏服务端数据错误，错误信息为：%s", err), logUtil.Error, true)
		return
	}

	// 将数据添加到channel中
	addData(requestObj)
}
