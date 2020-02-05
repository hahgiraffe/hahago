/*
 * @Author: haha_giraffe
 * @Date: 2020-01-31 17:14:55
 * @Description: file content
 */
package main

import (
	"fmt"
	"hahago/hahanet"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	fmt.Println("Clinet begin")
	time.Sleep(1 * time.Second)
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
	for {
		var buf []byte = []byte("Client send msgID 1")
		newmsg := hahanet.NewMsgPackage(1, buf)
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
		// _, err = conn.Read(tmp)
		if err != nil {
			fmt.Printf("read error %s\n", err)
		}
		recvdata, err := dp.Unpack(tmp)
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

		if err != nil {
			fmt.Println("dp unpack error ", err)
		}
		fmt.Printf("receiveID : %d, receiveContent : %s, len : %d\n", recvdata.GetMessageID(), string(recvdata.GetMessageData()), recvdata.GetMessageLen())
		time.Sleep(1 * time.Second)

	}
}
