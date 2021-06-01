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

type CommentContentJson struct {
	CommentBody string
}

func Test_CreateComment(t *testing.T) {
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
		in         *pb.CreateCommentReq
		postIn     *pb.CreatePostReq
		exp        *pb.Comment
		expErr     error
	}
	tests := []test{
		{
			name:       "CreateComment successfully creates a comment",
			testDbName: "test_CreateComment_1",
			in: &pb.CreateCommentReq{
				Content: testContentString,
			},
			postIn: &pb.CreatePostReq{
				Content: testContentString,
			},
			exp: &pb.Comment{
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
			postIn := tt.postIn
			// set the post User id to the user that was created
			postIn.UserId = user.Id
			post, err := s.CreatePost(context.TODO(), postIn)
			require.NoError(t, err)
			// set expected post id based on created post
			tt.exp.PostId = post.GetId()

			comment, err := s.CreateComment(context.TODO(), &pb.CreateCommentReq{
				PostId:  post.GetId(),
				Content: testContentString,
			})
			if tt.expErr != nil {
				require.Error(t, err)
				require.Equal(t, tt.expErr, err)
				return
			}
			require.NoError(t, err)
			// validate the content was serialized correctly
			commentContent := comment.GetContent()
			comment.Content = ""
			require.Equal(t, tt.exp, comment)
			commentContentSerialized := &CommentContentJson{}
			require.NoError(t, json.Unmarshal([]byte(commentContent), commentContentSerialized))
			require.Equal(t, testContentStruct, commentContentSerialized)

		})
	}
}
