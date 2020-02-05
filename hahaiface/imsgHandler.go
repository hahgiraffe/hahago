/*
 * @Author: haha_giraffe
 * @Date: 2020-02-03 13:14:09
 * @Description: 消息管理（多路由实现）
 */
package hahaiface

type IMsgHandle interface {
	//调度执行相应的路由
	DoMsgHandle(req IRequest)
	//添加路由
	AddRouter(msgID uint32, router IRouter)
	//启动对象池
	StartWorkerPool()
	//将消息交给对象池中的对象处理
	SendMsgToTaskQueue(req IRequest)
}
