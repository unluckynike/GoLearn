package main

import (
	"log"
	"net"
	"net/rpc"
)

// rpc服务对象
type HelloService struct{}

// 其中Hello方法必须满足Go语言的RPC规则：
// 方法只能有两个可序列化的参数，其中第二个参数是指针 类型
// 并且返回一个error类型，同时必须是公开的方法。
func (s HelloService) Hello(request string, reply *string) error {
	*reply = "reply: " + request
	return nil
}

func (s HelloService) SayByeBye(request string, reply *string) error {
	*reply = "Bye Bye  " + request
	return nil
}

func main() {
	//rpc.Register 函数调用会将对象类型中所有满足 RPC 规则的对象方法注册为 RPC 函数，所有注册的方法会放在 “HelloService” 服务空间之下。
	rpc.RegisterName("HelloService", new(HelloService))

	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	accept, err := listen.Accept()
	if err != nil {
		log.Fatal("Accept error", err)
	}

	//建立一个唯一的TCP链接，并且通过 rpc.ServeConn 函数在该 TCP 链接上为对方提供 RPC 服务。
	rpc.ServeConn(accept)
}
