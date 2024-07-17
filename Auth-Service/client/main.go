package main

import (
	"log"
	"net"

	"auth/genprotos/auth"
	pbu "auth/genprotos/user"
	"auth/service"
	"auth/storage/postgres"

	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}
	liss, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatal("Error while connection on tcp: ", err.Error())
	}

	s := grpc.NewServer()
	auth.RegisterAuthServiceServer(s, service.NewAuthService(db))
	pbu.RegisterUserServiceServer(s, service.NewUserService(db))
	
	log.Printf("server listening at %v", liss.Addr())
	if err := s.Serve(liss); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
