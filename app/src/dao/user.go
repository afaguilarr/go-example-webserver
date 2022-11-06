package dao

import "context"

type User struct {
	Username          string
	Description       *string
	EncryptedPassword *string
	PetMasterInfo     *PetMaster
}

type DaoUsersHandler interface {
	InsertUser(ctx context.Context, u *User) error
	GetPasswordByUsername(ctx context.Context, u string) (string, error)
	SetUserRefreshTokenSecret(ctx context.Context, u string, encryptedRefreshToken []byte) error
}
