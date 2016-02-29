package clientSocket

import (
	"sync"
)

var (
	// 客户端连接列表
	clientMap = make(map[int32]*Client, 1024)
	playerMap = make(map[string]*Client, 1024)

	// 读写锁
	mutex_client sync.RWMutex
	mutex_player sync.RWMutex
)

// 添加新的客户端
// clientObj：客户端对象
func RegisterClient(clientObj *Client) {
	mutex_client.Lock()
	defer mutex_client.Unlock()

	clientMap[clientObj.Id()] = clientObj
}

// 移除客户端
// clientObj：客户端对象
func UnRegisterClient(clientObj *Client) {
	mutex_client.Lock()
	defer mutex_client.Unlock()

	// 删除玩家缓存
	if clientObj.PlayerId != "" {
		delete(playerMap, clientObj.PlayerId)
	}

	// 删除客户端缓存
	delete(clientMap, clientObj.Id())
}

// 玩家登陆
// playerId：玩家id
// 返回值：无
func PlayerLogin(clientObj *Client, playerId string) {
	clientObj.PlayerLogin(playerId)

	// 添加到玩家列表中
	playerMap[playerId] = clientObj
}

// 获取所有的客户端对象
// 返回值：
// 客户端对象列表
func GetAllClient() (clientList []*Client) {
	mutex_client.RLock()
	defer mutex_client.RUnlock()

	for _, clientObj := range clientMap {
		clientList = append(clientList, clientObj)
	}

	return
}

// 返回过期的客户端列表
// 返回值：
// 过期的客户端列表
func GetExpiredClientList() (expiredClientList []*Client) {
	mutex_client.RLock()
	defer mutex_client.RUnlock()

	for _, item := range clientMap {
		if item.HasExpired() {
			expiredClientList = append(expiredClientList, item)
		}
	}

	return
}

// 根据客户端Id获取对应的客户端对象
// id：客户端Id
// 返回值：
// 客户端对象
func GetClient(id int32) (*Client, bool) {
	mutex_client.RLock()
	defer mutex_client.RUnlock()

	if clientObj, ok := clientMap[id]; ok {
		return clientObj, true
	}

	return nil, false
}

// 根据玩家Id获取对应的客户端对象
// playerId：玩家Id
// 返回值：
// 客户端对象
func GetClientByPlayerId(playerId string) (*Client, bool) {
	mutex_player.RLock()
	defer mutex_player.RUnlock()

	if clientObj, ok := playerMap[playerId]; ok {
		return clientObj, true
	}

	return nil, false
}

// 根据玩家Id列表获取对应的客户端列表
// playerIdList：玩家Id列表
// 返回值：
// 客户端列表
func GetClientByPlayerIdList(playerIdList []string) (clientList []*Client) {
	if len(playerIdList) == 0 {
		return
	}

	for _, playerId := range playerIdList {
		if clientObj, ok := GetClientByPlayerId(playerId); ok {
			clientList = append(clientList, clientObj)
		}
	}

	return
}

// 获取客户端的数量
// 返回值：
// 客户端数量
func GetClientCount() int {
	mutex_client.RLock()
	defer mutex_client.RUnlock()

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
