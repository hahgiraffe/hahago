/*
 * @Author: haha_giraffe
 * @Date: 2020-01-31 16:04:46
 * @Description: file content
 */
package hahanet

import (
	"github.com/hahgiraffe/hahago/hahaiface"
)

type Request struct {
	//客户端连接
	conn hahaiface.IConnection

	//客户端数据
	msg hahaiface.IMessage
}

func (req *Request) GetConnection() hahaiface.IConnection {
	return req.conn
}

func (req *Request) GetData() []byte {
	return req.msg.GetMessageData()
}

func (req *Request) GetMsgID() uint32 {
	return req.msg.GetMessageID()
}
