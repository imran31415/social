package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	pb "social/app/grpc/protos"
	"social/app/grpc/protos/serializers"
)

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.User, error) {
	if user, err := s.r.GetUserByUserName(req.GetUserName()); err == nil {
		if user.Password == req.GetPassword() {
			return serializers.User(user), nil
		} else {
			return nil, fmt.Errorf("invalid user")
		}
	}
	insertedId, err := s.r.InsertUser(serializers.CreateUserReq(req))
	if err != nil {
		return nil, err
	}
	user, err := s.r.GetUserById(insertedId)
	if err != nil {
		return nil, status.Error(codes.Internal, "err GetUserById")
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
			return nil, status.Error(codes.Internal, "err GetUserById")
		}
		return serializers.User(user), nil
	case *pb.GetUserReq_UserName:
		name := req.GetUserName()
		user, err := s.r.GetUserByUserName(name)
		if err != nil {
			return nil, status.Error(codes.Internal, "err GetUserByUserName")
		}
		return serializers.User(user), nil
	default:
		return nil, status.Error(codes.Internal, "err unknown getBy")
	}
}
