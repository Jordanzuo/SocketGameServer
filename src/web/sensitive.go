package web

import (
	"github.com/Jordanzuo/SocketGameServer/src/model/responseDataObject"
	"net/http"
)

var (
	sensitiveAPIName = "sensitive"
)

func init() {
	registerAPI(sensitiveAPIName, sensitiveCallback)
}

func sensitiveCallback(w http.ResponseWriter, r *http.Request) *responseDataObject.WebResponseObject {
	responseObj := responseDataObject.NewWebResponseObject()

	// 解析Form数据并记录日志
	parseFormAndLog(sensitiveAPIName, r)

	// 验证签名
	if verifySign(r) == false {
		responseObj.SetResultStatus(responseDataObject.SignError)
		return responseObj
	}

	// 重新加载敏感词
	// if err := sensitiveWordsBLL.Reload(); err != nil {
	// 	responseObj.SetDataError()
	// 	return responseObj
	// }

	return responseObj
}
