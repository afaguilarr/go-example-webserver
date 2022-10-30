package dao

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

const (
	hostKey     = "HELLO_WORLD_POSTGRES_HOST"
	portKey     = "HELLO_WORLD_POSTGRES_PORT"
	userKey     = "HELLO_WORLD_POSTGRES_USER"
	passwordKey = "HELLO_WORLD_POSTGRES_PASSWORD"
	dbnameKey   = "HELLO_WORLD_POSTGRES_DB"
)

func getPSQLInfo() (string, error) {
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
	log.Println("The connection to the Hello World DB was successful!")
	return db
}
