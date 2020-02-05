/*
 * @Author: haha_giraffe
 * @Date: 2020-02-04 12:00:24
 * @Description: 连接管理
 */
package hahanet

import (
	"errors"
	"fmt"
	"hahago/hahaiface"
	"sync"
)

type ConnManager struct {
	//连接ID --> 连接映射
	connections map[uint32]hahaiface.IConnection
	//保护连接集合的读写锁
	connLock sync.RWMutex
}

func (cm *ConnManager) Add(conn hahaiface.IConnection) {
	//加上写锁
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	cm.connections[conn.GetConnID()] = conn
	fmt.Printf("connection %d added \n", conn.GetConnID())
}

func (cm *ConnManager) Remove(conn hahaiface.IConnection) {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	delete(cm.connections, conn.GetConnID())
	fmt.Printf("connection %d deleted \n", conn.GetConnID())
}

func (cm *ConnManager) Get(connID uint32) (hahaiface.IConnection, error) {
	cm.connLock.RLock()
	defer cm.connLock.RUnlock()

	if conn, ok := cm.connections[connID]; ok {
		return conn, nil
	} else {
		return nil, errors.New("connection not found")
	}
}

func (cm *ConnManager) Len() int {
	return len(cm.connections)
}

func (cm *ConnManager) ClearConn() {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	for connID, conn := range cm.connections {
		conn.Stop()
		delete(cm.connections, connID)
	}
	fmt.Printf("clear all connections")
}

//创建连接管理初始化对象
func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]hahaiface.IConnection),
	}
}
