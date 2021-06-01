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

type ContentJson struct {
	PostBody string
}

func Test_CreatePost(t *testing.T) {
	t.Parallel()
	testContentStruct := &ContentJson{
		PostBody: "A test post",
	}
	testContent, err := json.Marshal(testContentStruct)
	require.NoError(t, err)
	testContentString := string(testContent)

	testUser := &repo.User{Username: "foo", Password: "foo_pw"}
	type test struct {
		name       string
		testDbName string
		in         *pb.CreatePostReq
		exp        *pb.Post
		expErr     error
	}
	tests := []test{
		{
			name:       "CreatePost successfully creates a post",
			testDbName: "test_CreatePost_1",
			in: &pb.CreatePostReq{
				Content: testContentString,
			},
			exp: &pb.Post{
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
			postIn := tt.in
			// set the post User id to the user that was created
			postIn.UserId = user.Id
			post, err := s.CreatePost(context.TODO(), postIn)
			if tt.expErr != nil {
				require.Error(t, err)
				require.Equal(t, tt.expErr, err)
				return
			}
			require.NoError(t, err)

			// validate the content was serialized correctly
			postContent := post.GetContent()
			post.Content = ""
			require.Equal(t, tt.exp, post)
			postContentSerialized := &ContentJson{}
			require.NoError(t, json.Unmarshal([]byte(postContent), postContentSerialized))
			require.Equal(t, testContentStruct, postContentSerialized)

		})
	}
}
