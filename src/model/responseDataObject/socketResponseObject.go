package responseDataObject

// Socket服务器的响应对象
type SocketResponseObject struct {
	// 响应结果的状态值
	Code ResultStatus

	// 响应结果的状态值所对应的描述信息
	Message string

	// 响应结果的数据
	Data interface{}
}

func NewSocketResponseObject() *SocketResponseObject {
	return &SocketResponseObject{
		Code:    Success,
		Message: "",
		Data:    nil,
	}
}

func (responseObject *SocketResponseObject) SetDataError() {
	responseObject.SetResultStatus(DataError)
}

func (responseObject *SocketResponseObject) SetClientDataError() {
	responseObject.SetResultStatus(APIDataError)
}

func (responseObject *SocketResponseObject) SetResultStatus(rs ResultStatus) {
	responseObject.Code = rs
	responseObject.Message = rs.String()
}

func (responseObject *SocketResponseObject) SetData(data interface{}) {
	responseObject.Data = data
}
