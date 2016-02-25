package web

import (
	"github.com/Jordanzuo/SocketGameServer/src/model/responseDataObject"
	"net/http"
)

func init() {
	registerAPI("sensitive", sensitiveCallback)
}

func sensitiveCallback(w http.ResponseWriter, r *http.Request) *responseDataObject.WebResponseObject {
	responseObj := responseDataObject.NewWebResponseObject()

	return responseObj
}
