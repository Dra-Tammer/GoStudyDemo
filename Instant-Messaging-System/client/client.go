package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn // 连接句柄，套接字句柄
	flag       int      // 当前client的模式
}

func NewClient(serverIp string, serverPort int) *Client {
	// 创建客户端对象
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       5418,
	}
	// 链接server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.dial error: ", err)
		return nil
	}
	client.conn = conn
	// 返回对象
	return client
}

// DealResponse 处理server回应的go程，不是go程的话他就是一个串型的，无法处理大量的请求，直接显示到标准输出
func (client *Client) DealResponse() {
	//for {
	//	buf := make()
	//	client.conn.Read(buf)
	//	fmt.Println(buf)
	//}
	// 不断永久的阻塞，下面这是io的方法，当然你也可以不断for循环去读，简写成下面这个样子，os.stdout标准输出
	// 一旦client.conn 有数据，就直接copy到stdout标准输出上，永久阻塞监听
	io.Copy(os.Stdout, client.conn)
}

func (client *Client) menu() bool {
	var flag int
	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("4.查看当前登录的所有的用户")
	fmt.Println("0.退出")
	fmt.Scanln(&flag)
	if flag >= 0 && flag <= 4 {
		client.flag = flag
		return true
	} else {
		fmt.Println(">>> Illegal input! >>>")
		return false
	}
}

func (client *Client) UpdateName() bool {
	fmt.Println(">>> 请输入用户名 >>>")
	fmt.Scanln(&client.Name)
	sendMsg := "rename|" + client.Name + "\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err: ", err)
		return false
	}
	return true
}

func (client *Client) PublicChat() {
	// 提示用户输入消息
	var chatMsg string
	fmt.Println(">>> please input what you want to send:")
	fmt.Scanln(&chatMsg)
	for chatMsg != "exit" {
		// 发送给服务器

		// 消息不为空则发送
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			_, err := client.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("conn.Write error: ", err)
				break
			}
			chatMsg = ""
			fmt.Println(">>> please input what you want to send:")
			fmt.Scanln(&chatMsg)
		}
	}
}

// SelectUsers 查询在线用户
func (client *Client) SelectUsers() {
	sendMsg := "erg\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err: ", err)
		return
	}
}

// PrivateChat 查询当前都有哪些用户在线，提示用户选择一个用户进入私聊
func (client *Client) PrivateChat() {
	var remoteName string
	var chatMsg string
	client.SelectUsers()
	fmt.Println(">>> please input who you want to chat with: ")
	fmt.Scanln(&remoteName)
	for remoteName != "exit" {
		fmt.Println(">>> please input what you want to say to him/her, input 'exit' to exit")
		fmt.Scanln(&chatMsg)
		for chatMsg != "exit" {
			// 消息不为空则发送
			if len(chatMsg) != 0 {
				fmt.Println("你想聊天的对象是：", remoteName)
				sendMsg := "to|" + remoteName + "|" + chatMsg + "\n\n"
				_, err := client.conn.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println("conn.Write error: ", err)
					break
				}
			}
			chatMsg = ""
			fmt.Println(">>> please input what you want to say to him/her, input 'exit' to exit")
			fmt.Scanln(&chatMsg)
		}
		client.SelectUsers()
		fmt.Println(">>> please input who you want to chat with: ")
		fmt.Scanln(&remoteName)
	}
}

// Run 主业务，根据用户的输入的结果选择不同的模式
func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() != true {
		}
		// 根据不同的模式处理不同的业务
		switch client.flag {
		case 1:
			// 公聊模式
			client.PublicChat()
			break
		case 2:
			// 私聊模式
			client.PrivateChat()
			break
		case 3:
			// 更新用户名
			client.UpdateName()
			break
		case 4:
			// 查看当前登录的所有用户
			client.SelectUsers()
			break
		}
	}
}

/* 命令行解析
var serverIp string
var serverPort int

// ./client -ip 127.0.0.1 -port 8888
func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器IP地址（默认是127.0.0.1）")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口（默认是8888）")
}
*/

func main() {
	// 命令行解析
	//flag.Parse()
	client := NewClient("127.0.0.1", 8888)
	if client == nil {
		fmt.Println(">>> Fail to connect! >>>")
		return
	}
	// 单独开启一个goroutine去处理server的回执消息
	go client.DealResponse()
	fmt.Println(">>> Success to connect! >>>")
	// 启动客户端的业务
	client.Run()
}
