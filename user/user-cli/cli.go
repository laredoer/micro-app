package main

import (
	"context"
	"fmt"
	pb "micro-app/user/user-service/proto"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro"
	"github.com/micro/micro/cmd"
	"github.com/micro/cli"
	"log"
	"os"
)

func main() {

	cmd.Init()
	//创建 user-service 客户端
	client := pb.NewUserService("go.micro.srv.user", microclient.DefaultClient)
	//设置命令行参数
	service := micro.NewService(

		micro.Name("go.micro.srv.user.client"),
		micro.Flags(
			cli.StringFlag{
				Name: "name",
				Usage: "you full name",
			},
			cli.StringFlag{
				Name: "email",
				Usage: "you email",
			},
			cli.StringFlag{
				Name: "password",
				Usage: "you password",
			},
			cli.StringFlag{
				Name: "company",
				Usage: "you company",
			},
			),
		)

	service.Init(
		micro.Action(func(c *cli.Context) {
			name := c.String("name")
			email := c.String("email")
			password := c.String("password")
			company := c.String("company")

			r,err := client.Create(context.TODO(),&pb.User{
				Name:name,
				Email:email,
				Password:password,
				Company:company,
			})
			if err != nil {
				log.Printf("Could not create: %v",err)
			}
			log.Fatalf("Created: %v",r.User.Id)

			getAll,err := client.GetAll(context.Background(),&pb.Request{})

			if err != nil {
				log.Fatalf("could not list users: %v",err)
			}

			for _, v := range getAll.Users {
				fmt.Println(v)
			}
			os.Exit(0)
		}),
		)
	if err := service.Run(); err != nil{
		log.Println(err)
	}
}
