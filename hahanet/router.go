/*
 * @Author: haha_giraffe
 * @Date: 2020-01-31 16:15:59
 * @Description: 路由接口的实现
 */
package hahanet

import (
	"github.com/hahgiraffe/hahago/hahaiface"
)

//这里只是一个基类Router，用户使用的时候需要对方法进行重写就可以啦
type BaseRouter struct{}

//这里对BaseRouter的方法都为空方法,目的是可以让后继Router重写三个方法中的某个就可以了，不需要重写所有的
func (br *BaseRouter) PreHandle(request hahaiface.IRequest) {}

func (br *BaseRouter) Handle(request hahaiface.IRequest) {}

func (br *BaseRouter) PostHandle(requst hahaiface.IRequest) {}
