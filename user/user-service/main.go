package main

import (
	"github.com/micro/go-micro"
	pb "micro-app/user/user-service/proto"
	pbid "micro-app/user/userid-service/proto"
	"context"
	"fmt"
	"strconv"
)

type Users struct {
	userid pbid.UserIdService
}

func (u *Users) GetUserInfo(ctx context.Context,req *pb.Request,resp *pb.Response) error {
	vReq := &pbid.Request{Name:req.Name}

	vResp,err := u.userid.GetUserId(ctx,vReq)

	if err != nil {
		return err
	}


	resp.Getting = "hello   " + req.Name + strconv.Itoa(int(vResp.Id))
	return nil
}



func main() {

	service := micro.NewService(micro.Name("user"))

	service.Init()

	client := pbid.NewUserIdService("userid", service.Client())

	pb.RegisterUsersHandler(service.Server(),&Users{userid:client})
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}
