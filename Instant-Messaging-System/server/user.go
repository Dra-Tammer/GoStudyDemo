package main

import (
	"fmt"
	"net"
	"strings"
)

type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn // 客户端连接句柄
	server *Server  // 用户类无法访问到当前的server，当前用户属于哪个server，通过user访问当前server句柄
}

// NewUser 创建一个用户的api
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}
	// 启动监听当前user channel的go程
	go user.ListenMessage()
	return user
}

// ListenMessage 监听当前user channel的方法，一旦有消息，就直接发给对应的客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}

// Online 用户上线的业务
func (this *User) Online() {
	// 用户上线，加入onlinemap中
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()
	// 广播当前用户上线的消息
	this.server.BroadCast(this, "is online")
}

// Offline 用户下线的业务
func (this *User) Offline() {
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()
	// 广播当前用户上线的消息
	this.server.BroadCast(this, "is offline")
}

// SendMsg 给当前用户发消息
func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

// DoMessage 用户处理消息的业务
func (this *User) DoMessage(msg string) {
	fmt.Println("用户输入了： " + msg)
	if msg == "erg" {
		// 用户输入erg，返回所有的在线的用户
		this.server.mapLock.Lock()
		for _, user := range this.server.OnlineMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + ":" + " is online......" + "\n"
			this.SendMsg(onlineMsg)
		}
		this.server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		// 消息格式rename|zhangsan
		newName := strings.Split(msg, "|")[1]
		// 判断name是否存在
		_, ok := this.server.OnlineMap[newName]
		if ok {
			this.SendMsg(">>> This username is already used !\n")
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this
			this.server.mapLock.Unlock()
			this.Name = newName
			this.SendMsg(">>> Update username success !\n")
		}
	} else if len(msg) > 4 && msg[:3] == "to|" {
		// 消息格式，to|zhangsan|消息内容
		// 获取对方的用户名
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			this.SendMsg(">>> err input, please input something as 'to|username|message'")
			return
		}
		// 根据用户名，得到对方的user对象
		remoteUser, ok := this.server.OnlineMap[remoteName]
		if !ok {
			this.SendMsg("this user is not exist")
			return
		}
		// 获取消息内容，通过对方的user对象，将消息内容发送过去
		content := strings.Split(msg, "|")[2]
		if content == "" {
			this.SendMsg("no message, please send once again! ")
			return
		}
		remoteUser.SendMsg(this.Name + "  :  " + content)
	} else {
		this.server.BroadCast(this, msg)
	}
}
