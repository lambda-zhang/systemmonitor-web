/* copy from: https://blog.csdn.net/dodod2012/article/details/81744526 */

package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"sync"
	"time"

	s "github.com/lambda-zhang/systemmonitor-web/cron"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Connection struct {
	wsConn    *websocket.Conn // 存放websocket连接
	inChan    chan []byte     // 用于存放数据
	outChan   chan []byte     // 用于读取数据
	closeChan chan byte
	mutex     sync.Mutex
	isClosed  bool // chan是否被关闭
}

// 读取Api
func (conn *Connection) ReadMessage() (data []byte, err error) {
	//select是Go中的一个控制结构，类似于用于通信的switch语句。每个case必须是一个通信操作，要么是发送要么是接收。
	//select随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。一个默认的子句应该总是可运行的。
	select {
	case data = <-conn.inChan:
	case <-conn.closeChan:
		err = errors.New("connection is closed")
	}
	return
}

// 发送Api
func (conn *Connection) WriteMessage(data []byte) (err error) {
	select {
	case conn.outChan <- data:
	case <-conn.closeChan:
		err = errors.New("connection is closed")
	}
	return
}

// 关闭连接的Api
func (conn *Connection) Close() {
	// 线程安全的Close，可以并发多次调用也叫做可重入的Close
	conn.wsConn.Close()
	conn.mutex.Lock()
	if !conn.isClosed {
		// 关闭chan,但是chan只能关闭一次
		close(conn.closeChan)
		conn.isClosed = true
	}
	conn.mutex.Unlock()

}

// 初始化长连接
func InitConnection(wsConn *websocket.Conn) (conn *Connection, err error) {
	conn = &Connection{
		wsConn:    wsConn,
		inChan:    make(chan []byte, 1000),
		outChan:   make(chan []byte, 1000),
		closeChan: make(chan byte, 1),
	}

	// 启动读协程
	go conn.readLoop()

	// 启动写协程
	go conn.writeLoop()

	return
}

// 内部实现
func (conn *Connection) readLoop() {
	var (
		data []byte
		err  error
	)
	for {
		if _, data, err = conn.wsConn.ReadMessage(); err != nil {
			goto ERR
		}
		// 容易阻塞到这里，等待inChan有空闲的位置
		select {
		case conn.inChan <- data:
		case <-conn.closeChan: // closeChan关闭的时候执行
			goto ERR
		}
	}

ERR:
	conn.Close()
}

func (conn *Connection) writeLoop() {
	var (
		data []byte
		err  error
	)
	for {
		select {
		case data = <-conn.outChan:
		case <-conn.closeChan:
			goto ERR
		}
		data = <-conn.outChan
		if err = conn.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}
ERR:
	conn.Close()
}

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  4096, // 读取存储空间大小
		WriteBufferSize: 4096, // 写入存储空间大小
		CheckOrigin: func(r *http.Request) bool {
			return true // 允许跨域
		},
	}
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		wsConn *websocket.Conn
		err    error
		// data []byte
		conn *Connection
		data []byte
	)
	// 完成http应答，在httpheader中放下如下参数
	if wsConn, err = upgrader.Upgrade(w, r, nil); err != nil {
		return // 获取连接失败直接返回
	}

	if conn, err = InitConnection(wsConn); err != nil {
		goto ERR
	}

	go func() {
		var (
			err error
		)
		for {
			// 每隔一秒发送一次心跳
			jsonBytes, _ := json.Marshal(s.Info)
			if err = conn.WriteMessage(jsonBytes); err != nil {
				return
			}
			time.Sleep(5 * time.Second)
		}

	}()

	for {
		if data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		if err = conn.WriteMessage(data); err != nil {
			goto ERR
		}
	}

ERR:
	// 关闭当前连接
}

func WsHandler(c *gin.Context) {
	wsHandler(c.Writer, c.Request)
}
