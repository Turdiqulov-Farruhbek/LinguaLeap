package main

import (
	gen "learning/genprotos"
	srvc "learning/service"
	"learning/storage/postgres"
	cfg "learning/config"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.DBConn()
	if err != nil {
		log.Fatalf("error listening: %v", err)
	}
	liss, err := net.Listen("tcp", cfg.Load().HTTPPort)
	if err != nil {
		log.Fatalf("error listening: %v", err)
	}

	server := grpc.NewServer()

	gen.RegisterQuestionServiceServer(server, srvc.NewQuestionService(db))
	gen.RegisterLanguageServiceServer(server, srvc.NewLanguageService(db))
	gen.RegisterLessonServiceServer(server, srvc.NewLessonService(db))
	gen.RegisterUserProgressServiceServer(server, srvc.NewUserProgressService(db))
	gen.RegisterUserLessonServiceServer(server, srvc.NewUserLessonService(db))
	gen.RegisterVocabularyServiceServer(server, srvc.NewVocabularyService(db))
	gen.RegisterUserServiceServer(server, srvc.NewUserService(db))

	log.Println("Server is running on port :50020")

	if err := server.Serve(liss); err != nil {
		log.Fatalf("error listening: %v", err)
	}

}
