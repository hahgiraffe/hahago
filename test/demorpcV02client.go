/*
 * @Author: haha_giraffe
 * @Date: 2020-02-18 21:20:35
 * @Description: file content
 */
package main

import (
	"fmt"
	"hahago/rpcclient"
)

//调用参数类型（注意要和server注册的参数类型一直）
type Args struct {
	A int
	B string
}

//返回参数类型（注意要和server注册的参数类型一直）
type ReplyArgs struct {
	Replynum int
	Replystr string
}

func main() {

	//请求的ip地址即端口号
	addrport := "127.0.0.1:8888"

	//
	funcname := "Chs.Multiply"

	//
	req := Args{
		A: 4,
		B: "chschs",
	}

	//
	var reply ReplyArgs

	err := rpcclient.RPCcall(addrport, funcname, req, &reply)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("recive Data %v\n", reply)
}
