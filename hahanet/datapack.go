/*
 * @Author: haha_giraffe
 * @Date: 2020-02-01 18:15:03
 * @Description: TLV格式拆包与封包
 */
package hahanet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"hahago/hahaiface"
	"hahago/hahautils"
)

type DataPack struct{}

//初始化一个实例
func NewDataPack() *DataPack {
	return &DataPack{}
}

func (d *DataPack) GetHeadLen() uint32 {
	//4字节长度 + 4字节ID
	return 8
}

func (d *DataPack) Pack(msg hahaiface.IMessage) ([]byte, error) {
	//新建一个byte缓冲
	dataBuff := bytes.NewBuffer([]byte{})

	//写长度到缓冲
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMessageLen()); err != nil {
		return nil, err
	}

	//写ID到缓冲
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMessageID()); err != nil {
		return nil, err
	}

	//写数据到缓冲
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMessageData()); err != nil {
		return nil, err
	}
	return dataBuff.Bytes(), nil
}

func (d *DataPack) Unpack(binarydata []byte) (hahaiface.IMessage, error) {
	//二进制的ioReader
	dataBuff := bytes.NewReader(binarydata)
	msg := &Message{}
	//从ioReader中读取datalen
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}
	//读取ID
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}
	//检测读取到的数据是否超过规定最大数据包大小
	if hahautils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > hahautils.GlobalObject.MaxPackageSize {
		return nil, errors.New("too large data")
	}
	//最后可以通过读取的datalen得到数据即可
	return msg, nil
}
