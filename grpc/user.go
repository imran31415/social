package grpc

import (
	"context"
	"fmt"
	"log"
	pb "social/app/grpc/protos"
	"social/app/grpc/protos/serializers"
)

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.User, error) {
	insertedId, err := s.r.InsertUser(serializers.CreateUserReq(req))
	if err != nil {
		return nil, err
	}
	user, err := s.r.GetUserById(insertedId)
	if err != nil {
		return nil, err
	}
	log.Println("Successfully created user")
	return serializers.User(user), nil
}

func (s *Server) GetUser(ctx context.Context, req *pb.GetUserReq) (*pb.User, error) {
	switch req.GetGetBy().(type) {
	case *pb.GetUserReq_Id:
		id := req.GetId()
		user, err := s.r.GetUserById(id)
		if err != nil {
			return nil, err
		}
		return serializers.User(user), nil
	case *pb.GetUserReq_UserName:
		name := req.GetUserName()
		user, err := s.r.GetUserByUserName(name)
		if err != nil {
			return nil, err
		}
		return serializers.User(user), nil
	default:
		return nil, fmt.Errorf("unknown req.GetGetBy")
	}
}
