package web

import (
	"errors"
	"fmt"
	"github.com/Jordanzuo/goutil/logUtil"
	"net/http"
	"sync"
)

// 启动服务器
// wg：WaitGroup对象
func StartServer(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()

	logUtil.Log("Web服务器开始监听...", logUtil.Info, true)

	// 设置访问的路由
	mux := new(SelfDefineMux)

	// 启动Web服务器监听
	err := http.ListenAndServe(ServerHost(), mux)
	if err != nil {
		panic(errors.New(fmt.Sprintf("ListenAndServe失败，错误信息为：%s", err)))
	}
}
