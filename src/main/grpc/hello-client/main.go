package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc/hello-sever/proto" //pb是别名 方便后面引用
	"log"
)

//hello 客户端

func main() {
	//连接导server服务端 这里禁用安全传输 没有加密合验证
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("未连接: %v", err)
	}
	defer conn.Close()

	//建立连接
	client := pb.NewSayHelloClient(conn)

	//执行rpc调用 （用这个方法在服务端实现并返回结果）
	hello, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "LinLin"})

	fmt.Println(hello.ResponseMsg)
}
