/*
 * @Author: haha_giraffe
 * @Date: 2020-02-07 20:18:47
 * @Description: 第一个版本的RPC，单纯用多路由实现的RPC
 */
package haharpc

import (
	"github.com/hahgiraffe/hahago/hahaiface"
	"github.com/hahgiraffe/hahago/hahanet"
)

type RpcServer struct {
	server      hahaiface.IServer
	name        string
	functionmap map[string]string
}

func NewRpcServer(name string) (*RpcServer, error) {
	server := &RpcServer{
		name:        name,
		functionmap: make(map[string]string),
	}
	server.server = hahanet.NewServer(server.name)
	return server, nil
}

func (s *RpcServer) Start() {
	s.server.Serve()

}

func (s *RpcServer) AddMethod(funnum int, router hahaiface.IRouter) {
	s.server.AddRouter(0, router)
}
