package main

import (
	pb "micro-app/user/user-service/proto"
)


type Repository interface {
	Get(id string) (*pb.User,error)
	GetAll() ([]*pb.User,error)
	Create(user *pb.User) error
	GetByEmailAndPassword(user *pb.User) (*pb.User,error)
}
