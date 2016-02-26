package web

// 服务端响应结果的状态对象，成功是0，非成功以负数来表示
type ResultStatus int

// 定义所有的响应结果的状态枚举值，此种实现方式是GO语言的标准方式
const (
	// 成功
	Success ResultStatus = 0

	// API数据错误
	APIDataError ResultStatus = -1

	// 只支持POST
	OnlySupportPOST ResultStatus = -2

	// API未定义
	APINotDefined ResultStatus = -3
)

// 定义所有的响应结果的状态值所对应的字符串描述信息，如果要增加状态枚举，则此处也要相应地增加
var status = map[ResultStatus]string{
	0:  "Success",
	-1: "APIDataError",
	-2: "OnlySupportPOST",
	-3: "APINotDefined",
}

// 返回响应状态枚举值对应的描述信息字符串
func (rs ResultStatus) String() string {
	return status[rs]
}
