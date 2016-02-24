/*
请求数据对象包
*/
package requestDataObject

type RequestDataObject struct {
	// 模块名称
	ModuleName string

	// 方法名称
	MethodName string

	// 玩家Id（登录以外的方法此参数不为空）
	PlayerId string

	// 数据（
	// 如果是Login，则参数为:PartnerId, ServerId, UserId, LoginInfo, GameVersionId, ResourceVersionId, MAC, IDFA,
	// Device, OS, RandNum, EncryptString
	// 如果是Relogin，则参数为:PlayerId
	Data []interface{}

	// 下面的字段是在除登录外的方法时才会赋值

	// 合作商Id
	PartnerId int

	// 服务器Id
	ServerId int

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
}
