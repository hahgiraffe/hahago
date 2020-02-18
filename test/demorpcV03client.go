/*
 * @Author: haha_giraffe
 * @Date: 2020-02-18 21:20:35
 * @Description: file content
 */
package main

import (
	"fmt"
	"hahago/hahagoRPC"
	"hahago/hahanet"
	"io"
	"log"
	"net"
	"reflect"
	"time"
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
	// fmt.Println("Clinet begin")
	// time.Sleep(1 * time.Second)
	// conn, err := net.Dial("tcp", "localhost:8888")
	tcpaddr, err := net.ResolveTCPAddr("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("ResolveTcpAddr error ", err)
		return
	}
	conn, err := net.DialTCP("tcp", nil, tcpaddr)
	if err != nil {
		log.Fatalf("Dialtcp error %s\n", err)
		return
	}
	defer conn.Close()

	for {
		var reply ReplyArgs
		// var reply interface{}
		req := Args{
			A: 4,
			B: "chschs",
		}
		var methodName string = "Chs.Multiply"
		//构造一个Request
		request := hahagoRPC.NewRequest(methodName, req)

		//如果是GOB编码，则要注册相应类型，防止gob编解码错误
		//这个先注册请求参数的类型
		err := hahagoRPC.RegisterGobArgsType(reflect.TypeOf(request.Args))
		if err != nil {
			log.Println(err.Error())
			break
		}

		data, err := hahagoRPC.Encode(request)
		if err != nil {
			log.Println(err.Error())
			break
		}
		newmsg := hahanet.NewMsgPackage(0, data)
		dp := hahanet.NewDataPack()
		senddata, err := dp.Pack(newmsg)
		if err != nil {
			fmt.Println("dp pack error ", err)
		}
		if _, err := conn.Write(senddata); err != nil {
			fmt.Printf("write error %s\n", err)
		}

		var tmp []byte = make([]byte, 8)
		_, err = io.ReadFull(conn, tmp)
		if err != nil {
			fmt.Printf("read error %s\n", err)
		}
		recvdata, err := dp.Unpack(tmp)
		if err != nil {
			fmt.Println("dp Unpack error ", err)
		}
		//根据长度读取body数据并放到message中
		var body []byte
		if recvdata.GetMessageLen() > 0 {
			body = make([]byte, recvdata.GetMessageLen())
			if _, err := io.ReadFull(conn, body); err != nil {
				fmt.Println("readbody full error ", err)
				break
			}
		}
		recvdata.SetMessageData(body)

		err = hahagoRPC.Decode(body, &reply)
		if err != nil {
			fmt.Println("Decode error ", err)
			break
		}
		fmt.Printf("receiveID : %d, reply : %v, len : %d\n", recvdata.GetMessageID(), reply, recvdata.GetMessageLen())
		time.Sleep(2 * time.Second)
	}
}
