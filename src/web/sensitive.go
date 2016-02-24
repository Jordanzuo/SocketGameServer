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

	return responseObj
}
