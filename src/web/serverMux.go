package web

import (
	"encoding/json"
	"fmt"
	"github.com/Jordanzuo/SocketGameServer/src/bll/configBLL"
	"github.com/Jordanzuo/SocketGameServer/src/model/responseDataObject"
	"github.com/Jordanzuo/goutil/logUtil"
	"github.com/Jordanzuo/goutil/securityUtil"
	"net/http"
	"sort"
	"strings"
)

var (
	// 处理的方法映射表
	funcMap = make(map[string]func(http.ResponseWriter, *http.Request) *responseDataObject.WebResponseObject)
)

// 注册API
// apiName：API名称
// callback：回调方法
func registerAPI(apiName string, callback func(http.ResponseWriter, *http.Request) *responseDataObject.WebResponseObject) {
	funcMap[fmt.Sprintf("/API/%s", apiName)] = callback
}

// 定义自定义的Mux对象
type SelfDefineMux struct {
}

func (mux *SelfDefineMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	responseObj := responseDataObject.NewWebResponseObject()

	// 判断是否是POST方法
	if r.Method != "POST" {
		responseObj.SetResultStatus(responseDataObject.OnlySupportPOST)
		responseResult(w, responseObj)
		return
	}

	// 根据路径选择不同的处理方法
	if f, ok := funcMap[r.RequestURI]; !ok {
		responseObj.SetResultStatus(responseDataObject.APINotDefined)
		responseResult(w, responseObj)
	} else {
		responseObj = f(w, r)
		responseResult(w, responseObj)
	}
}

// 解析Form数据并记录日志
func parseFormAndLog(apiName string, r *http.Request) {
	r.ParseForm()

	logUtil.Log(fmt.Sprintf("APIName:%s，请求数据：%v", apiName, r.Form), logUtil.Info, true)
}

// 验证签名
func verifySign(r *http.Request) bool {
	keySlice := make([]string, 0, 10)
	valueMap := make(map[string][]string)
	sign := ""

	// 取出所有的Form数据
	for key, value := range r.Form {
		if strings.ToLower(key) == "sign" {
			sign = value[0]
		} else {
			keySlice = append(keySlice, key)
			valueMap[key] = value
		}
	}

	// 对key进行排序
	sort.Strings(keySlice)

	// 组装原始加密字符串
	rawString := ""
	for _, key := range keySlice {
		if values, ok := valueMap[key]; ok {
			rawString = rawString + fmt.Sprintf("%s=%s&", key, strings.Join(values, ""))
		}
	}
	rawString += configBLL.SecretKey

	// 判断签名是否正确
	if sign == securityUtil.Md5String(rawString, false) {
		return true
	}

	return false
}

// 输出结果
func responseResult(w http.ResponseWriter, responseObj *responseDataObject.WebResponseObject) {
	responseBytes, _ := json.Marshal(responseObj)
	fmt.Fprintln(w, string(responseBytes))
}
