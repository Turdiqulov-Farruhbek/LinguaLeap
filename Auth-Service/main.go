package main

import (
	"fmt"
	"log"

	"auth/api"
	reds "auth/api/handler"
	"auth/genprotos/auth"
	"auth/genprotos/user"

	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ClientCon, err := grpc.NewClient(fmt.Sprintf("localhost%s", ":8088"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while NEwclient: ", err.Error())
	}
	defer ClientCon.Close()

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	redis := reds.NewInMemoryStorage(rdb)
	auth := auth.NewAuthServiceClient(ClientCon)
	user := user.NewUserServiceClient(ClientCon)

	hander := reds.NewHandler(auth, user, redis)
	routir := api.NewGin(hander)

	fmt.Println("Server started on port:8080")
	err = routir.Run(":8080")
	if err != nil {
		log.Fatal("Error while running server: ", err.Error())
	}
}
