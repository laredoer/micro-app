package main

import (
	"github.com/micro/go-micro"
	pb "micro-app/user/userid-service/proto"
	"context"
	"fmt"
)

type UserId struct {}

func (u *UserId) GetUserId(ctx context.Context,req *pb.Request,resp *pb.Response) error {
	resp.Id = 123456
	return nil
}



func main() {

	service := micro.NewService(micro.Name("userid"))

	service.Init()

	pb.RegisterUserIdHandler(service.Server(),new(UserId))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}
