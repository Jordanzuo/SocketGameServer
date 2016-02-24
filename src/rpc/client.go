package rpc

import (
	"encoding/binary"
	"github.com/Jordanzuo/goutil/intAndBytesUtil"
	"net"
	"strings"
	"sync/atomic"
	"time"
)

const (
	// 包头的长度
	HEADER_LENGTH = 4

	// 定义请求、响应数据的前缀的长度
	ID_LENGTH = 4
)

var (
	// 字节的大小端顺序
	byterOrder = binary.LittleEndian

	// 全局客户端的id，从1开始进行自增
	globalClientId int32 = 0
)

// 获得自增的id值
func getIncrementId() int32 {
	atomic.AddInt32(&globalClientId, 1)

	return globalClientId
}

// 定义客户端对象，以实现对客户端连接的封装
type Client struct {
	// 唯一标识
	id int32

	// 客户端连接对象
	conn net.Conn

	// 接收到的消息内容
	content []byte

	// 上次活跃时间
	activeTime time.Time

	// 玩家id
	PlayerId string

	// 合作商Id
	PartnerId int

	// 服务器Id
	ServerId int

	// 游戏版本号
	GameVersionId int

	// 资源版本号
	ResourceVersionId int
}

// 新建客户端对象
// conn：连接对象
// 返回值：客户端对象的指针
func NewClient(conn net.Conn) *Client {
	return &Client{
		id:                getIncrementId(),
		conn:              conn,
		content:           make([]byte, 0, 1024),
		activeTime:        time.Now(),
		PlayerId:          "",
		PartnerId:         0,
		ServerId:          0,
		GameVersionId:     0,
		ResourceVersionId: 0,
	}
}

// 获取唯一标识
// 返回值：
// 客户端唯一标识
func (clientObj *Client) Id() int32 {
	return clientObj.id
}

// 获取客户端IP地址
// 返回值：
// 客户端IP地址
func (clientObj *Client) IP() string {
	ipAndPort := strings.Split(clientObj.conn.RemoteAddr().String(), ":")
	if len(ipAndPort) > 0 {
		return ipAndPort[0]
	} else {
		return ""
	}
}

// 追加内容
// content：新的内容
// 返回值：无
func (c *Client) AppendContent(content []byte) {
	c.content = append(c.content, content...)
	c.activeTime = time.Now()
}

// 获取有效的消息
// 返回值：
// 消息对应客户端的唯一标识
// 消息内容
// 是否含有有效数据
func (clientObj *Client) GetValidMessage() (int, []byte, bool) {
	// 判断是否包含头部信息
	if len(clientObj.content) < HEADER_LENGTH {
		return 0, nil, false
	}

	// 获取头部信息
	header := clientObj.content[:HEADER_LENGTH]

	// 将头部数据转换为内部的长度
	contentLength := intAndBytesUtil.BytesToInt(header, byterOrder)

	// 判断长度是否满足
	if len(clientObj.content)-HEADER_LENGTH < contentLength {
		return 0, nil, false
	}

	// 提取消息内容
	content := clientObj.content[HEADER_LENGTH : HEADER_LENGTH+contentLength]

	// 将对应的数据截断，以得到新的数据
	clientObj.content = clientObj.content[HEADER_LENGTH+contentLength:]

	// 截取内容的前4位
	idBytes, content := content[:ID_LENGTH], content[ID_LENGTH:]

	// 提取id
	id := intAndBytesUtil.BytesToInt(idBytes, byterOrder)

	return id, content, true
}

// 发送字节数组消息
// id：需要添加到b前发送的数据
// b：待发送的字节数组
func (clientObj *Client) SendByteMessage(id int, b []byte) {
	idBytes := intAndBytesUtil.IntToBytes(id, byterOrder)

	// 将idByte和b合并
	b = append(idBytes, b...)

	// 获得数组的长度
	contentLength := len(b)

	// 将长度转化为字节数组
	header := intAndBytesUtil.IntToBytes(contentLength, byterOrder)

	// 将头部与内容组合在一起
	message := append(header, b...)

	// 发送消息
	clientObj.conn.Write(message)
}

// 判断客户端是否超时
// 返回值：是否超时
func (clientObj *Client) HasExpired() bool {
	return time.Now().Unix() > clientObj.activeTime.Add(ClientExpiredTime()*time.Second).Unix()
}

// 玩家登陆
// playerId：玩家id
// partnerId：合作商Id
// serverId：服务器Id
// gameVersionId：游戏版本Id
// resourceVersionId：资源版本Id
// 返回值：无
func (clientObj *Client) PlayerLogin(playerId string, partnerId, serverId, gameVersionId, resourceVersionId int) {
	clientObj.PlayerId = playerId
	clientObj.PartnerId = partnerId
	clientObj.ServerId = serverId
	clientObj.GameVersionId = gameVersionId
	clientObj.ResourceVersionId = resourceVersionId
}

// 玩家登出
// 返回值：无
func (clientObj *Client) PlayerLogout() {
	clientObj.PlayerId = ""
	clientObj.PartnerId = 0
	clientObj.ServerId = 0
	clientObj.GameVersionId = 0
	clientObj.ResourceVersionId = 0
}

// 退出
// 返回值：无
func (clientObj *Client) Quit() {
	clientObj.conn.Close()
}

// 玩家登出，客户端退出
// 返回值：无
func (clientObj *Client) LogoutAndQuit() {
	clientObj.PlayerLogout()
	clientObj.Quit()
}
