package crypto_client

import (
	"context"
	"fmt"

	"github.com/afaguilarr/go-example-webserver/proto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CryptoClientHandlerInterface interface {
	GetAddress() string
	CreateConnection() error
	CloseConnection() error
	Encrypt(ctx context.Context, req *proto.EncryptRequest) (*proto.EncryptResponse, error)
	Decrypt(ctx context.Context, req *proto.DecryptRequest) (*proto.DecryptResponse, error)
}

type CryptoClientHandler struct {
	Host       string
	Port       string
	Connection *grpc.ClientConn
	Client     proto.CryptoClient
}

func NewCryptoClientHandler(host, port string) *CryptoClientHandler {
	return &CryptoClientHandler{
		Host: host,
		Port: port,
	}
}

func (cc *CryptoClientHandler) GetAddress() string {
	return fmt.Sprintf("%s:%s", cc.Host, cc.Port)
}

func (cc *CryptoClientHandler) CreateConnection() error {
	// Set up a connection to the server.
	conn, err := grpc.Dial(cc.GetAddress(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return errors.Wrap(err, "while connecting to the Crypto server")
	}
	cc.Connection = conn
	cc.Client = proto.NewCryptoClient(conn)
	return nil
}

func (cc *CryptoClientHandler) CloseConnection() error {
	err := cc.Connection.Close()
	if err != nil {
		return errors.Wrap(err, "While closing connection with Crypto server")
	}
	return nil
}

func (cc *CryptoClientHandler) Encrypt(ctx context.Context, req *proto.EncryptRequest) (*proto.EncryptResponse, error) {
	resp, err := cc.Client.Encrypt(ctx, req)
	if err != nil {
		return resp, errors.Wrap(err, "while calling Encrypt RPC")
	}
	return resp, nil
}

func (cc *CryptoClientHandler) Decrypt(ctx context.Context, req *proto.DecryptRequest) (*proto.DecryptResponse, error) {
	resp, err := cc.Client.Decrypt(ctx, req)
	if err != nil {
		return resp, errors.Wrap(err, "while calling Decrypt RPC")
	}
	return resp, nil
}
