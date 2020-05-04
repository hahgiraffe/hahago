/*
 * @Author: haha_giraffe
 * @Date: 2020-02-03 13:18:37
 * @Description: 多路由消息处理方法
 */
package hahanet

import (
	"github.com/hahgiraffe/hahago/hahaiface"
	"github.com/hahgiraffe/hahago/hahautils"
)

type MsgHandle struct {
	//存放消息ID --> 路由方法（处理方法）
	Apis map[uint32]hahaiface.IRouter
	//消息队列,管道中存请求(每一个消息队列对应一个Worker)
	TaskQueue []chan hahaiface.IRequest
	//当前对象池中worker数量
	WorkerPoolSize uint32
}

//调度并执行Router消息处理方法
func (mhd *MsgHandle) DoMsgHandle(req hahaiface.IRequest) {
	router, ok := mhd.Apis[req.GetMsgID()]
	if !ok {
		hahautils.HaHalog.Error("MsgID not found")
		return
	}
	router.PostHandle(req)
	router.Handle(req)
	router.PostHandle(req)
}

//添加消息ID与路由方法的映射
func (mhd *MsgHandle) AddRouter(msgID uint32, router hahaiface.IRouter) {
	//如果map中已经注册了msgID则不更新
	if _, ok := mhd.Apis[msgID]; ok {
		hahautils.HaHalog.Errorf("msg %d has registed\n", msgID)
		return
	}
	mhd.Apis[msgID] = router
	hahautils.HaHalog.Debugf("msg %d registe success\n", msgID)

}

//启动Worker对象池
func (mh *MsgHandle) StartWorkerPool() {
	hahautils.HaHalog.Debug("WorkerPool start")
	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		mh.TaskQueue[i] = make(chan hahaiface.IRequest, hahautils.GlobalObject.MaxWorkerTaskLen)
		go mh.StartOneWorker(i, mh.TaskQueue[i])
	}
}

//Worker对象池中的每个对象的工作流程
func (mh *MsgHandle) StartOneWorker(workerID int, taskqueue chan hahaiface.IRequest) {
	//阻塞等待消息到来
	hahautils.HaHalog.Debug("WorkerID ", workerID, " is started...")
	for {
		select {
		case req := <-taskqueue:
			//指定request所绑定的业务
			mh.DoMsgHandle(req)
		}
	}
}

//将消息交给工作对象协程的消息队列
func (mh *MsgHandle) SendMsgToTaskQueue(req hahaiface.IRequest) {
	//负载均衡采用轮询分配，根据连接ID分配
	workerID := req.GetConnection().GetConnID() % mh.WorkerPoolSize
	hahautils.HaHalog.Debug("connID ", req.GetConnection().GetConnID(), " send msgID ", req.GetMsgID(), " to taskqueueID ", workerID)
	mh.TaskQueue[workerID] <- req
}

func NewMsgHandler() *MsgHandle {
	return &MsgHandle{
		Apis:           make(map[uint32]hahaiface.IRouter),
		WorkerPoolSize: hahautils.GlobalObject.WorkerPoolSize,
		TaskQueue:      make([]chan hahaiface.IRequest, hahautils.GlobalObject.WorkerPoolSize),
	}
}
