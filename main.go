package main

import (
	login2 "github.com/starbuy-commerce/auth-server/login"
	login "github.com/starbuy-commerce/auth-server/protobuf/protobuf_login"
	token "github.com/starbuy-commerce/auth-server/protobuf/protobuf_token"
	token2 "github.com/starbuy-commerce/auth-server/token"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Couldnt listen port %v: %v", port, err.Error())
		return
	}

	grpcServer := grpc.NewServer()
	login.RegisterLoginServiceServer(grpcServer, &login2.Server{})
	token.RegisterTokenValidationServiceServer(grpcServer, &token2.Server{})

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Couldnt serve port %v: %v", port, err.Error())
		return
	}

	defer listener.Close()
}
