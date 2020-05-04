/*
 * @Author: haha_giraffe
 * @Date: 2020-02-18 17:59:17
 * @Description: file content
 */
package main

import (
	"fmt"
	"github.com/hahgiraffe/hahago/hahanet"
	"github.com/hahgiraffe/hahago/rpcserver"
)

//请求参数类型
type Args struct {
	A int
	B string
}

//返回参数类型
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

//注册的被调用函数
func (c *Chs) Multiply(args Args, reply *ReplyArgs) error {
	obj := ChsInt{
		name:   "hahagiraffe",
		age:    2,
		school: "hust123123",
	}
	fmt.Printf("get call Args : [%d %s]\n", args.A, args.B)
	(*reply).Replynum = args.A * 5 * obj.age
	(*reply).Replystr = args.B + obj.school
	fmt.Printf("reply : [%d %s]\n", (*reply).Replynum, (*reply).Replystr)
	return nil
}

func main() {
	//新建server实例，配置文件中配置了地址，端口号，最大连接个数等
	s := hahanet.NewServer("chs")

	//注册函数（可以一次性注册多个）
	s.Register(new(Chs))

	//服务器添加rpc实例
	s.AddRouter(0, &rpcserver.RPCRouter{})

	//开启server
	s.Serve()
}
