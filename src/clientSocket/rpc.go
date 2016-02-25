package clientSocket

import (
	"sync"
)

var (
	// 客户端连接列表
	clientMap = make(map[int32]*Client, 1024)
	playerMap = make(map[string]*Client, 1024)

	// 读写锁
	mutex sync.RWMutex
)

// 添加新的客户端
// clientObj：客户端对象
func RegisterClient(clientObj *Client) {
	mutex.Lock()
	defer mutex.Unlock()

	clientMap[clientObj.Id()] = clientObj
}

// 移除客户端
// clientObj：客户端对象
func UnRegisterClient(clientObj *Client) {
	mutex.Lock()
	defer mutex.Unlock()

	// 删除玩家缓存
	if clientObj.PlayerId != "" {
		delete(playerMap, clientObj.PlayerId)
	}

	// 删除客户端缓存
	delete(clientMap, clientObj.Id())
}

// 玩家登陆
// playerId：玩家id
// partnerId：合作商Id
// serverId：服务器Id
// gameVersionId：游戏版本Id
// resourceVersionId：资源版本Id
// mac：MAC
// idfa：IDFA
// 返回值：无
func PlayerLogin(clientObj *Client, playerId string, partnerId, serverId, gameVersionId, resourceVersionId int, mac, idfa string) {
	clientObj.PlayerLogin(playerId, partnerId, serverId, gameVersionId, resourceVersionId, mac, idfa)

	// 添加到玩家列表中
	playerMap[playerId] = clientObj
}

// 返回过期的客户端列表
// 返回值：
// 过期的客户端列表
func GetExpiredClientList() (expiredClientList []*Client) {
	mutex.RLock()
	defer mutex.RUnlock()

	for _, item := range clientMap {
		if item.HasExpired() {
			expiredClientList = append(expiredClientList, item)
		}
	}

	return
}

// 根据客户端Id获取对应的客户端对象
// id：客户端Id
// 返回值：客户端对象
func GetClient(id int32) (*Client, bool) {
	mutex.RLock()
	defer mutex.RUnlock()

	if clientObj, ok := clientMap[id]; ok {
		return clientObj, true
	}

	return nil, false
}

// 获取客户端的数量
// 返回值：
// 客户端数量
func GetClientCount() int {
	mutex.RLock()
	defer mutex.RUnlock()

	return len(clientMap)
}

// 客户端过期
// clientObj：客户端对象
func Expire(clientObj *Client) {
	clientObj.Quit()
}

// 客户端断开连接
// clientObj：客户端对象
func Disconnect(clientObj *Client) {
	// 给GameServer发送玩家下线的数据
	playerId := clientObj.PlayerId
	_ = playerId

	// 移除客户端
	UnRegisterClient(clientObj)
}
