/*
 * @Author: haha_giraffe
 * @Date: 2020-02-01 17:09:56
 * @Description: file content
 */
package hahanet

type Message struct {
	//消息ID
	Id uint32
	//消息长度
	DataLen uint32
	//消息内容
	Data []byte
}

func (m *Message) GetMessageID() uint32 {
	return m.Id
}
func (m *Message) GetMessageLen() uint32 {
	return m.DataLen
}
func (m *Message) GetMessageData() []byte {
	return m.Data
}

func (m *Message) SetMessageID(id uint32) {
	m.Id = id
}
func (m *Message) SetMessageLen(len uint32) {
	m.DataLen = len
}
func (m *Message) SetMessageData(d []byte) {
	m.Data = d
}

//创建Message方法
func NewMsgPackage(id uint32, data []byte) *Message {
	return &Message{
		Id:      id,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}
