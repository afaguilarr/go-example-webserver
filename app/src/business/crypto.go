package business

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"

	"github.com/afaguilarr/go-example-webserver/app/src/dao"
	"github.com/afaguilarr/go-example-webserver/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	SECRET_BYTES  = 32
	PW_SALT_BYTES = 64
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
	secret, err := generateRandomBytes(SECRET_BYTES)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "while generating secret: %s", err.Error())
	}

	block, err := aes.NewCipher(secret)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "while generating cipher: %s", err.Error())
	}

	iv, err := generateRandomBytes(block.BlockSize())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "while generating iv: %s", err.Error())
	}

	cfb := cipher.NewCFBEncrypter(block, iv)

	plainText := []byte(req.UnencryptedValue)
	salts, err := generateRandomBytes(PW_SALT_BYTES)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "while generating salts: %s", err.Error())
	}
	saltedValue := getSaltedValue(plainText, salts)

	cipherText := make([]byte, len(saltedValue))
	cfb.XORKeyStream(cipherText, saltedValue)

	encryptedValue := base64.StdEncoding.EncodeToString(cipherText)
	return &proto.EncryptResponse{
		EncryptedValue: encryptedValue,
	}, nil
}

// generateRandomBytes generates a random bytes slice of the given length
// by using the specialized crypto package
func generateRandomBytes(n int) (randBytes []byte, err error) {
	randBytes = make([]byte, n)
	_, err = io.ReadFull(rand.Reader, randBytes)
	return
}

func getSaltedValue(plainText, salts []byte) []byte {
	saltedValue := make([]byte, len(plainText)+len(salts))
	saltedValue = append(saltedValue, plainText...)
	saltedValue = append(saltedValue, salts...)
	return saltedValue
}

func (bc *BusinessCrypto) Decrypt(ctx context.Context, req *proto.DecryptRequest) (*proto.DecryptResponse, error) {
	return &proto.DecryptResponse{}, nil
}
