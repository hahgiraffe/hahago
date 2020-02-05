/*
 * @Author: haha_giraffe
 * @Date: 2020-01-31 16:03:17
 * @Description: 请求封装
 */
package hahaiface

/*
	一个客户端连接请求的封装，包括连接和数据
*/
type IRequest interface {
	//获取连接的方法
	GetConnection() IConnection
	//获取数据的方法
	GetData() []byte
	//获取数据的ID
	GetMsgID() uint32
}
