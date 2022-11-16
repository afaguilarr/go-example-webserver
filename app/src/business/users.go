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
	Authenticate(ctx context.Context, req *proto.AuthenticateRequest) (*proto.AuthenticateResponse, error)
	RefreshAccessToken(ctx context.Context, req *proto.RefreshAccessTokenRequest) (*proto.RefreshAccessTokenResponse, error)
	LogOut(ctx context.Context, req *proto.LogOutRequest) (*proto.LogOutResponse, error)
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

func (bu *BusinessUsers) EncryptPassword(ctx context.Context, username, password string) ([]byte, error) {
	encryptReq := &proto.EncryptRequest{
		Context:          PasswordEncryptionContext(username),
		UnencryptedValue: []byte(password),
	}

	encryptResp, err := bu.CryptoClient.Encrypt(
		ctx,
		encryptReq,
	)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("while encrypting password: %s", err.Error()))
	}

	return encryptResp.EncryptedValue, nil
}

func ProtoUserToDaoUser(protoUser *proto.UserInfo, encryptedPassword []byte) *dao.User {
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
	if len(encryptedPassword) == 0 || err == sql.ErrNoRows {
		return nil, unmatchedUsernameAndPass
	}
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("while getting the user password: %s", err.Error()))
	}

	password, err := bu.DecryptPassword(ctx, req.Username, encryptedPassword)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("while decrypting the user password: %s", err.Error()))
	}

	if string(password) != req.Password {
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

func (bu *BusinessUsers) DecryptPassword(ctx context.Context, username string, password []byte) ([]byte, error) {
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

	return decryptResp.DecryptedValue, nil
}

func GenerateJWT(secret []byte, username string, d time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(d).Unix()
	claims["authorized"] = true
	claims["user"] = username

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", errors.Wrap(err, "while signing JWT")
	}
	return tokenString, nil
}

// Struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	ExpectedUsername string `json:"-"`
	Err              error  `json:"-"`
	Username         string `json:"user"`
	Authorized       bool   `json:"authorized"`
	jwt.StandardClaims
}

// Valid validates time based claims "exp", and also validates the expected user for the token.
func (c *Claims) Valid() error {
	now := time.Now().Unix()

	if !c.VerifyExpiresAt(now, false) {
		delta := time.Unix(now, 0).Sub(time.Unix(c.ExpiresAt, 0))
		c.Err = fmt.Errorf("token is expired by %v", delta)
	}

	if c.ExpectedUsername != c.Username {
		c.Err = fmt.Errorf("the token username is not correct")
	}

	return c.Err
}

func (bu *BusinessUsers) Authenticate(ctx context.Context, req *proto.AuthenticateRequest) (*proto.AuthenticateResponse, error) {
	secret := []byte(os.Getenv(accessTokenSecretKey))

	claims := &Claims{ExpectedUsername: req.Username}
	token, err := jwt.ParseWithClaims(req.AccessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, status.Error(codes.Unauthenticated, "invalid JWT token signature")
		}
		return nil, status.Error(codes.Internal, fmt.Sprintf("while parsing the access token: %s", err.Error()))
	}
	if !token.Valid {
		return nil, status.Error(codes.Unauthenticated, fmt.Sprintf("invalid JWT token signature: %s", claims.Err.Error()))
	}
	log.Println("Access Token is valid!")

	return &proto.AuthenticateResponse{}, nil
}

func (bu *BusinessUsers) RefreshAccessToken(ctx context.Context, req *proto.RefreshAccessTokenRequest) (*proto.RefreshAccessTokenResponse, error) {
	encryptedRefreshTokenSecret, err := bu.DaoUsers.GetRefreshTokenSecretByUsername(ctx, req.Username)
	if len(encryptedRefreshTokenSecret) == 0 || err == sql.ErrNoRows {
		// Not mentioning explicitly that the user is non-existent
		return nil, status.Error(codes.Internal, "while getting the user refresh token secret: unknown error occurred")
	}
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("while getting the user refresh token secret: %s", err.Error()))
	}

	claims := &Claims{ExpectedUsername: req.Username}
	token, err := jwt.ParseWithClaims(req.RefreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return encryptedRefreshTokenSecret, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, status.Error(codes.Unauthenticated, "invalid JWT token signature")
		}
		return nil, status.Error(codes.Internal, fmt.Sprintf("while parsing the refresh token: %s", err.Error()))
	}
	if !token.Valid {
		return nil, status.Error(codes.Unauthenticated, fmt.Sprintf("invalid JWT token signature: %s", claims.Err.Error()))
	}
	log.Println("Refresh Token is valid!")

	secret := []byte(os.Getenv(accessTokenSecretKey))
	accessToken, err := GenerateJWT(secret, req.Username, 1*time.Minute)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("while generating access JWT token: %s", err.Error()))
	}

	return &proto.RefreshAccessTokenResponse{
		AccessToken: accessToken,
	}, nil
}

func (bu *BusinessUsers) LogOut(ctx context.Context, req *proto.LogOutRequest) (*proto.LogOutResponse, error) {
	err := bu.DaoUsers.RevokeUserRefreshTokenSecret(ctx, req.Username)
	if err == sql.ErrNoRows {
		// Not mentioning explicitly that the user is non-existent
		return nil, status.Error(codes.Internal, "while revoking the user refresh token secret: unknown error occurred")
	}
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("while revoking the user refresh token secret: %s", err.Error()))
	}

	return &proto.LogOutResponse{}, nil
}
