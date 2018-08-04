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
	//作为客户端调用服务端
	userid pbid.UserIdService
}

func (u *Users) GetUserInfo(ctx context.Context,req *pb.Request,resp *pb.Response) error {

	//构建数据
	vReq := &pbid.Request{Name:req.Name}

	//调用服务端函数
	vResp,err := u.userid.GetUserId(ctx,vReq)

	if err != nil {
		return err
	}
	resp.Getting = "hello   " + req.Name + strconv.Itoa(int(vResp.Id))
	return nil
}



func main() {
	//创建一个服务
	service := micro.NewService(
		micro.Name("user"),    //
		)

	service.Init()

	//服务之间通信,使其作为客户端调用服务端,创建一个客户端
	client := pbid.NewUserIdService("userid", service.Client())

	//注册服务
	pb.RegisterUsersHandler(service.Server(),&Users{userid:client})

	//运行服务
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}
