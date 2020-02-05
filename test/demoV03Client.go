/*
 * @Author: haha_giraffe
 * @Date: 2020-01-31 17:14:55
 * @Description: file content
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
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		log.Fatalf("Dial error %s\n", err)
		return
	}
	for {
		var buf []byte = []byte("good morning")
		if _, err := conn.Write([]byte(buf)); err != nil {
			fmt.Printf("write error %s\n", err)
		}
		var tmp []byte = make([]byte, 1024)
		num, err := conn.Read(tmp)
		if err != nil {
			fmt.Printf("read error %s\n", err)
		}
		fmt.Printf("content %s, len %d\n", tmp, num)
		time.Sleep(1 * time.Second)

	}
}
