package clientSocket

// 请求数据对象
type RequestObject struct {
	// 模块名称
	ModuleName string

	// 方法名称
	MethodName string

	// 合作商Id
	PartnerId int

	// 服务器Id
	ServerId int

	// 玩家Id（登录以外的方法此参数不为空）
	PlayerId string

	// 游戏版本号
	GameVersionId int

	// 资源版本号
	ResourceVersionId int

	// 客户端的IP地址
	IP string

	// 客户端的MAC地址
	MAC string

	// 客户端的IDFA
	IDFA string

	// 请求发送的时间戳
	SendTick int64

	// 数据
	Data []interface{}
}
