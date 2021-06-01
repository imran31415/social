package grpc_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"social/app/grpc"
	pb "social/app/grpc/protos"
	"social/app/grpc/test_helpers"
	"testing"
)

func Test_CreateUser(t *testing.T) {
	t.Parallel()
	type test struct {
		name       string
		testDbName string
		in         *pb.CreateUserReq
		exp        *pb.User
		expErr     error
	}
	tests := []test{
		{
			name:       "CreateUser successfully creates a user",
			testDbName: "test_CreateUser_1",
			in: &pb.CreateUserReq{
				UserName: "foo",
				Password: "foopw",
			},
			exp: &pb.User{
				Id:       1,
				UserName: "foo",
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
			user, err := s.CreateUser(context.TODO(), &pb.CreateUserReq{
				UserName: tt.in.GetUserName(),
				Password: tt.in.GetPassword(),
			})
			if tt.expErr != nil {
				require.Error(t, err)
				require.Equal(t, tt.expErr, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.exp, user)

		})
	}
}
