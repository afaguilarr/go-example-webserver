package main

import (
	"log"
	"net"

	"github.com/afaguilarr/go-example-webserver/app/src/business"
	"github.com/afaguilarr/go-example-webserver/app/src/cmd"
	"github.com/afaguilarr/go-example-webserver/app/src/crypto_client"
	"github.com/afaguilarr/go-example-webserver/app/src/dao"
	"github.com/afaguilarr/go-example-webserver/app/src/dao/postgres"
	"github.com/afaguilarr/go-example-webserver/app/src/services"
	"github.com/afaguilarr/go-example-webserver/proto"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/lib/pq"
)

const (
	usersDBPrefix = "USERS_"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("there was an error loading the env variables: %s", err.Error())
	}

	dbch := dao.NewDBConnectionHandler(usersDBPrefix)
	usersDB, err := dbch.CreateDBConnection()
	if err != nil {
		panic(errors.Wrap(err, "while creating Users DB connection"))
	}
	defer func() {
		err := usersDB.Close()
		if err != nil {
			panic(errors.Wrap(err, "while closing Users DB Connection"))
		}
	}()

	cc := crypto_client.NewCryptoClientHandler(cmd.CryptoHost, cmd.DefaultPort)
	err = cc.CreateConnection()
	if err != nil {
		panic(errors.Wrap(err, "while creating Crypto server Connection"))
	}
	defer func() {
		err := cc.CloseConnection()
		if err != nil {
			panic(errors.Wrap(err, "while closing Crypto server Connection"))
		}
	}()

	sd := postgres.NewDaoUsers(usersDB)
	sb := business.NewBusinessUsers(sd, cc)
	su := services.NewServicesUsers(sb)

	s := grpc.NewServer()
	reflection.Register(s)
	proto.RegisterUsersServer(s, su)

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Users server listening at %v!", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
