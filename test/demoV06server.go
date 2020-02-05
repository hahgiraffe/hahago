/*
 * @Author: haha_giraffe
 * @Date: 2020-01-31 17:15:18
 * @Description: file content
 */
package main

import (
	"fmt"
	"hahago/hahaiface"
	"hahago/hahanet"
)

/*
	Pingtest 自定义路由
*/

type PingRouter struct {
	hahanet.BaseRouter
}

/*
	HelloRouter
*/
type HelloRouter struct {
	hahanet.BaseRouter
}

//对Handle路由方法进行重写
func (br *PingRouter) Handle(request hahaiface.IRequest) {
	fmt.Println("PingRouter Handle")
	fmt.Println("GetID ", request.GetMsgID(), " Get data : ", string(request.GetData()))
	err := request.GetConnection().SendMsg(202, []byte("Ping...Ping ..."))
	// _, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping ... "))
	if err != nil {
		fmt.Println("PreHandle Write error ", err)
	}
}

func (br *HelloRouter) Handle(request hahaiface.IRequest) {
	fmt.Println("HelloRouter Handle")
	fmt.Println("GetID ", request.GetMsgID(), " Get data : ", string(request.GetData()))
	err := request.GetConnection().SendMsg(201, []byte("hellohello"))
	if err != nil {
		fmt.Println("PreHandle Write error ", err)
	}
}

func DoConnectionBegin(conn hahaiface.IConnection) {
	fmt.Println("DoConnectionBegin called")
	if err := conn.SendMsg(202, []byte("DoConnection Begin")); err != nil {
		fmt.Println(err)
	}
}

func DoConnectionEnd(conn hahaiface.IConnection) {
	fmt.Println("DoConnectionEnd called, connID: ", conn.GetConnID())
}

func main() {
	s := hahanet.NewServer("[hahanetV0.5]")

	//注册hook函数
	s.SetOnConnStart(DoConnectionBegin)
	s.SetOnConnStop(DoConnectionEnd)

	//添加一个自定义Router
	s.AddRouter(0, &PingRouter{})  //如果消息0则回复ping
	s.AddRouter(1, &HelloRouter{}) //如果消息1则回复hello

	s.Serve()
}
