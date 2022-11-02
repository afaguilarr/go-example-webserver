package dao

import "context"

type User struct {
	Username      string
	Description   *string
	PetMasterInfo *PetMaster
}

type DaoUsersHandler interface {
	InsertUser(ctx context.Context, u *User) error
}
