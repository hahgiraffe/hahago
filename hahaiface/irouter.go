/*
 * @Author: haha_giraffe
 * @Date: 2020-01-31 16:13:18
 * @Description: 路由抽象
 */
package hahaiface

/*
	路由抽象接口
*/

type IRouter interface {
	//处理Connection业务之前的方法Hook
	PreHandle(request IRequest)
	//处理Connection业务的方法
	Handle(request IRequest)
	//处理Connection业务之后的方法
	PostHandle(requst IRequest)
}
