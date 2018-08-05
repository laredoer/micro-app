package main

import (
	"log"
	pb "micro-app/user/user-service/proto"
	"github.com/micro/go-micro"
)

func main() {
	connection, err := CreateConnection()
	if err != nil {
		log.Fatalf("connect err: %v\n",err)
	}

	db := connection.GetInstance()

	defer connection.Close()

	repo := &UserRepository{db}

	service := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
		)
	service.Init()

	pb.RegisterUserServiceHandler(service.Server(),&handler{repo})

	if err := service.Run(); err != nil {
		log.Fatalf("user service error:%v\n",err)
	}



}
