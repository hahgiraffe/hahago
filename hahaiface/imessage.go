/*
 * @Author: haha_giraffe
 * @Date: 2020-02-01 17:10:21
 * @Description: 消息的封装
 */
package hahaiface

/*
	消息封装
*/

type IMessage interface {
	GetMessageID() uint32
	GetMessageLen() uint32
	GetMessageData() []byte

	SetMessageID(uint32)
	SetMessageLen(uint32)
	SetMessageData([]byte)
}
