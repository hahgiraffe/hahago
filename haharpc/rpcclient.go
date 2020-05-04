/*
 * @Author: haha_giraffe
 * @Date: 2020-02-07 20:50:52
 * @Description: 这是第一个版本的RPC，单纯用多路由实现的
 */
package haharpc

import (
	"errors"
	"fmt"
	"github.com/hahgiraffe/hahago/hahanet"

	"io"
	"net"
)

type RpcClient struct {
	ip   string
	port string
	conn *net.TCPConn
}

func NewRpcClient(ip, port string) (*RpcClient, error) {
	obj := &RpcClient{
		ip:   ip,
		port: port,
	}
	tcpaddr, err := net.ResolveTCPAddr("tcp", ip+":"+port)
	if err != nil {
		fmt.Println("ResolveTcpAddr error ", err)
		return obj, errors.New("ResolveTcpAddr")
	}
	conn, err := net.DialTCP("tcp", nil, tcpaddr)
	if err != nil {
		fmt.Println("Dialtcp error ", err)
		return obj, errors.New("Dialtcp")
	}
	obj.conn = conn
	return obj, nil
}

func (client *RpcClient) Call(funcnum uint32, arg string) (string, error) {
	newmsg := hahanet.NewMsgPackage(funcnum, []byte(arg))
	dp := hahanet.NewDataPack()
	senddata, err := dp.Pack(newmsg)
	if err != nil {
		fmt.Println("dp pack error ", err)
	}
	if _, err := client.conn.Write(senddata); err != nil {
		fmt.Printf("write error %s\n", err)
	}
	var tmp []byte = make([]byte, 8)
	_, err = io.ReadFull(client.conn, tmp)
	if err != nil {
		fmt.Printf("read error %s\n", err)
	}
	recvdata, err := dp.Unpack(tmp)
	//根据长度读取body数据并放到message中
	var body []byte
	if recvdata.GetMessageLen() > 0 {
		body = make([]byte, recvdata.GetMessageLen())
		if _, err := io.ReadFull(client.conn, body); err != nil {
			fmt.Println("readbody full error ", err)
			return "nil", errors.New("readbody full")
		}
	}
	recvdata.SetMessageData(body)
	return string(recvdata.GetMessageData()), nil
	// fmt.Printf("receiveID : %d, receiveContent : %s, len : %d\n", recvdata.GetMessageID(), string(recvdata.GetMessageData()), recvdata.GetMessageLen())
}
