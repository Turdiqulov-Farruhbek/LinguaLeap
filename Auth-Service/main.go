package main

import (
	"fmt"
	"log"

)

func main() {
	db, err := postgres.ConnectDb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	router := api.NewRouter(db)
	
	fmt.Println("Server is running on port 2121")
	if err := router.Run(":2121"); err != nil {
		log.Fatal(err)
	}
	
}
