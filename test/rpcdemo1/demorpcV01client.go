/*
 * @Author: haha_giraffe
 * @Date: 2020-02-07 20:11:53
 * @Description: RPCClient
 */
package main

import (
	"fmt"
	"github.com/hahgiraffe/hahago/haharpc"
	)

func main() {
	myclient, err := haharpc.NewRpcClient("127.0.0.1", "8888")
	if err != nil {
		fmt.Println("NewRpcClient error ", err)
	}
	res, err := myclient.Call(0, "44")
	if err != nil {
		fmt.Println("Call error ", err)
	}
	fmt.Println("get result", res)
}
