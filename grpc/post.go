package grpc

import (
	"context"
	"fmt"
	"log"
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
	created := posts.Items[0]
	// TODO: move this out of user call path
	go s.populateFeed(created.Id)
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

func (s *Server) populateFeed(postId int64) error {
	users, err := s.r.GetAllUsers()
	if err != nil {
		return err
	}
	suc, errs := 0, 0
	for _, user := range users.Items {
		toCreate := &pb.CreateFeedItemReq{
			OwnerId: user.Id,
			PostId:  postId,
		}
		if _, err = s.CreateFeedItem(context.TODO(), toCreate); err != nil {
			log.Println("err inserting feed item: ", err)
			errs += 1
		} else {
			suc += 1
		}
	}
	log.Printf("Successfully inserted %d items", suc)

	if err != nil {
		log.Printf("Failed to insert  %d items", errs)
		return err
	}
	return nil

}
