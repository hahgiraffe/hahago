/*
 * @Author: haha_giraffe
 * @Date: 2020-02-18 17:48:15
 * @Description: file content
 */
package main

import (
	"fmt"
	"hahago/hahagoRPC"
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
		age:    33,
		school: "hust",
	}
	fmt.Printf("get call Args : [%d %s]\n", args.A, args.B)
	(*reply).Replynum = args.A * 5 * obj.age
	(*reply).Replystr = args.B + obj.school
	fmt.Printf("reply : [%d %s]\n", (*reply).Replynum, (*reply).Replystr)
	return nil
}

func (c *Chs) Add(args Args, reply *ReplyArgs) error {
	(*reply).Replynum = args.A + 5
	(*reply).Replystr = args.B
	return nil
}

func main() {
	newServer := hahagoRPC.NewServer()

	newServer.Register(new(Chs))

	// newServer.CheckServiceMap()

	newServer.Server("tcp", "127.0.0.1:1234")
}
