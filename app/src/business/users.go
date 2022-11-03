package business

import (
	"context"
	"fmt"
	"log"

	"github.com/afaguilarr/go-example-webserver/app/src/crypto_client"
	"github.com/afaguilarr/go-example-webserver/app/src/dao"
	"github.com/afaguilarr/go-example-webserver/app/src/dao/postgres"
	"github.com/afaguilarr/go-example-webserver/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BusinessUsersHandler interface {
	Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error)
}

type BusinessUsers struct {
	DaoUsers     dao.DaoUsersHandler
	CryptoClient crypto_client.CryptoClientHandlerInterface
}

func NewBusinessUsers(
	daoUsers dao.DaoUsersHandler,
	cryptoClient crypto_client.CryptoClientHandlerInterface,
) *BusinessUsers {
	return &BusinessUsers{
		DaoUsers:     daoUsers,
		CryptoClient: cryptoClient,
	}
}

func (bu *BusinessUsers) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	userInfo := req.UserInfo

	encryptedPassword, err := bu.EncryptPassword(ctx, userInfo.Username, req.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("while calling Encrypt RPC: %s", err.Error()))
	}
	log.Println("Password was encrypted successfully!")

	daoUser := ProtoUserToDaoUser(userInfo, encryptedPassword)
	err = bu.DaoUsers.InsertUser(ctx, daoUser)
	if err != nil {
		if err == postgres.UniqueUsernameErr {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}
		return nil, status.Error(codes.Internal, fmt.Sprintf("while inserting the user: %s", err.Error()))
	}
	log.Println("User was inserted successfully!")

	return &proto.RegisterResponse{
		UserInfo: userInfo,
	}, nil
}

func (bu *BusinessUsers) EncryptPassword(ctx context.Context, username, password string) (*string, error) {
	encryptReq := &proto.EncryptRequest{
		Context:          PasswordEncryptionContext(username),
		UnencryptedValue: password,
	}

	encryptResp, err := bu.CryptoClient.Encrypt(
		ctx,
		encryptReq,
	)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("while encrypting password: %s", err.Error()))
	}

	return &encryptResp.EncryptedValue, nil
}

func PasswordEncryptionContext(username string) string {
	return fmt.Sprintf("{username:%s,dataType:password}", username)
}

func ProtoUserToDaoUser(protoUser *proto.UserInfo, encryptedPassword *string) *dao.User {
	var d *string

	petMaster := protoUser.PetMasterInfo
	pm := ProtoPetMasterToDaoPetMaster(petMaster)

	if protoUser.Description != nil {
		d = &protoUser.Description.Value
	}
	return &dao.User{
		Username:          protoUser.Username,
		Description:       d,
		PetMasterInfo:     pm,
		EncryptedPassword: encryptedPassword,
	}
}
