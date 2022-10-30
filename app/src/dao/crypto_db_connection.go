package dao

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

const (
	cryptoHostKey     = "CRYPTO_POSTGRES_HOST"
	cryptoPortKey     = "CRYPTO_POSTGRES_PORT"
	cryptoUserKey     = "CRYPTO_POSTGRES_USER"
	cryptoPasswordKey = "CRYPTO_POSTGRES_PASSWORD"
	cryptoDBNameKey   = "CRYPTO_POSTGRES_DB"
)

func getCryptoPSQLInfo() (string, error) {
	host := os.Getenv(cryptoHostKey)
	port := os.Getenv(cryptoPortKey)
	user := os.Getenv(cryptoUserKey)
	password := os.Getenv(cryptoPasswordKey)
	dbname := os.Getenv(cryptoDBNameKey)
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	return psqlInfo, nil
}

func CreateCryptoDBConnection() *sql.DB {
	psqlInfo, err := getCryptoPSQLInfo()
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
	log.Println("The connection to the Crypto DB was successful!")
	return db
}
