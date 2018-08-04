package main

import (
	pb "micro-app/user/user-service/proto"
	"context"
	api "github.com/micro/micro/api/proto"
	"log"
	"github.com/micro/go-micro/errors"
	"strings"
	"encoding/json"
	"github.com/micro/go-micro"
)

type User struct {
	Client pb.UsersService
}

func (u *User) GetUserInfo(ctx context.Context,request *api.Request,response *api.Response ) error {
	log.Print("Received User.GetUserInfo API request")

	name, ok := request.Get["name"]
	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("user.api.getuserinfo", "Name cannot be blank")
	}

	resp, err := u.Client.GetUserInfo(ctx,&pb.Request{Name: strings.Join(name.Values," ")})
	if err != nil {
		return err
	}

	response.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"message": resp.Getting,
	})
	response.Body = string(b)

	return nil


}


func main(){
	service := micro.NewService(micro.Name("user.api.getuserinfo"))

	service.Init()

	service.Server().Handle(
			service.Server().NewHandler(
					&User{Client: pb.NewUsersService("user",service.Client())},
				),
		)
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}