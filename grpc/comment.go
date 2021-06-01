package grpc

import (
	"context"
	pb "social/app/grpc/protos"
	"social/app/grpc/protos/serializers"
)

func (s *Server) CreateComment(ctx context.Context, req *pb.CreateCommentReq) (*pb.Comment, error) {
	insertedId, err := s.r.InsertComment(serializers.CreateCommentReq(req))
	if err != nil {
		return nil, err
	}
	comment, err := s.r.GetCommentById(insertedId)
	if err != nil {
		return nil, err
	}
	return serializers.Comment(comment), nil
}

func (s *Server) GetComments(ctx context.Context, req *pb.GetCommentsReq) (*pb.Comments, error) {
	comments, err := s.r.GetCommentsByPostId(req.GetPostId())
	if err != nil {
		return nil, err
	}
	return serializers.Comments(comments), nil
}
