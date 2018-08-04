package main

import (
	"context"
	"fmt"
	pb "micro-app/user/user-service/proto"

	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(micro.Name("user.client"))

	service.Init()

	client := pb.NewUsersService("user", service.Client())

	rsp, err := client.GetUserInfo(context.TODO(), &pb.Request{Name: "tom"})

	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(rsp)
}
