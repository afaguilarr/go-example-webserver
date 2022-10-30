package dao

import "context"

type Salts struct {
	Context, Salts *string
}

type DaoSaltsHandler interface {
	UpsertSalts(ctx context.Context, s *Salts) error
	GetSalts(ctx context.Context, s *Salts) error
}
