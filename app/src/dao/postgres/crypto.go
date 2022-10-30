package postgres

import (
	"context"
	"database/sql"
	"log"

	"github.com/afaguilarr/go-example-webserver/app/src/dao"
	"github.com/pkg/errors"
)

type DaoSalts struct {
	DB *sql.DB
}

func NewDaoSalts(db *sql.DB) *DaoSalts {
	return &DaoSalts{
		DB: db,
	}
}

const UpsertSalts = `
INSERT INTO salts (context, salts)
  VALUES ($1, $2)
  RETURNING id, context, salts
ON CONFLICT (context)
  DO
  UPDATE salts SET salts = EXCLUDED.salts, updated_at = NOW()
    WHERE context = EXCLUDED.context;
`

func (d *DaoSalts) UpsertSalts(ctx context.Context, s *dao.Salts) error {
	var id *int

	if s == nil || s.Context == nil || s.Salts == nil || *s.Salts == "" || *s.Context == "" {
		return errors.New("Nil or empty value detected")
	}

	err := d.DB.QueryRowContext(ctx, UpsertSalts, *s.Context, *s.Salts).
		Scan(
			id,
			s.Context,
			s.Salts,
		)
	if err != nil {
		return errors.Wrap(err, "there was an error while inserting the salts")
	}

	log.Println("New salts record was inserted")
	return nil
}

const GetSalts = `
SELECT
  id,
  context,
  salts
FROM salts
  WHERE context = $1;
`

func (d *DaoSalts) GetSalts(ctx context.Context, s *dao.Salts) error {
	var id *int

	if s == nil || s.Context == nil || *s.Context == "" {
		return errors.New("Nil or empty value detected")
	}

	rows, err := d.DB.QueryContext(ctx, GetSalts, *s.Context)
	if err != nil {
		return errors.Wrap(err, "something went wrong while getting the salts")
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		count++
		err = rows.Scan(
			id,
			s.Context,
			s.Salts,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				return err
			}
			return errors.Wrap(err, "there was an error while getting the salts")
		}
	}

	if count == 0 {
		return sql.ErrNoRows
	}
	if count > 1 {
		return errors.New("there were more than one salts records for the provided context")
	}

	log.Println("Succeeded getting the salts record")
	return nil
}
