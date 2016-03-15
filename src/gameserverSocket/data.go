package gameserverSocket

import (
	"errors"
	"time"
)

var (
	dataChan       = make(chan *RequestObject, 1024)
	handleDataFunc func(*RequestObject)
)

func init() {
	go handleData()
}

func handleData() {
	for {
		select {
		case data := <-dataChan:
			if handleDataFunc == nil {
				panic(errors.New("处理数据的方法尚未注册"))
			}
			handleDataFunc(data)
		default:
			// 休眠一下，防止CPU过高
			time.Sleep(5 * time.Millisecond)
		}
	}
}

func addData(requestObj *RequestObject) {
	dataChan <- requestObj
}

// 注册处理数据的方法
// handle：处理数据的方法
func RegisterHandleDataFunc(handle func(*RequestObject)) {
	handleDataFunc = handle
}
