package main

import (
		"log"
		pb "micro-app/user/user-service/proto"
	"github.com/micro/go-micro"
)

func main() {
	db, err := CreateConnection()
	if err != nil {
		log.Fatalf("connect err: %v\n",err)
	}

	defer db.Close()


	db.AutoMigrate(&pb.User{})

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
