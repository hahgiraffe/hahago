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
	ping test 自定义路由
*/

type PingRouter struct {
	hahanet.BaseRouter
}

//对三个路由方法进行重写
func (br *PingRouter) PreHandle(request hahaiface.IRequest) {
	fmt.Println("PingRouter PreHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping ...\n"))
	if err != nil {
		fmt.Println("PreHandle Write error ", err)
	}
}

func (br *PingRouter) Handle(request hahaiface.IRequest) {
	fmt.Println("PingRouter Handle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping ping ping ...\n"))
	if err != nil {
		fmt.Println("Handle Write error ", err)
	}
}

func (br *PingRouter) PostHandle(request hahaiface.IRequest) {
	fmt.Println("PingRouter PostHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("After ping ..."))
	if err != nil {
		fmt.Println("PostHandle Write error ", err)
	}
}

func main() {
	s := hahanet.NewServer("[hahanetV0.2]")
	//添加一个自定义Router
	pr := PingRouter{}
	s.AddRouter(&pr)

	s.Serve()
}
