package web

import (
	"net/http"
)

func init() {
	registerAPI("sensitive", sensitiveCallback)
}

func sensitiveCallback(w http.ResponseWriter, r *http.Request) *ResponseObject {
	responseObj := NewResponseObject()

	return responseObj
}
