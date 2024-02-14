package main

import (
	"81jcpd.cn/grpcjchhh/client/auth"
	"81jcpd.cn/grpcjchhh/service"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {

	token := &auth.Authentication{
		Token: "123456",
	}
	//conn, err := grpc.Dial(":8002", grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(":8002", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithPerRPCCredentials(token))

	if err != nil {
		log.Fatal("服务端出错，连接不上:", err)
	}
	defer conn.Close()

	prodClient := service.NewProdServiceClient(conn)

	//request := &service.ProductRequest{
	//	ProdId: 123,
	//}
	//stockResponse, err1 := prodClient.GetProductStock(context.Background(), request)
	//if err1 != nil {
	//	log.Fatal("查询库存出错:", err1)
	//}
	//fmt.Println("查询成功:", stockResponse.ProdStock, stockResponse.Data.GetValue()[2:])

	//	--------------------------------------------------------------------------------

	//stream, err := prodClient.UpdateProductStock(context.Background())
	//if err != nil {
	//	log.Fatal("客户端获取流失败", err)
	//}
	//
	//rsp := make(chan struct{}, 1)
	//go prodRequest(stream, rsp)
	//select {
	//case <-rsp:
	//	recv, err2 := stream.CloseAndRecv()
	//	if err2 != nil {
	//		log.Fatal("wwww---", err2)
	//	}
	//	stock := recv.ProdStock
	//	fmt.Println("客户端收到响应：", stock)
	//}

	//	--------------------------------------------------------------------------------

	//request := &service.ProductRequest{
	//	ProdId: 123,
	//}
	//
	//stream, err := prodClient.GetProductStockStream(context.Background(), request)
	//if err != nil {
	//	log.Fatal("获取流出错")
	//}
	//for {
	//	recv, err1 := stream.Recv()
	//	if err1 != nil {
	//		if err1 == io.EOF {
	//			fmt.Println("客户端数据接收完成")
	//			err2 := stream.CloseSend()
	//			if err2 != nil {
	//				log.Fatal("wwdsdwd--", err2)
	//			}
	//			break
	//		}
	//		log.Fatal("www--", err)
	//	}
	//	fmt.Println("客户端收到的流", recv.ProdStock)
	//
	//}

	//	--------------------------------------------------------------------------------

	stream, err := prodClient.WoShiNiMaStream(context.Background())
	if err != nil {
		log.Fatal("获取流失败:", err)
	}
	count := 1
	for {
		//	先发后收
		request := &service.ProductRequest{
			ProdId: 123,
		}
		err2 := stream.Send(request)
		if err2 != nil {
			log.Fatal(err2)
		}
		time.Sleep(time.Second)

		recv, err1 := stream.Recv()
		if err1 != nil {
			log.Fatal(err1)
		}
		fmt.Println("客户端收到消息:", recv.ProdStock+int32(count))
		count++
	}

}

func prodRequest(stream service.ProdService_UpdateProductStockClient, rsp chan struct{}) {
	count := 1
	for {

		request := &service.ProductRequest{
			ProdId: 123,
		}
		err := stream.Send(request)
		if err != nil {
			log.Fatal(err)
		}
		count++
		if count > 10 {
			rsp <- struct{}{}
			break
		}

	}

}
