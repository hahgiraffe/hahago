/*
 * @Author: haha_giraffe
 * @Date: 2020-02-18 17:59:17
 * @Description: file content
 */
package main

import (
	"fmt"
	"hahago/hahagoRPC"
	"hahago/hahaiface"
	"hahago/hahanet"
	"reflect"
	"strings"
)

type Args struct {
	A int
	B string
}

type ReplyArgs struct {
	Replynum int
	Replystr string
}

type ChsInt struct {
	name   string
	age    int
	school string
}

type Chs int

func (c *Chs) Multiply(args Args, reply *ReplyArgs) error {
	obj := ChsInt{
		name:   "hahagiraffe",
		age:    10,
		school: "hust",
	}
	fmt.Printf("get call Args : [%d %s]\n", args.A, args.B)
	(*reply).Replynum = args.A * 5 * obj.age
	(*reply).Replystr = args.B + obj.school
	fmt.Printf("reply : [%d %s]\n", (*reply).Replynum, (*reply).Replystr)
	return nil
}

type RPCRouter struct {
	hahanet.BaseRouter
}

func (rpc *RPCRouter) Handle(request hahaiface.IRequest) {
	data := request.GetData()
	fmt.Println(data)

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
	fmt.Println("chs reply ", reply)
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

func main() {
	s := hahanet.NewServer("chs")
	s.Register(new(Chs))
	s.AddRouter(0, &RPCRouter{})
	s.Serve()
}
