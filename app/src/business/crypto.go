package business

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"database/sql"
	"io"

	"github.com/afaguilarr/go-example-webserver/app/src/dao"
	"github.com/afaguilarr/go-example-webserver/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	secretNumberOfBytes = 32
	pwSaltBytes         = 64
)

type BusinessCryptoHandler interface {
	Encrypt(ctx context.Context, req *proto.EncryptRequest) (*proto.EncryptResponse, error)
	Decrypt(ctx context.Context, req *proto.DecryptRequest) (*proto.DecryptResponse, error)
}

type BusinessCrypto struct {
	DaoEncryption dao.DaoEncryptionHandler
}

func NewBusinessCrypto(daoEncryption dao.DaoEncryptionHandler) *BusinessCrypto {
	return &BusinessCrypto{
		DaoEncryption: daoEncryption,
	}
}

func (bc *BusinessCrypto) Encrypt(ctx context.Context, req *proto.EncryptRequest) (*proto.EncryptResponse, error) {
	secret, err := generateRandomBytes(secretNumberOfBytes)
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
	salts, err := generateRandomBytes(pwSaltBytes)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "while generating salts: %s", err.Error())
	}
	saltedValue := getSaltedValue(plainText, salts)

	cipherText := make([]byte, len(saltedValue))
	cfb.XORKeyStream(cipherText, saltedValue)

	resp := &proto.EncryptResponse{
		EncryptedValue: cipherText,
	}

	ed := &dao.EncryptionData{
		Context:      &req.Context,
		CryptoSecret: secret,
		IV:           iv,
		Salts:        salts,
	}
	err = bc.DaoEncryption.UpsertEncryptionData(ctx, ed)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "while storing the encryption data: %s", err.Error())
	}

	return resp, nil
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
	ed := &dao.EncryptionData{
		Context: &req.Context,
	}
	err := bc.DaoEncryption.GetEncryptionData(ctx, ed)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "encryption_data not found")
		}
		return nil, status.Errorf(codes.Internal, "while storing the encryption data: %s", err.Error())
	}

	block, err := aes.NewCipher(ed.CryptoSecret)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "while generating cipher: %s", err.Error())
	}
	cfb := cipher.NewCFBDecrypter(block, ed.IV)

	cipherText := req.EncryptedValue
	saltedValue := make([]byte, len(cipherText))
	cfb.XORKeyStream(saltedValue, cipherText)
	decryptedValue := saltedValue[0:(len(saltedValue) - pwSaltBytes)]

	// For some reason, the decrypted value included an indefinite number of \u0000 characters
	// at the beginning of the slice. Trimming those looks like it fixes the issue.
	decryptedValue = bytes.Trim(decryptedValue, "\u0000")
	resp := &proto.DecryptResponse{
		DecryptedValue: decryptedValue,
	}
	return resp, nil
}
