package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int
	// 在线用户的列表
	OnlineMap map[string]*User
	// map是全局的，我们要加一个读写锁？有关同步的机制都在sync里
	// 多线程同时操纵公共区的时候要加锁
	mapLock sync.RWMutex
	// 消息广播的channel
	Message chan string
}

// NewServer 创建一个server的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

// ListenMessage 监听Message广播消息channel的goroutine，一旦有消息就发送给全部的在线User
func (this *Server) ListenMessage() {
	for {
		msg := <-this.Message
		// 将message发送给全部的在线user
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}

// BroadCast 广播消息的方法
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	this.Message <- sendMsg
}

func (this *Server) handler(conn net.Conn) {
	//...当前连接的业务
	fmt.Println("有用户连接")
	user := NewUser(conn, this) // 传入当前的server
	user.Online()

	// 监听用户是否活跃的channel
	isLive := make(chan bool)

	// 接收客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			read, err := conn.Read(buf)
			if err != nil && err != io.EOF { // 每次读完之后都会有一个eof的标识表示文件的末尾，这里面有非法操作
				fmt.Println("Conn read err:", err)
				return
			}
			if read == 0 {
				user.Offline()
				return
			}
			// 提取用户的消息，去除\n
			msg := string(buf[:read-1])
			fmt.Println("用户输入的msg是：", msg)
			fmt.Println("用户的输入的是：", []byte(msg))
			// 用户针对message进行处理
			user.DoMessage(msg)
			// 用户的任意消息，代表当前的用户是一个活跃的
			isLive <- true
		}
	}()
	// 当前handler阻塞
	for {
		select {
		case <-isLive:
			// 当前用户是活跃的，应该重置定时器
			// 不做任何事情，为了激活select，更新下面的定时器
		case <-time.After(time.Second * 10000):
			// 已经超时，将当前的user强制关闭
			user.SendMsg("you are out!")
			// 销毁用的资源
			close(user.C)
			// 关闭连接
			conn.Close()
			// 退出当前的handler
			return
		}
	}
}

// Start 启动服务器的接口
func (this *Server) Start() {
	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.listen err:", err)
		return
	}
	defer listener.Close()
	// 启动监听Message的goroutine
	go this.ListenMessage()
	for {
		// accept
		conn, err := listener.Accept() // 说明用户上线了
		if err != nil {
			fmt.Println("listener accept err:", err)
			continue
		}
		// do handler，开启一个go程去处理这个业务，这样的话go程立马开启一个协程处理，当前go程立即回来处理下一个accept的请求
		go this.handler(conn)
	}
}
