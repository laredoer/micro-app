package main

import (
	"github.com/micro/go-micro"
	pb "micro-app/src/user-api/proto"
	"log"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.user"),
		)
	service.Init()
	pb.RegisterUserServiceHandler(service.Server(),new(Handler))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
