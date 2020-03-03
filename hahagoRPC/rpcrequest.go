/*
 * @Author: haha_giraffe
 * @Date: 2020-02-18 15:14:53
 * @Description: file content
 */
package hahagoRPC

import (
	"reflect"
)

//请求结构体
type Request struct {
	//请求方法名称
	MethodName string
	//请求参数
	Args interface{}
}

// 返回Args的reflect.Value类型
func (request *Request) MakeArgs(service Service) (reflect.Value, error) {
	return reflect.ValueOf(request.Args), nil
}

//新建一个请求结构体
func NewRequest(methodName string, args interface{}) *Request {
	return &Request{
		MethodName: methodName,
		Args:       args}
}
