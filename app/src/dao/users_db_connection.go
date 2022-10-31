package dao

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

const (
	usersHostKey     = "USERS_POSTGRES_HOST"
	usersPortKey     = "USERS_POSTGRES_PORT"
	usersUserKey     = "USERS_POSTGRES_USER"
	usersPasswordKey = "USERS_POSTGRES_PASSWORD"
	usersDBNameKey   = "USERS_POSTGRES_DB"
)

func getUsersPSQLInfo() (string, error) {
	host := os.Getenv(usersHostKey)
	port := os.Getenv(usersPortKey)
	user := os.Getenv(usersUserKey)
	password := os.Getenv(usersPasswordKey)
	dbname := os.Getenv(usersDBNameKey)
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	return psqlInfo, nil
}

func CreateUsersDBConnection() *sql.DB {
	psqlInfo, err := getUsersPSQLInfo()
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
	log.Println("The connection to the Users DB was successful!")
	return db
}
