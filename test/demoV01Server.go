/*
 * @Author: haha_giraffe
 * @Date: 2020-01-30 17:24:06
 * @Description: V0.1模拟服务器，基于网络库进行测试
 */
package main

import (
	"hahago/hahanet"
)

func main() {
	//获得一个服务器句柄
	s := hahanet.NewServer("chs")
	//服务器开始运行
	s.Serve()
}
