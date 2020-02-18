/*
 * @Author: haha_giraffe
 * @Date: 2020-02-18 15:08:51
 * @Description: file content
 */
package hahagoRPC

import (
	"bytes"
	"encoding/gob"
	"net"
	"reflect"
)

const (
	EachReadBytes = 500
)

//读数据
func ReadData(conn net.Conn) ([]byte, error) {
	finalData := make([]byte, 0)
	for {
		data := make([]byte, EachReadBytes)
		i, err := conn.Read(data)
		if err != nil {
			return nil, err
		}
		finalData = append(finalData, data[:i]...)
		if i < EachReadBytes {
			//小于500表示读取完了
			break
		}
	}
	return finalData, nil
}

//写数据
func WriteData(conn net.Conn, data []byte) (int, error) {
	num, err := conn.Write(data)
	return num, err
}

//GOB编码
func Encode(v interface{}) ([]byte, error) {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	err := enc.Encode(v)
	if err != nil {
		return nil, err
	}
	return network.Bytes(), nil
}

//GOB解码
func Decode(data []byte, v interface{}) error {
	buf := bytes.NewBuffer(data)
	return gob.NewDecoder(buf).Decode(v)
}

//当结构体中有interface的时候需要注册，告诉GOBinterface有可能是该种类型
func RegisterGobArgsType(arg reflect.Type) error {
	args := reflect.New(arg)
	if args.Kind() == reflect.Ptr {
		args = args.Elem()
	}
	gob.Register(args.Interface())
	return nil
}
