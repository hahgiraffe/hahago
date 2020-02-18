/*
 * @Author: haha_giraffe
 * @Date: 2020-02-18 15:09:30
 * @Description: file content
 */
package hahagoRPC

import (
	"log"
	"net"
	"reflect"
)

type Client struct {
	conn net.Conn
}

func (client *Client) Close() {
	client.conn.Close()
}

func (client *Client) Call(methodName string, req interface{}, reply interface{}) error {

	//构造一个Request
	request := NewRequest(methodName, req)

	//如果是GOB编码，则要注册相应类型，防止gob编解码错误
	//这个先注册请求参数的类型
	err := RegisterGobArgsType(reflect.TypeOf(request.Args))
	if err != nil {
		log.Println(err.Error())
		return err
	}

	data, err := Encode(request)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	// write
	_, err = WriteData(client.conn, data)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	// read
	data2, err := ReadData(client.conn)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	// decode and assin to reply
	Decode(data2, reply)

	// return
	return nil
}

func NewClient(conn net.Conn) *Client {
	return &Client{
		conn: conn}
}
