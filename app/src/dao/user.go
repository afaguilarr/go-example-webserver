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
	CheckUserByUsername(ctx context.Context, u string) (bool, error)
}
