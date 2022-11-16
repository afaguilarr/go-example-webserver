package services

import (
	"context"
	"log"

	"github.com/afaguilarr/go-example-webserver/app/src/business"
	"github.com/afaguilarr/go-example-webserver/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ServicesCrypto struct {
	// This is necessary with the most recent versions of go-grpc
	// https://stackoverflow.com/questions/65079032/grpc-with-mustembedunimplemented-method
	// It auto-implements a mustEmbedUnimplemented*** method necessary to
	// implement the Service's Interface
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

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "got nil request")
	}

	if req.Context == "" {
		return nil, status.Error(codes.InvalidArgument, "context can't be an empty string")
	}

	if len(req.UnencryptedValue) == 0 {
		return nil, status.Error(codes.InvalidArgument, "unencrypted value can't be an empty string")
	}

	resp, err := sc.BusinessCrypto.Encrypt(ctx, req)
	if err != nil {
		return nil, status.Errorf(status.Code(err), "error encrypting provided value: %s", err.Error())
	}

	return resp, nil
}

// Decrypt encrypts a password by using a context string
func (sc *ServicesCrypto) Decrypt(ctx context.Context, req *proto.DecryptRequest) (*proto.DecryptResponse, error) {
	log.Println("Decrypt RPC was called!")

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "got nil request")
	}

	if req.Context == "" {
		return nil, status.Error(codes.InvalidArgument, "context can't be an empty string")
	}

	if len(req.EncryptedValue) == 0 {
		return nil, status.Error(codes.InvalidArgument, "encrypted value can't be an empty string")
	}

	resp, err := sc.BusinessCrypto.Decrypt(ctx, req)
	if err != nil {
		return nil, status.Errorf(status.Code(err), "error decrypting provided value: %s", err.Error())
	}

	return resp, nil
}
