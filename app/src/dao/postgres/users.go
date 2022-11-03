package postgres

import (
	"context"
	"database/sql"
	"log"
	"strings"

	"github.com/afaguilarr/go-example-webserver/app/src/dao"
	"github.com/pkg/errors"
)

type DaoUsers struct {
	DB *sql.DB
}

func NewDaoUsers(db *sql.DB) *DaoUsers {
	return &DaoUsers{
		DB: db,
	}
}

const InsertUserQuery = `
INSERT INTO users (username, description, encrypted_password) VALUES ($1, $2, $3)
RETURNING id, username, description, encrypted_password
`

const uniqueUsernameError = "duplicate key value violates unique constraint \"users_username_key\""

var UniqueUsernameErr = errors.New("username already exists")

// InsertUser executes a transaction that inserts the data
// for the User, the PetMaster, and the Location related to the user
func (d *DaoUsers) InsertUser(ctx context.Context, u *dao.User) error {
	var userID, petMasterID, locationID int

	if u == nil || u.PetMasterInfo == nil || u.PetMasterInfo.Location == nil || u.EncryptedPassword == nil {
		return errors.New("nil value detected")
	}

	petMasterInfo := u.PetMasterInfo
	location := u.PetMasterInfo.Location

	// Get a Tx (transaction) for making transaction requests.
	tx, err := d.DB.BeginTx(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "while beginning transaction")
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	err = tx.QueryRowContext(ctx, InsertUserQuery, u.Username, u.Description, u.EncryptedPassword).
		Scan(
			&userID,
			&u.Username,
			&u.Description,
			&u.EncryptedPassword,
		)
	if err != nil {
		if strings.Contains(err.Error(), uniqueUsernameError) {
			return UniqueUsernameErr
		}
		return errors.Wrap(err, "while inserting the user")
	}

	err = tx.QueryRowContext(ctx, InsertPetMasterQuery, petMasterInfo.Name, petMasterInfo.ContactNumber, userID).
		Scan(
			&petMasterID,
			&petMasterInfo.Name,
			&petMasterInfo.ContactNumber,
			&userID,
		)
	if err != nil {
		return errors.Wrap(err, "while inserting the pet master info")
	}

	err = tx.QueryRowContext(ctx, InsertLocationQuery, location.Country, location.StateOrProvince, location.CityOrMunicipality, location.Neighborhood, location.ZipCode, petMasterID).
		Scan(
			&locationID,
			&location.Country,
			&location.StateOrProvince,
			&location.CityOrMunicipality,
			&location.Neighborhood,
			&location.ZipCode,
			&petMasterID,
		)
	if err != nil {
		return errors.Wrap(err, "while inserting the location info")
	}

	// Commit the transaction, this will commit the changes,
	// even if the Rollback action is executed on the function defer
	if err = tx.Commit(); err != nil {
		return errors.Wrap(err, "while committing the transaction")
	}

	log.Println("New user, pet master, and location records were inserted")
	return nil
}

const GetPasswordByUsernameQuery = `SELECT encrypted_password FROM users WHERE username = $1`

// GetPasswordByUsername returns the encrypted password of a user, querying it by its username
func (d *DaoUsers) GetPasswordByUsername(ctx context.Context, u string) (string, error) {
	var encryptedPassword string

	if u == "" {
		return "", errors.New("username can't be empty")
	}

	err := d.DB.QueryRowContext(ctx, GetPasswordByUsernameQuery, u).Scan(
		&encryptedPassword,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", err
		}
		return "", errors.Wrap(err, "while getting the password by username")
	}

	log.Println("Found an encrypted password password with the provided username!")
	return encryptedPassword, nil
}
