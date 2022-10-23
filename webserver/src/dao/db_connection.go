package dao

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

const (
	hostKey     = "POSTGRES_HOST"
	portKey     = "POSTGRES_PORT"
	userKey     = "POSTGRES_USER"
	passwordKey = "POSTGRES_PASSWORD"
	dbnameKey   = "POSTGRES_DB"
)

func getPSQLInfo() (string, error) {
	err := godotenv.Load("./.env")
	if err != nil {
		return "", errors.Wrap(err, "there was an error loading the env variables")
	}
	host := os.Getenv(hostKey)
	port := os.Getenv(portKey)
	user := os.Getenv(userKey)
	password := os.Getenv(passwordKey)
	dbname := os.Getenv(dbnameKey)
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	return psqlInfo, nil
}

func CreateDBConnection() *sql.DB {
	psqlInfo, err := getPSQLInfo()
	if err != nil {
		panic(err)
	}
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("The connection to the DB was successful!")
	return db
}
