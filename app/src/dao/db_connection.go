package dao

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

const (
	hostKey     = "POSTGRES_HOST"
	portKey     = "POSTGRES_PORT"
	userKey     = "POSTGRES_USER"
	passwordKey = "POSTGRES_PASSWORD"
	dbnameKey   = "POSTGRES_DB"
)

type DBConnectionHandler struct {
	KeysPrefix, Host, Port, User, Password, DBName, DataSourceName string
}

func NewDBConnectionHandler(keysPrefix string) *DBConnectionHandler {
	return &DBConnectionHandler{
		KeysPrefix: keysPrefix,
	}
}

func (dbch *DBConnectionHandler) populatePSQLInfo() {
	dbch.Host = os.Getenv(dbch.KeysPrefix + hostKey)
	dbch.Port = os.Getenv(dbch.KeysPrefix + portKey)
	dbch.User = os.Getenv(dbch.KeysPrefix + userKey)
	dbch.Password = os.Getenv(dbch.KeysPrefix + passwordKey)
	dbch.DBName = os.Getenv(dbch.KeysPrefix + dbnameKey)
}

func (dbch *DBConnectionHandler) populateDataSourceName() {
	dbch.DataSourceName = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbch.Host, dbch.Port, dbch.User, dbch.Password, dbch.DBName)
}

func (dbch *DBConnectionHandler) CreateDBConnection() (*sql.DB, error) {
	dbch.populatePSQLInfo()
	dbch.populateDataSourceName()

	db, err := sql.Open("postgres", dbch.DataSourceName)
	if err != nil {
		return nil, errors.Wrap(err, "while opening database connection")
	}

	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "while pinging database database")
	}

	log.Printf("The connection to the %s DB was successful!\n", dbch.DBName)
	return db, nil
}
