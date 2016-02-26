package clientSocket

// Web服务器的响应对象
type ResponseObject struct {
	// 响应结果的状态值
	Code ResponseCode

	// 响应结果的状态值所对应的描述信息
	Message string

	// 响应结果的数据
	Data interface{}
}

// 创建新的响应对象
// 返回值：
// 相应对象
func NewResponseObject() *ResponseObject {
	return &ResponseObject{
		Code:    Success,
		Message: "",
		Data:    nil,
	}
}

// 设置数据错误的Code
func (responseObj *ResponseObject) SetDataError() {
	responseObj.SetResponseCode(DataError)
}

// 设置客户端数据错误的Code
func (responseObj *ResponseObject) SetClientDataError() {
	responseObj.SetResponseCode(ClientDataError)
}

// 设置Code
// rc：响应Code
func (responseObj *ResponseObject) SetResponseCode(rc ResponseCode) {
	responseObj.Code = rc
	responseObj.Message = rc.Message()
}

// 设置数据
// data：数据
func (responseObj *ResponseObject) SetData(data interface{}) {
	responseObj.Data = data
}
