package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc/hello-sever/proto" //pb是别名 方便后面引用
	"net"
)

// hello server 服务端编写
type server struct {
	pb.UnimplementedSayHelloServer
}

// 实现服务逻辑
// 服务器端方法:处理客户端发送的HelloRequest请求接收一个上下文对象ctx和一个HelloRequest对象req作为参数，返回一个HelloResponse对象和一个错误值。
// 功能:打印出"服务端 hello:"加上req.Name的值，并将"Hello "加上req.Name的值作为ResponseMsg字段的值，封装在HelloResponse对象中返回给客户端。
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	fmt.Println("服务端 hello:" + req.Name)
	//给客户端
	return &pb.HelloResponse{ResponseMsg: "Hello " + req.Name}, nil
}

func main() {
	//开启端口
	lsiten, err := net.Listen("tcp", ":9090") //端口不要漏掉冒号
	if err != nil {
		fmt.Printf("无法监听端口： %v\n", err)
		return
	}

	//创建grpc服务
	grpcServer := grpc.NewServer()

	//在grpc服务端中去注册自己编写的服务
	pb.RegisterSayHelloServer(grpcServer, &server{})

	//启动服务
	fmt.Println("服务已启动，监听端口 9090...")
	err = grpcServer.Serve(lsiten)
	if err != nil {
		fmt.Printf("服务启动失败： %v", err)
		return
	}
}
