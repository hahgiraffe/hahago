/*
 * @Author: haha_giraffe
 * @Date: 2020-02-01 18:14:17
 * @Description: Message进行拆包和装包
 */
package hahaiface

type IDataPack interface {
	//获取长度
	GetHeadLen() uint32
	//封装包
	Pack(msg IMessage) ([]byte, error)
	//拆包
	Unpack([]byte) (IMessage, error)
}
