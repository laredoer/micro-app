package main

import (
	pb "micro-app/user/user-service/proto"
	"github.com/gohouse/gorose"
)


type Repository interface {
	Get(id string) (*pb.User,error)
	GetAll() ([]*pb.User,error)
	Create(user *pb.User) error
	GetByEmailAndPassword(user *pb.User) (*pb.User,error)
}

type UserRepository struct {
	db *gorose.Database
}

func (repo *UserRepository) Get(id string) (*pb.User,error)  {
	return nil,nil
}

func (repo *UserRepository) GetAll () ([] *pb.User,error){
	return nil,nil
}

func (repo *UserRepository) Create(user *pb.User) (error){
	return nil
}
func(repo *UserRepository) GetByEmailAndPassword(user *pb.User) (*pb.User,error){
	return nil,nil
}
