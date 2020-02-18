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

//服务器类型
// type Server struct {
// 	//方法映射。第一个string是服务名，里面的string是方法名，对应的value就是Service指针
// 	ServiceMap map[string]map[string]*Service
// 	//互斥锁
// 	serviceLock sync.Mutex
// 	//服务类型
// 	ServerType reflect.Type

// 	// Netserver hahaiface.IServer
// }

//注册函数调用
// func (server *Server) Register(obj interface{}) error {
// 	server.serviceLock.Lock()
// 	defer server.serviceLock.Unlock()

// 	//通过obj得到其各个方法，存储在servicesMap中
// 	tp := reflect.TypeOf(obj)
// 	val := reflect.ValueOf(obj)
// 	serviceName := reflect.Indirect(val).Type().Name()
// 	if _, ok := server.ServiceMap[serviceName]; ok {
// 		return errors.New(serviceName + " already registed.")
// 	}

// 	s := make(map[string]*Service)
// 	numMethod := tp.NumMethod()
// 	for m := 0; m < numMethod; m++ {
// 		service := new(Service)
// 		method := tp.Method(m)
// 		mtype := method.Type
// 		mname := method.Name

// 		service.ArgType = mtype.In(1)
// 		service.ReplyType = mtype.In(2)
// 		service.Method = method
// 		s[mname] = service

// 		err := RegisterGobArgsType(service.ArgType)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	server.ServiceMap[serviceName] = s
// 	server.ServerType = reflect.TypeOf(obj)
// 	return nil
// }

// func (server *Server) ServeConn(conn net.Conn) {
// 	// trans := NewTransfer(conn)
// 	for {
// 		// 从conn读数据
// 		data, err := ReadData(conn)
// 		if err != nil {
// 			return
// 		}

// 		// decode
// 		var req Request

// 		err = Decode(data, &req)
// 		if err != nil {
// 			return
// 		}

// 		// 根据MethodName拿到service
// 		methodStr := strings.Split(req.MethodName, ".")
// 		//解码后只有两个参数，第一个是服务名，第二个是方法名
// 		if len(methodStr) != 2 {
// 			return
// 		}
// 		service := server.ServiceMap[methodStr[0]][methodStr[1]]

// 		// 构造argv
// 		argv, err := req.MakeArgs(*service)

// 		// 构造reply
// 		reply := reflect.New(service.ReplyType.Elem())

// 		// 调用对应的函数
// 		function := service.Method.Func
// 		out := function.Call([]reflect.Value{reflect.New(server.ServerType.Elem()), argv, reply})
// 		if out[0].Interface() != nil {
// 			return
// 		}

// 		// encode
// 		replyData, err := Encode(reply.Elem().Interface())
// 		if err != nil {
// 			return
// 		}

// 		// 向conn写数据
// 		_, err = WriteData(conn, replyData)
// 		if err != nil {
// 			return
// 		}
// 	}
// }

// func (server *Server) Server(network, address string) error {
// 	l, err := net.Listen(network, address)
// 	if err != nil {
// 		log.Fatalf("net.Listen tcp :0: %v", err)
// 		return err
// 	}

// 	for {
// 		// 阻塞直到收到一个网络连接
// 		conn, e := l.Accept()
// 		if e != nil {
// 			log.Fatalf("l.Accept: %v", e)
// 		}

// 		//开始工作
// 		go server.ServeConn(conn)
// 	}
// 	// return nil
// }

// func (server *Server) CheckServiceMap() {
// 	var m map[string]map[string]*Service = server.ServiceMap
// 	for k, v := range m {
// 		fmt.Println(k) //这个是服务的名称，如Chs
// 		for k2, v2 := range v {
// 			fmt.Println(k2, v2) //这个是服务所绑定的方法，如 Multiply Add
// 		}
// 	}
// }

// func NewServer() *Server {
// 	return &Server{
// 		ServiceMap:  make(map[string]map[string]*Service),
// 		serviceLock: sync.Mutex{},
// 		// Netserver:   s,
// 	}
// }
