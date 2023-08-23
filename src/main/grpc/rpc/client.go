package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	/*
		通过rpc.Dial拨号RPC服务，然后通过client.Call调用具体的RPC方法。
		再调用 client.Call时，第一个参数是用点号链接的RPC服务名字和方法名字，第二和第三个参数分别我们定 义RPC方法的两个参数。
	*/
	client, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing erro： ", err)
	}

	var reply string
	//调用了名为"HelloService.Hello"的远程过程调用（RPC）方法，并传递了一个字符串参数"test_rpc"。该方法的返回值被存储在reply变量中。
	err = client.Call("HelloService.Hello", "test_rpc", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)

	var bye string
	ctx := "RPC(Remote Procedure Call，远程过程调用)是一种计算机通信协议，允许调用不同进程空间的程序。RPC 的客户端和服务器可以在一台机器上，也可以在不同的机器上。程序员使用时，就像调用本地程序一样，无需关注内部的实现细节。"
	client.Call("HelloService.SayByeBye", ctx, &bye)
	fmt.Println(bye)
}
