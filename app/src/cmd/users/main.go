package main

import (
	"log"

	"github.com/afaguilarr/go-example-webserver/app/src/dao"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"

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
		panic(errors.Wrap(err, "while creating DB connection"))
	}
	defer usersDB.Close()

	// sd := postgres.NewDaoEncryptionData(usersDB)
	// sb := business.NewBusinessCrypto(sd)
	// sc := services.NewServicesCrypto(sb)

	// s := grpc.NewServer()
	// // Register reflection service on gRPC server.
	// // https://github.com/grpc/grpc-go/blob/master/Documentation/server-reflection-tutorial.md
	// reflection.Register(s)
	// proto.RegisterCryptoServer(s, sc)

	// lis, err := net.Listen("tcp", ":8080")
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }

	// log.Printf("Crypto server listening at %v!", lis.Addr())
	// if err := s.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }
}
