package services

import (
	"context"
	"log"

	"github.com/afaguilarr/go-example-webserver/app/src/business"
	"github.com/afaguilarr/go-example-webserver/proto"
)

type ServicesCrypto struct {
	proto.UnimplementedCryptoServer
	BusinessCrypto business.BusinessCryptoHandler
}

func NewServicesCrypto(businessCrypto business.BusinessCryptoHandler) *ServicesCrypto {
	return &ServicesCrypto{
		BusinessCrypto: businessCrypto,
	}
}

// Encrypt encrypts a password by using a context string
func (sc *ServicesCrypto) Encrypt(ctx context.Context, req *proto.EncryptRequest) (*proto.EncryptResponse, error) {
	log.Println("Encrypt RPC was called!")
	return &proto.EncryptResponse{EncryptedValue: "jiji"}, nil
}

// Decrypt encrypts a password by using a context string
func (sc *ServicesCrypto) Decrypt(ctx context.Context, req *proto.DecryptRequest) (*proto.DecryptResponse, error) {
	return &proto.DecryptResponse{DecryptedValue: "jiji"}, nil
}
