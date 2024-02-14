package main

import (
	"81jcpd.cn/grpcjchhh/service"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

func grpcServer() {
	//	实现 token认证
	//	实现一个拦截器
	var authInterceptor grpc.UnaryServerInterceptor
	authInterceptor = func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		err = Auth(ctx)
		if err != nil {
			return
		}
		return handler(ctx, req)
	}

	//	创建一个服务端
	rpcServer := grpc.NewServer(grpc.UnaryInterceptor(authInterceptor))
	//	注册一个服务到 rpc
	service.RegisterProdServiceServer(rpcServer, service.ProductService)

	listener, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatal("启动监听出错:", err)
	}
	err = rpcServer.Serve(listener)
	if err != nil {
		log.Fatal("启动服务出错:", err)
	}
	fmt.Println("启动grpc服务端成功")
}

func Auth(ctx context.Context) error {
	//	实际上就是拿到我们需要的 token
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("missing credentials")
	}
	var tokenString string
	if token_, ok := md[TokenKey]; ok {
		tokenString = token_[0]
	}
	//	判断 token是否合法
	if tokenString != "123456" {
		return status.Errorf(codes.Unauthenticated, "token不合法")
	}
	return nil
}
