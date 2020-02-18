/*
 * @Author: haha_giraffe
 * @Date: 2020-01-30 22:56:50
 * @Description: 连接接口
 */
package hahaiface

import (
	"net"
)

//连接接口
type IConnection interface {
	//启动连接，当前连接开始工作
	Start()
	//停止连接，停止当前连接工作
	Stop()
	//获取当前连接绑定的socket
	GetTCPConnection() *net.TCPConn
	//获取连接ID
	GetConnID() uint32
	//获取TCPServer
	GetTcpServer() IServer
	//获取远程客户端TCP状态（ip，port）
	RemoteAddr() net.Addr
	//发送数据
	SendMsg(msgId uint32, data []byte) error

	//设置连接属性
	SetProperty(key string, value interface{})
	//获取连接属性
	GetProperty(key string) (interface{}, error)
	//删除连接属性
	RemoveProperty(key string)
}

//声明一个处理连接业务的方法，参数分别为连接socket，字节数组，字节数，返回一个error
type HandleFunc func(*net.TCPConn, []byte, int) error
