/*
 * @Author: haha_giraffe
 * @Date: 2020-02-07 20:18:47
 * @Description: RPCServer
 */
package haharpc

import (
	"hahago/hahaiface"
	"hahago/hahanet"
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
