/*
 * @Author: haha_giraffe
 * @Date: 2020-02-07 20:00:20
 * @Description: Test RPCServer
 */
package main

import (
	"fmt"
	"hahago/hahaiface"
	"hahago/hahanet"
	"hahago/haharpc"
	"strconv"
)

type Method1 struct {
	hahanet.BaseRouter
	args  string
	reply string
}

func (m *Method1) Multiply() {
	intarg, err := strconv.Atoi(m.args)
	if err != nil {
		fmt.Println("strconv.Atoi error ", err)
	}
	m.reply = strconv.Itoa(intarg * 3)
}

func (m *Method1) Handle(request hahaiface.IRequest) {
	fmt.Println("Get ID ", request.GetMsgID(), "Get data : ", string(request.GetData()))
	m.args = string(request.GetData())
	m.Multiply()
	err := request.GetConnection().SendMsg(33, []byte(m.reply))
	if err != nil {
		fmt.Println("Handle Write error ", err)
	}
}

func main() {
	s, err := haharpc.NewRpcServer("chsserver")
	if err != nil {
		fmt.Println("NewRpcserver error ", err)
	}
	s.AddMethod(0, &Method1{})
	s.Start()
}
