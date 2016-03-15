package gameserverSocket

// 游戏服务端的请求数据对象
type RequestObject struct {
	// 需要发送的玩家Id集合，空代表发送给所有
	To []string

	// 需要发送的数据
	Data interface{}
}
