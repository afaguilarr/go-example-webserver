package users_client

import (
	"context"
	"fmt"

	"github.com/afaguilarr/go-example-webserver/proto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UsersClientHandlerInterface interface {
	GetAddress() string
	CreateConnection() error
	CloseConnection() error
	Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error)
	LogIn(ctx context.Context, req *proto.LogInRequest) (*proto.LogInResponse, error)
	Authenticate(ctx context.Context, req *proto.AuthenticateRequest) (*proto.AuthenticateResponse, error)
	RefreshAccessToken(ctx context.Context, req *proto.RefreshAccessTokenRequest) (*proto.RefreshAccessTokenResponse, error)
	LogOut(ctx context.Context, req *proto.LogOutRequest) (*proto.LogOutResponse, error)
}

type UsersClientHandler struct {
	Host       string
	Port       string
	Connection *grpc.ClientConn
	Client     proto.UsersClient
}

func NewCryptoClientHandler(host, port string) *UsersClientHandler {
	return &UsersClientHandler{
		Host: host,
		Port: port,
	}
}

func (uc *UsersClientHandler) GetAddress() string {
	return fmt.Sprintf("%s:%s", uc.Host, uc.Port)
}

func (uc *UsersClientHandler) CreateConnection() error {
	// Set up a connection to the server.
	conn, err := grpc.Dial(uc.GetAddress(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return errors.Wrap(err, "while connecting to the Users server")
	}
	uc.Connection = conn
	uc.Client = proto.NewUsersClient(conn)
	return nil
}

func (uc *UsersClientHandler) CloseConnection() error {
	err := uc.Connection.Close()
	if err != nil {
		return errors.Wrap(err, "While closing connection with Users server")
	}
	return nil
}

func (uc *UsersClientHandler) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	resp, err := uc.Client.Register(ctx, req)
	if err != nil {
		return resp, errors.Wrap(err, "while calling Register RPC")
	}
	return resp, nil
}

func (uc *UsersClientHandler) LogIn(ctx context.Context, req *proto.LogInRequest) (*proto.LogInResponse, error) {
	resp, err := uc.Client.LogIn(ctx, req)
	if err != nil {
		return resp, errors.Wrap(err, "while calling LogIn RPC")
	}
	return resp, nil
}

func (uc *UsersClientHandler) Authenticate(ctx context.Context, req *proto.AuthenticateRequest) (*proto.AuthenticateResponse, error) {
	resp, err := uc.Client.Authenticate(ctx, req)
	if err != nil {
		return resp, errors.Wrap(err, "while calling Authenticate RPC")
	}
	return resp, nil
}

func (uc *UsersClientHandler) RefreshAccessToken(ctx context.Context, req *proto.RefreshAccessTokenRequest) (*proto.RefreshAccessTokenResponse, error) {
	resp, err := uc.Client.RefreshAccessToken(ctx, req)
	if err != nil {
		return resp, errors.Wrap(err, "while calling RefreshAccessToken RPC")
	}
	return resp, nil
}

func (uc *UsersClientHandler) LogOut(ctx context.Context, req *proto.LogOutRequest) (*proto.LogOutResponse, error) {
	resp, err := uc.Client.LogOut(ctx, req)
	if err != nil {
		return resp, errors.Wrap(err, "while calling LogOut RPC")
	}
	return resp, nil
}
