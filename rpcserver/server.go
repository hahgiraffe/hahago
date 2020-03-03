/*
 * @Author: haha_giraffe
 * @Date: 2020-03-03 20:48:03
 * @Description: file content
 */
package rpcserver

import (
	"fmt"
	"hahago/hahagoRPC"
	"hahago/hahaiface"
	"hahago/hahanet"
	"reflect"
	"strings"
)

//RPCRouter RPC特定的路由方法
type RPCRouter struct {
	hahanet.BaseRouter
}

//RPC路由对应的Handle
func (rpc *RPCRouter) Handle(request hahaiface.IRequest) {
	data := request.GetData()
	// fmt.Println(data)

	var req hahagoRPC.Request
	err := hahagoRPC.Decode(data, &req)
	if err != nil {
		fmt.Println("Decode error ", err)
		return
	}

	methodStr := strings.Split(req.MethodName, ".")
	if len(methodStr) != 2 {
		fmt.Println("methodStr len error")
		return
	}

	m := request.GetConnection().GetTcpServer().GetServiceMap()
	service := m[methodStr[0]][methodStr[1]]

	argv, err := req.MakeArgs(*service)

	reply := reflect.New(service.ReplyType.Elem())

	function := service.Method.Func
	out := function.Call([]reflect.Value{reflect.New(request.GetConnection().GetTcpServer().GetServerType().Elem()), argv, reply})
	if out[0].Interface() != nil {
		fmt.Println("functionCall error ", err)
		return
	}
	// fmt.Println("chs reply ", reply)
	// encode
	replyData, err := hahagoRPC.Encode(reply.Elem().Interface())
	if err != nil {
		fmt.Println("Encode error ", err)
		return
	}

	err = request.GetConnection().SendMsg(202, replyData)
	if err != nil {
		fmt.Println("SendMsg error ", err)
		return
	}
}
