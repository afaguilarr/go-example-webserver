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

// Authenticate receives a JWT access token and its username and verifies
// if the access token is valid.
func (su *ServicesUsers) Authenticate(ctx context.Context, req *proto.AuthenticateRequest) (*proto.AuthenticateResponse, error) {
	log.Println("Authenticate RPC was called!")

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "got nil value in the request")
	}

	if req.Username == "" {
		return nil, status.Error(codes.InvalidArgument, "username can't be an empty string")
	}

	if req.AccessToken == "" {
		return nil, status.Error(codes.InvalidArgument, "access token can't be an empty string")
	}

	resp, err := su.BusinessUsers.Authenticate(ctx, req)
	if err != nil {
		return nil, status.Errorf(status.Code(err), "while authenticating user: %s", err.Error())
	}

	return resp, nil
}

// RefreshAccessToken receives and verifies the correctness of a JWT refresh token,
// and returns a new valid JWT access token.
func (su *ServicesUsers) RefreshAccessToken(ctx context.Context, req *proto.RefreshAccessTokenRequest) (*proto.RefreshAccessTokenResponse, error) {
	log.Println("RefreshAccessToken RPC was called!")

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "got nil value in the request")
	}

	if req.Username == "" {
		return nil, status.Error(codes.InvalidArgument, "username can't be an empty string")
	}

	if req.RefreshToken == "" {
		return nil, status.Error(codes.InvalidArgument, "refresh token can't be an empty string")
	}

	resp, err := su.BusinessUsers.RefreshAccessToken(ctx, req)
	if err != nil {
		return nil, status.Errorf(status.Code(err), "while refreshing access token for user: %s", err.Error())
	}

	return resp, nil
}

// LogOut receives the information of a user, and revokes its refresh token secret, so that a new
// Login process has to be made in order to create a new refresh token secret.
func (su *ServicesUsers) LogOut(ctx context.Context, req *proto.LogOutRequest) (*proto.LogOutResponse, error) {
	log.Println("LogOut RPC was called!")

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "got nil value in the request")
	}

	if req.Username == "" {
		return nil, status.Error(codes.InvalidArgument, "username can't be an empty string")
	}

	resp, err := su.BusinessUsers.LogOut(ctx, req)
	if err != nil {
		return nil, status.Errorf(status.Code(err), "while logging user out: %s", err.Error())
	}

	return resp, nil
}
