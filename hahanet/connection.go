/*
 * @Author: haha_giraffe
 * @Date: 2020-01-30 23:09:27
 * @Description: 连接实现
 */
package hahanet

import (
	"errors"
	"hahago/hahaiface"
	"hahago/hahautils"
	"io"
	"net"
	"sync"
)

type Connection struct {
	//连接属于的服务器实例
	TcpServer hahaiface.IServer
	//连接socket
	Conn *net.TCPConn
	//ID
	ConnID uint32
	//连接状态
	isClosed bool
	//连接所绑定的业务方法API（用于处理自定义业务，类似函数指针）
	// handleAPI hahaiface.HandleFunc
	//Read告知Write连接是否已经退出
	ExitChan chan bool
	//消息管理 消息ID --> 路由方法
	MsgHandler hahaiface.IMsgHandle
	//无缓冲管道，用于读写之间消息通信
	MsgChan chan []byte
	//连接属性
	property map[string]interface{}
	//保护连接属性的锁
	propertyLock sync.RWMutex
}

//服务器从客户端读数据的方法
func (c *Connection) StartRead() {
	hahautils.HaHalog.Debug("[Read goroutine is running]")
	defer hahautils.HaHalog.Debugf("[Read is ending, connID %d]\n", c.ConnID)
	defer c.Stop()

	for {

		//当有datapack功能之后，接收后需要解包
		dp := NewDataPack()

		//读取消息头部（Len + ID） 8字节
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			if err == io.EOF {
				break
			}
			hahautils.HaHalog.Error("readfull error ", err)
			break
		}

		msg, err := dp.Unpack(headData)
		if err != nil {
			hahautils.HaHalog.Error("unpack error ", err)
			break
		}
		//根据长度读取body数据并放到message中
		var body []byte
		if msg.GetMessageLen() > 0 {
			body = make([]byte, msg.GetMessageLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), body); err != nil {
				hahautils.HaHalog.Error("readbody full error ", err)
				break
			}
		}
		msg.SetMessageData(body)
		//得到当前连接的Request请求(连接 + 数据)
		req := Request{
			conn: c,
			msg:  msg,
		}

		if hahautils.GlobalObject.WorkerPoolSize > 0 {
			//将任务交给已经开启的工作池
			c.MsgHandler.SendMsgToTaskQueue(&req)
		} else {
			//如果没有开始工作池，就开启一个goroutine自己处理
			go c.MsgHandler.DoMsgHandle(&req)
		}
	}
}

//服务器发送给客户端的方法
func (c *Connection) StartWrite() {
	hahautils.HaHalog.Debug("[Writer goroutine running]")
	defer hahautils.HaHalog.Debug(c.RemoteAddr().String(), " [conn Writer exit]")

	for {
		select {
		case data := <-c.MsgChan:
			if _, err := c.Conn.Write(data); err != nil {
				hahautils.HaHalog.Error("send data error ", err)
				return
			}
		case <-c.ExitChan:
			//收到ExitChan管道说明Read时候客户端退出了，所以Write也要退出
			return
		}
	}
}

//启动连接，当前连接开始工作
func (c *Connection) Start() {
	// fmt.Printf("Conn ID %d Start\n", c.ConnID)
	//启动一个用于读的goroutine和一个用于写的goroutine
	go c.StartRead()
	go c.StartWrite()

	//调用创建连接之后需要处理的业务OnConnStart
	c.TcpServer.CallOnConnStart(c)
}

//停止连接，停止当前连接工作
func (c *Connection) Stop() {
	// fmt.Printf("Conn ID %d Stop\n", c.ConnID)
	if c.isClosed {
		return
	}
	c.isClosed = true
	//调用OnConnStop，处理连接关闭之前的业务
	c.TcpServer.CallOnConnStop(c)
	//关闭连接和管道
	c.Conn.Close()
	//这里通过管道告诉write退出
	c.ExitChan <- true

	//将当前连接从ConnectionManager取出
	c.TcpServer.GetConnMgr().Remove(c)
	close(c.ExitChan)
	close(c.MsgChan)
}

//获取当前连接绑定的socket
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

//获取连接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) GetTcpServer() hahaiface.IServer {
	return c.TcpServer
}

//获取远程客户端TCP状态（ip，port）
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

//发送给客户端的数据进行封包并发送
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed == true {
		return errors.New("Connection closed")
	}

	dp := NewDataPack()
	binarymsg, err := dp.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		hahautils.HaHalog.Error("dp pack error ", err)
		return errors.New("dp pack error")
	}

	//通过MsgChan管道发送给Write
	c.MsgChan <- binarymsg
	return nil
}

//设置连接属性
func (c *Connection) SetProperty(key string, value interface{}) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	c.property[key] = value
}

//获取连接属性
func (c *Connection) GetProperty(key string) (interface{}, error) {
	c.propertyLock.RLock()
	defer c.propertyLock.RUnlock()

	if value, ok := c.property[key]; ok {
		return value, nil
	} else {
		return nil, errors.New("no property found")
	}
}

//删除连接属性
func (c *Connection) RemoveProperty(key string) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	delete(c.property, key)
}

func NewConnection(server hahaiface.IServer, conn *net.TCPConn, connID uint32, handler hahaiface.IMsgHandle) *Connection {
	c := &Connection{
		TcpServer: server,
		Conn:      conn,
		ConnID:    connID,
		isClosed:  false,
		// handleAPI: callbackAPI,
		MsgHandler: handler,
		ExitChan:   make(chan bool, 1),
		MsgChan:    make(chan []byte),
		property:   make(map[string]interface{}),
	}
	c.TcpServer.GetConnMgr().Add(c)
	return c
}
