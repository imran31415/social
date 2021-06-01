package serializers

import (
	"encoding/json"
	pb "social/app/grpc/protos"
	"social/repo"
)

func User(user *repo.User) *pb.User {
	u := &pb.User{
		Id:       user.Id,
		UserName: user.Username,
	}
	if user.Profile != nil {
		u.Profile = string(*user.Profile)
	}
	return u
}

func CreateUserReq(req *pb.CreateUserReq) *repo.User {
	return &repo.User{
		Password: req.GetPassword(),
		Username: req.GetUserName(),
	}
}

func CreatePostReq(req *pb.CreatePostReq) *repo.Post {
	var content json.RawMessage
	if req.GetContent() != "" {
		content = []byte(req.GetContent())
	}
	return &repo.Post{
		Content: &content,
		UserId:  req.GetUserId(),
	}
}
func CreateCommentReq(req *pb.CreateCommentReq) *repo.Comment {
	var content json.RawMessage
	if req.GetContent() != "" {
		content = []byte(req.GetContent())
	}
	return &repo.Comment{
		Content:         &content,
		PostId:          req.GetPostId(),
		ParentCommentId: req.GetParentCommentId(),
	}
}
func CreateFeedItemReq(req *pb.CreateFeedItemReq) *repo.FeedItem {
	return &repo.FeedItem{
		PostId:  req.GetPostId(),
		OwnerId: req.GetOwnerId(),
	}
}
func Comment(comment *repo.Comment) *pb.Comment {
	c := &pb.Comment{
		Id:       comment.Id,
		ParentId: comment.ParentCommentId,
		PostId:   comment.PostId,
	}
	if comment.Content != nil {
		c.Content = string(*comment.Content)
	}
	return c
}
func Post(post *repo.Post) *pb.Post {
	p := &pb.Post{
		Id: post.Id,
	}
	if post.Content != nil {
		p.Content = string(*post.Content)
	}
	return p
}

func Posts(posts *repo.Posts) *pb.Posts {
	out := make([]*pb.Post, 0, len(posts.Items))
	for _, p := range posts.Items {
		out = append(out, Post(p))
	}
	return &pb.Posts{Items: out}
}

func Comments(posts *repo.Comments) *pb.Comments {
	out := make([]*pb.Comment, 0, len(posts.Items))
	for _, p := range posts.Items {
		out = append(out, Comment(p))
	}
	return &pb.Comments{Items: out}
}

func FeedItem(item *repo.FeedItem) *pb.FeedItem {
	return &pb.FeedItem{
		Id:     item.Id,
		PostId: item.PostId,
	}
}
func Feed(items *repo.Feed, posts *pb.Posts) *pb.Feed {
	postMap := map[int64]string{}
	for _, p := range posts.Items {
		if p.GetContent() != "" {
			postMap[p.Id] = string(p.GetContent())
		}
	}
	out := make([]*pb.FeedItem, 0, len(items.Items))
	for _, item := range items.Items {
		fi := FeedItem(item)
		if v, ok := postMap[item.PostId]; ok {
			fi.PostContent = v
			out = append(out)
		}
	}
	return &pb.Feed{Items: out}
}
