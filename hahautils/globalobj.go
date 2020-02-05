/*
 * @Author: haha_giraffe
 * @Date: 2020-01-31 20:32:53
 * @Description: 全局工具文件
 */
package hahautils

import (
	"encoding/json"
	"fmt"
	"hahago/hahaiface"
	"io/ioutil"
)

/*
	服务器JSON配置解析
*/

type GlobalObj struct {
	//Server object
	TcpServer hahaiface.IServer
	//Server ip
	Host string
	//Server port
	TcpPort int
	//Server name
	Name string

	//Version
	Version string
	//max connections
	MaxConn int
	//max data package size
	MaxPackageSize uint32
	//消息队列对象池Worker数量
	WorkerPoolSize uint32
	//每个Worker对象所对应的消息队列中消息个数最大值
	MaxWorkerTaskLen uint32
}

//全局对象
var GlobalObject *GlobalObj

func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("configure/conf.json")
	if err != nil {
		fmt.Println("ioutial.ReadFile error ", err)
		return
	}
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		fmt.Println("json.Unmarshal error ", err)
		return
	}
}

//初始化对象,只要这个文件被导入，init函数就会被main函数之前执行
func init() {
	//默认配置
	GlobalObject = &GlobalObj{
		Name:             "server",
		Version:          "V0.1",
		Host:             "0.0.0.0",
		TcpPort:          9999,
		MaxConn:          1000,
		MaxPackageSize:   4096,
		WorkerPoolSize:   10,
		MaxWorkerTaskLen: 1024,
	}

	//从配置文件中加载参数替换默认值
	GlobalObject.Reload()
}
