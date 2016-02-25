package gameserverSocket

import (
	"encoding/binary"
	"github.com/Jordanzuo/goutil/intAndBytesUtil"
	"net"
	"sync/atomic"
	"time"
)

const (
	// 包头的长度
	HEADER_LENGTH = 4
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
}

// 新建客户端对象
// conn：连接对象
// 返回值：客户端对象的指针
func NewClient(conn net.Conn) *Client {
	return &Client{
		id:         getIncrementId(),
		conn:       conn,
		content:    make([]byte, 0, 1024*10),
		activeTime: time.Now(),
	}
}

// 获取唯一标识
func (clientObj *Client) Id() int32 {
	return clientObj.id
}

// 追加内容
// content：新的内容
// 返回值：无
func (clientObj *Client) AppendContent(content []byte) {
	clientObj.content = append(clientObj.content, content...)
	clientObj.activeTime = time.Now()
}

// 获取有效的消息
// 返回值：消息内容
//		：是否含有有效数据
func (clientObj *Client) GetValieMessage() ([]byte, bool) {
	// 判断是否包含头部信息
	if len(clientObj.content) < HEADER_LENGTH {
		return nil, false
	}

	// 获取头部信息
	header := clientObj.content[:HEADER_LENGTH]

	// 将头部数据转换为内部的长度
	contentLength := intAndBytesUtil.BytesToInt(header, byterOrder)

	// 判断长度是否满足
	if len(clientObj.content)-HEADER_LENGTH < contentLength {
		return nil, false
	}

	// 提取消息内容
	content := clientObj.content[HEADER_LENGTH : HEADER_LENGTH+contentLength]

	// 将对应的数据截断，以得到新的数据
	clientObj.content = clientObj.content[HEADER_LENGTH+contentLength:]

	return content, true
}
