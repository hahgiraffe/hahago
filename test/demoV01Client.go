/*
 * @Author: haha_giraffe
 * @Date: 2020-01-30 20:41:41
 * @Description: V0.1模拟客户端
 */
package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	fmt.Println("Clinet begin")
	time.Sleep(1 * time.Second)
	for {
		conn, err := net.Dial("tcp", "localhost:9999")
		if err != nil {
			log.Fatalf("Dial error %s\n", err)
			return
		}
		var buf []byte = []byte("good morning")
		if _, err := conn.Write([]byte(buf)); err != nil {
			fmt.Printf("write error %s\n", err)
		}
		var tmp []byte = make([]byte, 1024)
		if _, err := conn.Read(tmp); err != nil {
			fmt.Printf("read error %s\n", err)
		}
		fmt.Printf("%s\n", tmp)
		time.Sleep(1 * time.Second)
	}
}
