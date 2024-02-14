package main

import (
	"81jcpd.cn/grpcjchhh/service"
	"fmt"
	"google.golang.org/protobuf/proto"
)

const (
	TokenKey = "token"
)

func main() {

	grpcServer()

}

func testUser() {
	//	1. 创建 user
	user := &service.User{
		Username: "jchhh_1",
		Age:      18,
	}
	//	2. 转成二进制 -- 也就是字节数组 --  也叫序列化过程
	marshal, err := proto.Marshal(user)
	if err != nil {
		panic(err)
	}

	//	3. 反序列化
	newUser := &service.User{}
	err = proto.Unmarshal(marshal, newUser)
	if err != nil {
		panic(err)
	}

	fmt.Println(newUser.String())
}
