package clientSocket

// 服务端响应结果的状态对象，成功是0，非成功以负数来表示
type ResponseCode int

// 定义所有的响应结果的状态枚举值，此种实现方式是GO语言的标准方式
const (
	// 成功
	Success ResponseCode = 0

	// 数据错误
	DataError ResponseCode = -31

	// 客户端数据错误
	ClientDataError ResponseCode = -32

	// 在另一台设备上登录
	LoginOnAnotherDevice ResponseCode = -1101
)

// 定义所有的响应结果的状态值所对应的字符串描述信息，如果要增加状态枚举，则此处也要相应地增加
var codeMap = map[ResponseCode]string{
	0:     "Success",
	-31:   "DataError",
	-32:   "ClientDataError",
	-1101: "LoginOnAnotherDevice",
}

// 返回响应状态枚举值对应的描述信息字符串
// 返回值：
// 枚举值对应的描述信息字符串
func (rc ResponseCode) Message() string {
	return codeMap[rc]
}
