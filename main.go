package main

import (
	"github.com/Jordanzuo/SocketGameServer/src/clientSocket"
	"github.com/Jordanzuo/SocketGameServer/src/gameserverSocket"
	"github.com/Jordanzuo/SocketGameServer/src/web"
	"github.com/Jordanzuo/goutil/fileUtil"
	"github.com/Jordanzuo/goutil/logUtil"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
)

var (
	wg sync.WaitGroup
)

func init() {
	// 设置日志文件的存储目录
	logUtil.SetLogPath(filepath.Join(fileUtil.GetCurrentPath(), "LOG"))

	// 设置WaitGroup需要等待的数量
	wg.Add(3)
}

// 处理系统信号
func signalProc() {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)

	for {
		// 准备接收信息
		<-sigs

		logUtil.Log("收到退出程序的命令，开始退出……", logUtil.Info, true)

		// 做一些收尾的工作

		logUtil.Log("收到退出程序的命令，退出完成……", logUtil.Info, true)

		// 一旦收到信号，则表明管理员希望退出程序，则先保存信息，然后退出
		os.Exit(0)
	}
}

func main() {
	// 处理系统信号
	go signalProc()

	// 启动为客户端服务的Socket服务器
	go clientSocket.StartServer(&wg)

	// 启动为游戏服务器服务的Socket服务器
	go gameserverSocket.StartServer(&wg)

	// 启动Web服务器
	go web.StartServer(&wg)

	// 阻塞等待，以免main线程退出
	wg.Wait()
}
