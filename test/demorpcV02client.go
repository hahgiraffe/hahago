/*
 * @Author: haha_giraffe
 * @Date: 2020-02-18 17:48:32
 * @Description: file content
 */
package main

import (
	"fmt"
	"hahago/hahagoRPC"
	"net"
)

type Args struct {
	A int
	B string
}

type ReplyArgs struct {
	Replynum int
	Replystr string
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("Dial error ", err)
	}
	defer conn.Close()

	client := hahagoRPC.NewClient(conn)
	defer client.Close()

	var reply ReplyArgs
	args := Args{
		A: 4,
		B: "chschs",
	}
	if err := client.Call("Chs.Multiply", args, &reply); err != nil {
		fmt.Println("NewClient call error ", err)
	}
	fmt.Println("the reply is ", reply)

	var reply2 ReplyArgs
	if err := client.Call("Chs.Add", args, &reply2); err != nil {
		fmt.Println("NewClient call error ", err)
	}
	fmt.Println("the reply is ", reply2)

}
