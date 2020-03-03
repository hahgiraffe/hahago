/*
 * @Author: haha_giraffe
 * @Date: 2020-03-03 21:27:36
 * @Description: rpcclient调用
 */
package rpcclient

import (
	"errors"
	"fmt"
	"hahago/hahagoRPC"
	"hahago/hahanet"
	"io"
	"log"
	"net"
	"reflect"
)

//调用RPC，参数分别是地址端口，函数名，请求参数，返回参数的指针
func RPCcall(addrport string, funcname string, req interface{}, reply interface{}) error {

	tcpaddr, err := net.ResolveTCPAddr("tcp", addrport)
	if err != nil {
		fmt.Println("ResolveTcpAddr error ", err)
		return errors.New("ResolveTcpAddr error")
	}
	conn, err := net.DialTCP("tcp", nil, tcpaddr)
	if err != nil {
		log.Fatalf("Dialtcp error %s\n", err)
		return errors.New("DialTCP error")
	}
	defer conn.Close()

	var methodName string = funcname
	//构造一个Request
	request := hahagoRPC.NewRequest(methodName, req)

	//如果是GOB编码，则要注册相应类型，防止gob编解码错误
	//这个先注册请求参数的类型
	err = hahagoRPC.RegisterGobArgsType(reflect.TypeOf(request.Args))
	if err != nil {
		log.Println(err.Error())
		return errors.New("registerGobArgsType error")
	}

	data, err := hahagoRPC.Encode(request)
	if err != nil {
		log.Println(err.Error())
		return errors.New("Encode error")
	}
	newmsg := hahanet.NewMsgPackage(0, data)
	dp := hahanet.NewDataPack()
	senddata, err := dp.Pack(newmsg)
	if err != nil {
		fmt.Println("dp pack error ", err)
		return errors.New("Pack error")
	}
	if _, err := conn.Write(senddata); err != nil {
		fmt.Printf("write error %s\n", err)
		return errors.New("Write error")
	}

	var tmp []byte = make([]byte, 8)
	_, err = io.ReadFull(conn, tmp)
	if err != nil {
		fmt.Printf("read error %s\n", err)
		return errors.New("readFull error")
	}
	recvdata, err := dp.Unpack(tmp)
	if err != nil {
		fmt.Println("dp Unpack error ", err)
		return errors.New("unpack error")
	}
	//根据长度读取body数据并放到message中
	var body []byte
	if recvdata.GetMessageLen() > 0 {
		body = make([]byte, recvdata.GetMessageLen())
		if _, err := io.ReadFull(conn, body); err != nil {
			fmt.Println("readbody full error ", err)
			return errors.New("read body full error")
		}
	}
	recvdata.SetMessageData(body)

	err = hahagoRPC.Decode(body, reply)
	if err != nil {
		fmt.Println("Decode error ", err)
		return errors.New("Decode error")
	}
	// fmt.Printf("receiveID : %d, reply : %v, len : %d\n", recvdata.GetMessageID(), reply, recvdata.GetMessageLen())
	// time.Sleep(2 * time.Second)
	return nil
}
