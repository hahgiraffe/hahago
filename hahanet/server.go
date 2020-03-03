/*
 * @Author: haha_giraffe
 * @Date: 2020-01-30 17:11:25
 * @Description: file content
 */
package hahanet

import (
	"errors"
	"fmt"
	"hahago/hahagoRPC"
	"hahago/hahaiface"
	"hahago/hahautils"
	"net"
	"reflect"
	"sync"
)

type Server struct {

	//服务器名称
	Name string
	//服务器绑定ip版本
	IPVersion string
	//服务器监听ip
	IP string
	//服务器监听端口
	Port int
	//消息管理，管理多个router对象
	MsgHandler hahaiface.IMsgHandle
	//连接管理器
	ConnMgr hahaiface.IConnManager
	//创建连接之后调用的Hook函数
	OnConnStart func(conn hahaiface.IConnection)
	//连接关闭之前调用的Hook函数
	OnConnStop func(conn hahaiface.IConnection)

	//RPC相关
	//服务映射表，第一个map的key是服务名称，value又是一个map，这个map的key是函数名称，value是函数类型
	ServiceMap map[string]map[string]*hahagoRPC.Service

	//互斥锁
	serviceLock sync.Mutex

	//服务类型
	ServerType reflect.Type
}

//开启服务
func (s *Server) Start() {
	//socket -> bind -> listen -> accept
	hahautils.HaHalog.Debugf("Server %s start, listen addr at %s, port at %d\n", s.Name, s.IP, s.Port)
	//开启对象池
	s.MsgHandler.StartWorkerPool()

	//开一个goroutine
	go func() {
		//第一步先创建套接字并绑定ip和端口
		tcpaddr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			hahautils.HaHalog.Fatal("ResolveTCPAddr error")
			return
		}
		//第二步开始监听
		listenner, err := net.ListenTCP(s.IPVersion, tcpaddr)
		if err != nil {
			hahautils.HaHalog.Fatalf("ListenTCP error %s\n", err)
		}
		var cid uint32 = 0
		//到这里已经给监听成功，开始阻塞等待连接并处理业务
		for {
			conn, err := listenner.AcceptTCP()
			if err != nil {
				hahautils.HaHalog.Fatal("AcceptTCP error ", err)
				continue
			}
			// fmt.Printf("the %d client is connected\n", ClientNum)

			if s.ConnMgr.Len() >= hahautils.GlobalObject.MaxConn {
				//超过系统规定的最大连接个数
				hahautils.HaHalog.Info("too many connections")
				conn.Close()
				continue
			}

			//接下来可以使用封装好的connection
			dealConn := NewConnection(s, conn, cid, s.MsgHandler)
			cid++
			go dealConn.Start()
		}
	}()
}

//关闭服务
func (s *Server) Stop() {
	//关闭所有连接
	s.ConnMgr.ClearConn()
}

//运行服务
func (s *Server) Serve() {
	//因为Start是非阻塞的，所以要在Serve中阻塞，并且可以处理一些其他业务逻辑
	s.CheckConfig()
	s.Start()

	select {}

}

//添加路由方法
func (s *Server) AddRouter(msgID uint32, router hahaiface.IRouter) {
	s.MsgHandler.AddRouter(msgID, router)
	hahautils.HaHalog.Debug("AddRouter success")
}

//检配置信息
func (s *Server) CheckConfig() {
	//打印配置文件信息检验
	hahautils.HaHalog.Debugf("[Config] ServerName : %s, ServerIPVersion : %s, ServerIP : %s, Server Port : %d", s.Name, s.IPVersion, s.IP, s.Port)
	hahautils.HaHalog.Debugf(" ServerVersion : %s, ServerMaxConnections : %d, ServerMaxDataPackageSize : %d\n", hahautils.GlobalObject.Version,
		hahautils.GlobalObject.MaxConn, hahautils.GlobalObject.MaxPackageSize)
}

//获取连接管理器
func (s *Server) GetConnMgr() hahaiface.IConnManager {
	return s.ConnMgr
}

//获取服务映射表
func (s *Server) GetServiceMap() map[string]map[string]*hahagoRPC.Service {
	return s.ServiceMap
}

//获取服务类型
func (s *Server) GetServerType() reflect.Type {
	return s.ServerType
}

//注册OnConnStart连接调用方法
func (s *Server) SetOnConnStart(hookfunc func(connection hahaiface.IConnection)) {
	s.OnConnStart = hookfunc
}

//注册OnConnStop方法
func (s *Server) SetOnConnStop(hookfunc func(connection hahaiface.IConnection)) {
	s.OnConnStop = hookfunc
}

//调动OnConnStart
func (s *Server) CallOnConnStart(conn hahaiface.IConnection) {
	if s.OnConnStart != nil {
		hahautils.HaHalog.Debug("call OnConnStart")
		s.OnConnStart(conn)
	}
}

//调用OnConnStop
func (s *Server) CallOnConnStop(conn hahaiface.IConnection) {
	if s.OnConnStop != nil {
		hahautils.HaHalog.Debug("call OnConnStop")
		s.OnConnStop(conn)
	}
}

//注册RPC
func (server *Server) Register(obj interface{}) error {
	server.serviceLock.Lock()
	defer server.serviceLock.Unlock()

	//通过obj得到其各个方法，存储在servicesMap中
	tp := reflect.TypeOf(obj)
	val := reflect.ValueOf(obj)
	serviceName := reflect.Indirect(val).Type().Name()
	if _, ok := server.ServiceMap[serviceName]; ok {
		return errors.New(serviceName + " already registed.")
	}

	s := make(map[string]*hahagoRPC.Service)
	numMethod := tp.NumMethod()
	for m := 0; m < numMethod; m++ {
		service := new(hahagoRPC.Service)
		method := tp.Method(m)
		mtype := method.Type
		mname := method.Name

		service.ArgType = mtype.In(1)
		service.ReplyType = mtype.In(2)
		service.Method = method
		s[mname] = service

		err := hahagoRPC.RegisterGobArgsType(service.ArgType)
		if err != nil {
			return err
		}
	}
	server.ServiceMap[serviceName] = s
	server.ServerType = reflect.TypeOf(obj)
	return nil
}

//初始化实例
func NewServer(name string) hahaiface.IServer {
	s := &Server{
		Name:      hahautils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP:        hahautils.GlobalObject.Host,
		Port:      hahautils.GlobalObject.TcpPort,
		// Router:    nil,
		MsgHandler: NewMsgHandler(),
		ConnMgr:    NewConnManager(),

		ServiceMap:  make(map[string]map[string]*hahagoRPC.Service),
		serviceLock: sync.Mutex{},
	}
	return s
}
