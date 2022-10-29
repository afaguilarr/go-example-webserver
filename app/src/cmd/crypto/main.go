package main

import (
	"context"
	"log"
	"net"

	"github.com/afaguilarr/go-example-webserver/app/src/dao"
	"github.com/afaguilarr/go-example-webserver/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/lib/pq"
)

// server is used to implement CryptoServer
type cryptoServer struct {
	proto.UnimplementedCryptoServer
}

// Encrypt encrypts a password by using a context string
func (cs *cryptoServer) Encrypt(ctx context.Context, in *proto.EncryptRequest) (*proto.EncryptResponse, error) {
	log.Println("Encrypt RPC was called!")
	return &proto.EncryptResponse{EncryptedValue: "jiji"}, nil
}

// Decrypt encrypts a password by using a context string
func (cs *cryptoServer) Decrypt(ctx context.Context, in *proto.DecryptRequest) (*proto.DecryptResponse, error) {
	return &proto.DecryptResponse{DecryptedValue: "jiji"}, nil
}

func main() {
	cryptoDB := dao.CreateCryptoDBConnection()
	defer cryptoDB.Close()

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	// Register reflection service on gRPC server.
	// https://github.com/grpc/grpc-go/blob/master/Documentation/server-reflection-tutorial.md
	reflection.Register(s)
	proto.RegisterCryptoServer(s, &cryptoServer{})

	log.Printf("Crypto server listening at %v!", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
