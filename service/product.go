package service

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/anypb"
	"io"
	"log"
	"time"
)

var ProductService = &productService{}

type productService struct {
}

func (p *productService) mustEmbedUnimplementedProdServiceServer() {
	//TODO implement me
	panic("mustEmbedUnimplementedProdServiceServer 错误")
}

func (p *productService) GetProductStock(context context.Context, request *ProductRequest) (*ProductResponse, error) {
	//	实现具体的业务逻辑
	stock := p.GetStockById(request.ProdId)

	content := Content{Msg: "123456"}
	any_, err := anypb.New(&content)
	if err != nil {
		return nil, err
	}
	return &ProductResponse{
		ProdStock: stock,
		Data:      any_,
	}, nil
}

func (p *productService) GetStockById(id int32) int32 {
	return id
}

func (p *productService) UpdateProductStock(stream ProdService_UpdateProductStockServer) error {
	count := 1
	for {
		//	源源不断的去接收客户端发来的消息
		recv, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		fmt.Println("服务端接收到的流", recv.ProdId, count)
		count++
		if count > 10 {
			rsp := &ProductResponse{
				ProdStock: recv.ProdId,
			}
			err1 := stream.SendAndClose(rsp)
			if err1 != nil {
				return err1
			}
			return nil
		}
	}

}

func (p *productService) GetProductStockStream(request *ProductRequest, stream ProdService_GetProductStockStreamServer) error {
	count := 1
	for {
		rsp := &ProductResponse{
			ProdStock: request.ProdId,
		}
		err1 := stream.Send(rsp)
		if err1 != nil {
			return err1
		}
		time.Sleep(time.Second)
		count++
		if count > 5 {
			return nil
		}
	}

}

func (p *productService) WoShiNiMaStream(stream ProdService_WoShiNiMaStreamServer) error {
	count := 1
	for {
		//	服务端先收后发
		recv, err := stream.Recv()
		if err != nil {
			log.Fatal("ww11", err)
			return err
		}
		fmt.Println("服务端收到客户端的消息:", recv.ProdId)
		time.Sleep(time.Second)
		rsp := &ProductResponse{
			ProdStock: recv.ProdId + int32(count),
		}
		count++
		err1 := stream.Send(rsp)
		if err1 != nil {
			log.Fatal("ww22", err1)
			return err1
		}
	}
}
