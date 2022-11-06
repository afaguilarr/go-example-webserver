package services

import (
	"context"
	"log"

	"github.com/afaguilarr/go-example-webserver/app/src/business"
	"github.com/afaguilarr/go-example-webserver/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ServicesUsers struct {
	proto.UnimplementedUsersServer
	BusinessUsers business.BusinessUsersHandler
}

func NewServicesUsers(businessUsers business.BusinessUsersHandler) *ServicesUsers {
	return &ServicesUsers{
		BusinessUsers: businessUsers,
	}
}

// Register registers a user, its pet master associated entity, and its location associated entity
func (su *ServicesUsers) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	log.Println("Register RPC was called!")

	if req == nil || req.UserInfo == nil ||
		req.UserInfo.PetMasterInfo == nil || req.UserInfo.PetMasterInfo.Location == nil {
		return nil, status.Error(codes.InvalidArgument, "got nil value in the request")
	}

	if req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password can't be an empty string")
	}

	if req.UserInfo.Username == "" {
		return nil, status.Error(codes.InvalidArgument, "username can't be an empty string")
	}

	if req.UserInfo.PetMasterInfo.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "pet master name can't be an empty string")
	}

	if req.UserInfo.PetMasterInfo.Location.Country == "" {
		return nil, status.Error(codes.InvalidArgument, "country name can't be an empty string")
	}

	resp, err := su.BusinessUsers.Register(ctx, req)
	if err != nil {
		return nil, status.Errorf(status.Code(err), "while registering user: %s", err.Error())
	}

	return resp, nil
}

// LogIn returns an access token and a refresh token to the client after checking the correctness
// of a username-password pair
func (su *ServicesUsers) LogIn(ctx context.Context, req *proto.LogInRequest) (*proto.LogInResponse, error) {
	log.Println("LogIn RPC was called!")

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "got nil value in the request")
	}

	if req.Username == "" {
		return nil, status.Error(codes.InvalidArgument, "username can't be an empty string")
	}

	if req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password can't be an empty string")
	}

	resp, err := su.BusinessUsers.LogIn(ctx, req)
	if err != nil {
		return nil, status.Errorf(status.Code(err), "while logging user in: %s", err.Error())
	}

	return resp, nil
}
