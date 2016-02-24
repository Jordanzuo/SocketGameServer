package web

import (
	"errors"
)

var (
	// 服务器监听地址
	mServerHost string
)

// 设置服务器参数
// serverHost：服务器监听地址
func SetParam(serverHost string) {
	mServerHost = serverHost
}

// 获取服务器的监听地址
// 返回值：
// 服务器的监听地址
func ServerHost() string {
	if mServerHost == "" {
		panic(errors.New("mServerHost尚未设置，请先设置"))
	}

	return mServerHost
}
