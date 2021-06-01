package grpc_test

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"social/app/grpc"
	pb "social/app/grpc/protos"
	"social/app/grpc/test_helpers"
	"social/repo"
	"testing"
)

func Test_CreateFeedItem(t *testing.T) {
	t.Parallel()
	testContentStruct := &CommentContentJson{
		CommentBody: "A test comment",
	}
	testContent, err := json.Marshal(testContentStruct)
	require.NoError(t, err)
	testContentString := string(testContent)

	testUser := &repo.User{Username: "foo", Password: "foo_pw"}
	type test struct {
		name       string
		testDbName string
		postIn     *pb.CreatePostReq
		exp        *pb.FeedItem
		expErr     error
	}
	tests := []test{
		{
			name:       "CreateFeedItem successfully creates a feed item",
			testDbName: "test_CreateFeedItem_1",
			postIn: &pb.CreatePostReq{
				Content: testContentString,
			},
			exp: &pb.FeedItem{
				Id: 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := test_helpers.SetupDbForTest(tt.testDbName, "", "../db/schema")
			defer test_helpers.DeleteTestDatabase(tt.testDbName, r)
			require.NoError(t, err)
			require.NoError(t, r.Ping())
			s := grpc.NewServer(r)
			// create a test user for the test
			user, err := s.CreateUser(context.TODO(), &pb.CreateUserReq{
				UserName: testUser.Username,
				Password: testUser.Password,
			})
			require.NoError(t, err)
			user2, err := s.CreateUser(context.TODO(), &pb.CreateUserReq{
				UserName: testUser.Username + "2",
				Password: testUser.Password + "2",
			})
			require.NoError(t, err)
			postIn := tt.postIn
			// set the post User id to the user that was created
			postIn.UserId = user.Id
			post, err := s.CreatePost(context.TODO(), postIn)
			require.NoError(t, err)
			// set expected post id based on created post
			tt.exp.PostId = post.GetId()

			feedItem, err := s.CreateFeedItem(context.TODO(), &pb.CreateFeedItemReq{
				OwnerId: user2.Id,
				PostId:  post.Id,
			})
			if tt.expErr != nil {
				require.Error(t, err)
				require.Equal(t, tt.expErr, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, &pb.FeedItem{
				Id:     1,
				PostId: post.Id,
			}, feedItem)

		})
	}
}
