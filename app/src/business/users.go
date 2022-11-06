package business

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/afaguilarr/go-example-webserver/app/src/crypto_client"
	"github.com/afaguilarr/go-example-webserver/app/src/dao"
	"github.com/afaguilarr/go-example-webserver/app/src/dao/postgres"
	"github.com/afaguilarr/go-example-webserver/proto"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	accessTokenSecretKey = "ACCESS_TOKEN_SECRET"
)

type BusinessUsersHandler interface {
	Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error)
	LogIn(ctx context.Context, req *proto.LogInRequest) (*proto.LogInResponse, error)
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

func (bu *BusinessUsers) LogIn(ctx context.Context, req *proto.LogInRequest) (*proto.LogInResponse, error) {
	unmatchedUsernameAndPass := status.Error(codes.InvalidArgument, "username and password don't match")

	encryptedPassword, err := bu.DaoUsers.GetPasswordByUsername(ctx, req.Username)
	if encryptedPassword == "" || err == sql.ErrNoRows {
		return nil, unmatchedUsernameAndPass
	}
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("while getting the user password: %s", err.Error()))
	}

	password, err := bu.DecryptPassword(ctx, req.Username, encryptedPassword)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("while decrypting the user password: %s", err.Error()))
	}

	if *password != req.Password {
		return nil, unmatchedUsernameAndPass
	}
	log.Println("User and Password matched!")

	secret := []byte(os.Getenv(accessTokenSecretKey))
	accessToken, err := GenerateJWT(secret, req.Username, 1*time.Minute)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("while generating access JWT token: %s", err.Error()))
	}

	refreshTokenSecret, err := generateRandomBytes(secretNumberOfBytes)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "while generating refresh token secret: %s", err.Error())
	}
	bu.DaoUsers.SetUserRefreshTokenSecret(ctx, req.Username, refreshTokenSecret)

	refreshToken, err := GenerateJWT(refreshTokenSecret, req.Username, 150*time.Hour)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("while generating refresh JWT token: %s", err.Error()))
	}

	return &proto.LogInResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (bu *BusinessUsers) DecryptPassword(ctx context.Context, username, password string) (*string, error) {
	decryptReq := &proto.DecryptRequest{
		Context:        PasswordEncryptionContext(username),
		EncryptedValue: password,
	}

	decryptResp, err := bu.CryptoClient.Decrypt(
		ctx,
		decryptReq,
	)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("while decrypting password: %s", err.Error()))
	}

	return &decryptResp.DecryptedValue, nil
}

func GenerateJWT(secret []byte, username string, d time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(d)
	claims["authorized"] = true
	claims["user"] = username

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", errors.Wrap(err, "while signing JWT")
	}
	return tokenString, nil
}
