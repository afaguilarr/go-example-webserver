package main

import (
	"log"
	"net"

	"github.com/afaguilarr/go-example-webserver/app/src/business"
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
	cryptoDBPrefix = "CRYPTO_"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("there was an error loading the env variables: %s", err.Error())
	}

	dbch := dao.NewDBConnectionHandler(cryptoDBPrefix)
	cryptoDB, err := dbch.CreateDBConnection()
	if err != nil {
		panic(errors.Wrap(err, "while creating DB connection"))
	}
	defer cryptoDB.Close()

	sd := postgres.NewDaoEncryptionData(cryptoDB)
	sb := business.NewBusinessCrypto(sd)
	sc := services.NewServicesCrypto(sb)

	s := grpc.NewServer()
	// Register reflection service on gRPC server.
	// https://github.com/grpc/grpc-go/blob/master/Documentation/server-reflection-tutorial.md
	reflection.Register(s)
	proto.RegisterCryptoServer(s, sc)

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Crypto server listening at %v!", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
