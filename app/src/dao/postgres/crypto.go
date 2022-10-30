package postgres

import (
	"context"
	"database/sql"
	"log"

	"github.com/afaguilarr/go-example-webserver/app/src/dao"
	"github.com/pkg/errors"
)

type DaoEncryptionData struct {
	DB *sql.DB
}

func NewDaoEncryptionData(db *sql.DB) *DaoEncryptionData {
	return &DaoEncryptionData{
		DB: db,
	}
}

// Just as a note, Upsert queries (ON CONFLICT DO UPDATE SET) can't have
// the table as UPDATE table SET nor WHERE statements. The rows are already identified
const UpsertSalts = `
INSERT INTO encryption_data (context, crypto_secret, iv, salts)
  VALUES ($1, $2, $3, $4)
ON CONFLICT (context)
  DO UPDATE SET crypto_secret = EXCLUDED.crypto_secret, iv = EXCLUDED.iv, salts = EXCLUDED.salts, updated_at = NOW()
RETURNING id, context, crypto_secret, iv, salts
`

func (d *DaoEncryptionData) UpsertEncryptionData(ctx context.Context, s *dao.EncryptionData) error {
	var id int

	if s == nil || s.Context == nil || *s.Context == "" ||
		len(s.CryptoSecret) == 0 || len(s.IV) == 0 || len(s.Salts) == 0 {
		return errors.New("Nil or empty value detected")
	}

	// Scans should reference not-nil pointers
	err := d.DB.QueryRowContext(ctx, UpsertSalts, *s.Context, s.CryptoSecret, s.IV, s.Salts).
		Scan(
			&id,
			&s.Context,
			&s.CryptoSecret,
			&s.IV,
			&s.Salts,
		)
	if err != nil {
		return errors.Wrap(err, "there was an error while upserting the encryption data")
	}

	log.Println("New salts record was inserted")
	return nil
}

const GetEncryptionDataByContext = `
SELECT
  id,
  context,
  crypto_secret,
  iv,
  salts
FROM encryption_data
  WHERE context = $1
`

func (d *DaoEncryptionData) GetEncryptionData(ctx context.Context, s *dao.EncryptionData) error {
	var id int

	if s == nil || s.Context == nil || *s.Context == "" {
		return errors.New("Nil or empty value detected")
	}

	err := d.DB.QueryRowContext(ctx, GetEncryptionDataByContext, *s.Context).
		Scan(
			&id,
			&s.Context,
			&s.CryptoSecret,
			&s.IV,
			&s.Salts,
		)
	if err != nil {
		if err == sql.ErrNoRows {
			return err
		}
		return errors.Wrap(err, "there was an error while getting the encryption_data")
	}

	log.Println("Succeeded getting the salts record")
	return nil
}
