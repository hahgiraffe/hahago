/*
 * @Author: haha_giraffe
 * @Date: 2020-02-01 21:12:38
 * @Description: datapack的数据封包和拆包单元测试
 */
package hahanet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

func TestDataPack(t *testing.T) {
	/*
		模拟服务器
	*/
	listenner, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("listen error ", err)
		return
	}

	go func() {
		conn, err := listenner.Accept()
		if err != nil {
			fmt.Println("accept error ", err)
		}

		go func(conn net.Conn) {
			//处理客户端请求
			dp := NewDataPack()
			for {
				//读取header
				headData := make([]byte, dp.GetHeadLen())
				_, err := io.ReadFull(conn, headData) //一次读满8字节
				if err != nil {
					fmt.Println("read header error")
					break
				}

				messheader, err := dp.Unpack(headData)
				if err != nil {
					fmt.Println("unpack error ", err)
					return
				}
				if messheader.GetMessageLen() > 0 {
					//则说明数据长度已经读取到（意味着数据包中有数据），则需要继续读取数据
					msg := messheader.(*Message)
					//根据数据长度开辟数据空间
					msg.Data = make([]byte, msg.GetMessageLen())
					//根据长度再次从conn中读取数据
					_, err := io.ReadFull(conn, msg.Data)
					if err != nil {
						fmt.Println("server unpack error ", err)
						return
					}

					fmt.Println("Recv message: ", msg.Id, ", datalen: ", msg.DataLen, ", data: ", string(msg.Data))
				}
			}
		}(conn)
	}()

	/*
		模拟客户端
	*/
	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("net Dial error ", err)
		return
	}
	//第一个包
	dp := NewDataPack()
	mess1 := &Message{
		Id:      1,
		DataLen: 5,
		Data:    []byte{'h', 'e', 'l', 'l', 'o'},
	}
	senddata1, err := dp.Pack(mess1)
	if err != nil {
		fmt.Println("dp pack error ", err)
		return
	}
	//第二个包
	dp2 := NewDataPack()
	mess2 := &Message{
		Id:      2,
		DataLen: 9,
		Data:    []byte{'c', 'h', 's', 'c', 'h', 's', 'c', 'h', 's'},
	}
	senddata2, err := dp2.Pack(mess2)
	if err != nil {
		fmt.Println("dp pack error ", err)
		return
	}

	senddata1 = append(senddata1, senddata2...)
	conn.Write(senddata1)

	// time.Sleep(10 * time.Second)
	select {}
}
