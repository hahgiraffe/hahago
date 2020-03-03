/*
 * @Author: haha_giraffe
 * @Date: 2020-02-18 15:09:30
 * @Description: file content
 */
package hahagoRPC

import (
	"net"
)

type Client struct {
	conn net.Conn
}

func (client *Client) Close() {
	client.conn.Close()
}

func NewClient(conn net.Conn) *Client {
	return &Client{
		conn: conn}
}
