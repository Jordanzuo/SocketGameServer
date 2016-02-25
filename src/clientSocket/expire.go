package clientSocket

import (
	"fmt"
	"github.com/Jordanzuo/goutil/logUtil"
	"time"
)

// 清理过期的客户端
func clearExpiredClient() {
	for {
		// 休眠指定的时间（单位：秒）(放在此处是因为程序刚启动时并没有过期的客户端，所以先不用占用资源；并且此时LogPath尚未设置，如果直接执行后面的代码会出现panic异常)
		time.Sleep(CheckExpiredInterval * time.Second)

		beforeClientCount := GetClientCount()

		// 获取过期的客户端列表
		expiredClientList := GetExpiredClientList()
		expiredClientCount := len(expiredClientList)
		if expiredClientCount == 0 {
			continue
		}

		for _, item := range expiredClientList {
			Expire(item)
		}

		// 记录日志
		logUtil.Log(fmt.Sprintf("清理前的客户端数量为：%d，本次清理不活跃的数量为：%d", beforeClientCount, expiredClientCount), logUtil.Debug, true)
	}
}
