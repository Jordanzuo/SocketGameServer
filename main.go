package main

import (
	"github.com/Jordanzuo/SocketGameServer/src/rpc"
	"github.com/Jordanzuo/goutil/fileUtil"
	"github.com/Jordanzuo/goutil/logUtil"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
)

import (
	_ "github.com/Jordanzuo/SocketGameServer/src/bll/configBLL"
	_ "github.com/Jordanzuo/SocketGameServer/src/bll/sensitiveWordsBLL"
	_ "github.com/Jordanzuo/SocketGameServer/src/bll/webBLL"
)

var (
	wg sync.WaitGroup
)

func init() {
	// 设置日志文件的存储目录
	logUtil.SetLogPath(filepath.Join(fileUtil.GetCurrentPath(), "LOG"))

	// 设置WaitGroup需要等待的数量
	wg.Add(1)
}

// 处理系统信号
func signalProc() {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)

	for {
		// 准备接收信息
		<-sigs

		// 一旦收到信号，则表明管理员希望退出程序，则先保存信息，然后退出
		os.Exit(0)
	}
}

func main() {
	// 处理系统信号
	go signalProc()

	// 启动服务器
	go rpc.StartServer(&wg)

	// 阻塞等待，以免main线程退出
	wg.Wait()
}
