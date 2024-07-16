package handlers

import (
	"getway/genproto/authors"
	"getway/genproto/books"
	"getway/genproto/borrowers"
	"getway/genproto/genres"
	"getway/genproto/users"

	"google.golang.org/grpc"
)

type HandlerStruct struct {
	
}


func NewHandler(libraryConn *grpc.ClientConn ) *HandlerStruct {
	return &HandlerStruct{

    }
}