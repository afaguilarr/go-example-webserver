package dao

import "context"

type EncryptionData struct {
	Context                 *string
	CryptoSecret, IV, Salts []byte
}

type DaoEncryptionHandler interface {
	UpsertEncryptionData(ctx context.Context, s *EncryptionData) error
	GetEncryptionData(ctx context.Context, s *EncryptionData) error
}
