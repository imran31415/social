package grpc

import (
	"context"
	pb "social/app/grpc/protos"
	"social/app/grpc/protos/serializers"
)

func (s *Server) CreateFeedItem(ctx context.Context, req *pb.CreateFeedItemReq) (*pb.FeedItem, error) {
	insertedId, err := s.r.InsertFeedItem(serializers.CreateFeedItemReq(req))
	if err != nil {
		return nil, err
	}
	item, err := s.r.GetFeedItemById(insertedId)
	if err != nil {
		return nil, err
	}
	return serializers.FeedItem(item), nil
}

func (s *Server) GetFeed(ctx context.Context, req *pb.GetFeedReq) (*pb.Feed, error) {
	feed, err := s.r.GetFeedByOwnerId(req.GetOwnerId())
	if err != nil {
		return nil, err
	}
	return serializers.Feed(feed), nil
}
