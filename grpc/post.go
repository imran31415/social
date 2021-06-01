package grpc

import (
	"context"
	"fmt"
	pb "social/app/grpc/protos"
	"social/app/grpc/protos/serializers"
	"social/repo"
)

func (s *Server) CreatePost(ctx context.Context, req *pb.CreatePostReq) (*pb.Post, error) {
	insertedId, err := s.r.InsertPost(serializers.CreatePostReq(req))
	if err != nil {
		return nil, err
	}
	posts, err := s.r.GetPostsIds([]int64{insertedId}, repo.FieldNameSocialPostId)
	if err != nil {
		return nil, err
	}
	if posts.Items == nil || len(posts.Items) != 1 {
		return nil, fmt.Errorf("err unexp items")
	}
	return serializers.Post(posts.Items[0]), nil
}

func (s *Server) GetPosts(ctx context.Context, req *pb.GetPostsReq) (*pb.Posts, error) {
	switch req.GetGetBy() {
	case pb.GetPostsReq_GetPostsIdType_post:
		posts, err := s.r.GetPostsIds(req.GetIds(), repo.FieldNameSocialPostId)
		if err != nil {
			return nil, err
		}
		return serializers.Posts(posts), nil
	case pb.GetPostsReq_GetPostsIdType_user:
		posts, err := s.r.GetPostsIds(req.GetIds(), repo.FieldNameSocialPostUserId)
		if err != nil {
			return nil, err
		}
		return serializers.Posts(posts), nil
	default:
		return nil, fmt.Errorf("invalid req")
	}

}
