package main

import (
	"context"
	"fmt"
	"net"
	"rpc/hello_grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

// HelloServer 得有一个结构体，需要实现这个服务的全部方法,叫什么名字不重要
type HelloServer struct {
}

func (HelloServer) SayHello(ctx context.Context, request *hello_grpc.HelloRequest) (*hello_grpc.HelloResponse, error) {
	fmt.Println("入参：", request.Name, request.Message)
	return &hello_grpc.HelloResponse{
		Name:    "server",
		Message: "hello " + request.Name,
	}, nil
}

func main() {
	// 监听端口
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
	}

	// 创建一个gRPC服务器实例。
	s := grpc.NewServer()
	server := HelloServer{}
	// 将server结构体注册为gRPC服务。
	hello_grpc.RegisterHelloServiceServer(s, &server)
	fmt.Println("grpc server running :8080")
	// 开始处理客户端请求。
	err = s.Serve(listen)
}
