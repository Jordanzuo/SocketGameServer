package web

// Web服务器的响应对象
type ResponseObject struct {
	// 响应结果的状态值
	Code ResultStatus

	// 响应结果的状态值所对应的描述信息
	Message string

	// 响应结果的数据
	Data interface{}
}

func NewResponseObject() *ResponseObject {
	return &WebResponseObject{
		Code:    Success,
		Message: "",
		Data:    nil,
	}
}

func (responseObj *ResponseObject) SetDataError() {
	responseObj.SetResultStatus(DataError)
}

func (responseObj *ResponseObject) SetAPIDataError() {
	responseObj.SetResultStatus(APIDataError)
}

func (responseObj *ResponseObject) SetResultStatus(rs ResultStatus) {
	responseObj.Code = rs
	responseObj.Message = rs.String()
}

func (responseObj *ResponseObject) SetData(data interface{}) {
	responseObj.Data = data
}
