/*
 * @Author: haha_giraffe
 * @Date: 2020-01-30 17:12:40
 * @Description: 接口
 */
package hahaiface

//服务接口
type IServer interface {
	//开启服务
	Start()
	//关闭服务
	Stop()
	//运行服务
	Serve()
	//给客户端连接添加一个路由方法
	AddRouter(msgID uint32, router IRouter)
	//获取连接管理器
	GetConnMgr() IConnManager

	//注册OnConnStart连接调用方法
	SetOnConnStart(func(connection IConnection))
	//注册OnConnStop方法
	SetOnConnStop(func(connection IConnection))
	//调动OnConnStart
	CallOnConnStart(conn IConnection)
	//调用OnConnStop
	CallOnConnStop(conn IConnection)
}
