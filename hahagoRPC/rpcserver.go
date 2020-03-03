/*
 * @Author: haha_giraffe
 * @Date: 2020-02-18 14:57:30
 * @Description: file content
 */
package hahagoRPC

import (
	"reflect"
)

//每种服务相当于一个方法
type Service struct {
	//方法
	Method reflect.Method
	//请求参数
	ArgType reflect.Type
	//返回参数
	ReplyType reflect.Type
}
