package main

import (
	"encoding/json"
	"fmt"
	"geerpc"
	"geerpc/codec"
	"log"
	"net"
	"time"
)

func startServer(addr chan string) {
	// pick a free port
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("network error:", err)
	}
	log.Println("start rpc server on", l.Addr())
	//信道 addr，确保服务端端口监听成功，客户端再发起请求。
	addr <- l.Addr().String()
	geerpc.Accept(l)
}

/*
已经实现了一个消息的编解码器 GobCodec，
并且客户端与服务端实现了简单的协议交换(protocol exchange)，即允许客户端使用不同的编码方式。
同时实现了服务端的雏形，建立连接，读取、处理并回复客户端的请求。
*/

func main() {
	addr := make(chan string)
	go startServer(addr)

	// in fact, following code is like a simple geerpc client
	//1.  net.Dial("tcp", <-addr) ：使用TCP协议连接到  <-addr  中指定的地址。
	//2.  conn, _ := ：将连接的结果存储在  conn  变量中。
	conn, _ := net.Dial("tcp", <-addr)
	//3.  defer func() { _ = conn.Close() }() ：在函数结束时执行  conn.Close() ，即关闭连接。
	//使用  defer  关键字可以确保在函数结束时始终执行该语句，无论函数是否出现错误。
	defer func() { _ = conn.Close() }()

	time.Sleep(time.Second)
	// send options
	//客户端首先发送 Option 进行协议交换，接下来发送消息头 h := &codec.Header{}，和消息体 geerpc req ${h.Seq}。
	_ = json.NewEncoder(conn).Encode(geerpc.DefaultOption)
	cc := codec.NewGobCodec(conn)
	// send request & receive response
	for i := 0; i < 10; i++ {
		h := &codec.Header{
			ServiceMethod: "Foo.Sum",
			Seq:           uint64(i),
		}
		_ = cc.Write(h, fmt.Sprintf("geerpc req %d", h.Seq))
		_ = cc.ReadHeader(h)
		//最后解析服务端的响应 reply，并打印出来。
		var reply string
		_ = cc.ReadBody(&reply)
		log.Println("reply:", reply)
	}
}
