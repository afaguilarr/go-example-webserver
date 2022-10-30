package business

import (
	"context"

	"github.com/afaguilarr/go-example-webserver/app/src/dao"
	"github.com/afaguilarr/go-example-webserver/proto"
)

type BusinessCryptoHandler interface {
	Encrypt(ctx context.Context, req *proto.EncryptRequest) (*proto.EncryptResponse, error)
	Decrypt(ctx context.Context, req *proto.DecryptRequest) (*proto.DecryptResponse, error)
}

type BusinessCrypto struct {
	DaoSalts dao.DaoSaltsHandler
}

func NewBusinessCrypto(daoSalts dao.DaoSaltsHandler) *BusinessCrypto {
	return &BusinessCrypto{
		DaoSalts: daoSalts,
	}
}

func (bc *BusinessCrypto) Encrypt(ctx context.Context, req *proto.EncryptRequest) (*proto.EncryptResponse, error) {
	return &proto.EncryptResponse{}, nil
}

func (bc *BusinessCrypto) Decrypt(ctx context.Context, req *proto.DecryptRequest) (*proto.DecryptResponse, error) {
	return &proto.DecryptResponse{}, nil
}
