package rpc

import (
	"errors"
	"time"
)

var (
	// 服务器监听地址
	mServerHost string

	// 检测客户端过期的时间间隔（单位：秒）
	mCheckExpiredInterval time.Duration = time.Duration(0)

	// 客户端过期的时间（单位：秒）
	mClientExpiredTime time.Duration = time.Duration(0)
)

// 设置服务器参数
// serverHost：服务器监听地址
// checkExpiredInterval：检测客户端过期的时间间隔（单位：秒）
// clientExpireTime：客户端过期的时间（单位：秒）
func SetParam(serverHost string, checkExpiredInterval, clientExpireTime time.Duration) {
	mServerHost = serverHost
	mCheckExpiredInterval = checkExpiredInterval
	mClientExpiredTime = clientExpireTime
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

// 获取检测客户端过期的时间间隔（单位：秒）
// 返回值：
// 检测客户端过期的时间间隔（单位：秒）
func CheckExpiredInterval() time.Duration {
	if mCheckExpiredInterval == time.Duration(0) {
		panic(errors.New("mCheckExpiredInterval尚未设置，请先设置"))
	}

	return mCheckExpiredInterval
}

// 获取客户端过期的秒数
// 返回值：
// 客户端过期的秒数
func ClientExpiredTime() time.Duration {
	if mClientExpiredTime == time.Duration(0) {
		panic(errors.New("mClientExpiredTime尚未设置，请先设置"))
	}

	return mClientExpiredTime
}
