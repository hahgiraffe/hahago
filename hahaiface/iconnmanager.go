/*
 * @Author: haha_giraffe
 * @Date: 2020-02-04 11:59:52
 * @Description: 连接管理
 */
package hahaiface

type IConnManager interface {
	//添加连接
	Add(conn IConnection)
	//删除连接
	Remove(conn IConnection)
	//根据connID获取连接
	Get(connID uint32) (IConnection, error)
	//获取已经连接的个数
	Len() int
	//清楚并终止所有连接
	ClearConn()
}
