package main

import (
	"context"
	pb "micro-app/user/user-service/proto"
	)


//此handler结构体,必须实现 user.micro.go 中的 UserServiceHandler 接口
type handler struct {
	repo Repository
}

func (h *handler) Create(ctx context.Context,req *pb.User,resp *pb.Response) error {
	if err := h.repo.Create(req); err != nil {
		return nil
	}
	resp.User = req
	return nil
}

func (h *handler) Get(ctx context.Context,req *pb.User,resp *pb.Response)error{
	u, err := h.repo.Get(req.Id)
	if err != nil {
		return err
	}
	resp.User = u
	return nil
}

func (h *handler) GetAll(ctx context.Context,req *pb.Request,resp *pb.Response) error {
	users, err := h.repo.GetAll()
	if err != nil {
		return err
	}

	resp.Users = users
	return nil
}

func (h *handler) Auth(ctx context.Context,req *pb.User,resp *pb.Token) error {
	_, err := h.repo.GetByEmailAndPassword(req)
	if err != nil {
		return nil
	}
	resp.Token = "1233"
	return nil

}

func (h *handler) ValidateToken(ctx context.Context,req *pb.Token,resp *pb.Token) error{
	return nil
}


