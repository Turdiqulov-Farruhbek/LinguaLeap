package handlers

import (
	au "gateway/genprotos/auth"
	gm "gateway/genprotos/game"
	ln "gateway/genprotos/learn"
	"gateway/genprotos/progres"
	// us "gateway/genprotos/user"

	"google.golang.org/grpc"
)

type AuthHangler struct {
	Auth au.AuthServiceClient
}

func NewAuthHangler(authCon *grpc.ClientConn) *AuthHangler {
	return &AuthHangler{
		Auth: au.NewAuthServiceClient(authCon),
	}
}

// type UserHandler struct {
// 	User us.UserServiceClient
// }

// func NewUserHandler(userCon *grpc.ClientConn) *UserHandler {
// 	return &UserHandler{
// 		User: us.NewUserServiceClient(userCon),
// 	}
// }

type GameHandler struct {
	Game gm.GameServiceClient
}

func NewGameHandler(gameCon *grpc.ClientConn) *GameHandler {
	return &GameHandler{
		Game: gm.NewGameServiceClient(gameCon),
	}
}

type ProgressHandler struct {
	Progress progres.ProgressServiceClient
}

func NewProgressHandler(progressCon *grpc.ClientConn) *ProgressHandler {
	return &ProgressHandler{
		Progress: progres.NewProgressServiceClient(progressCon),
	}
}

type LearnHandler struct {
	Exercises         ln.QuestionServiceClient
	Language          ln.LanguageServiceClient
	LessonS           ln.LessonServiceClient
	UserProgress      ln.UserProgressServiceClient
	UserLessonService ln.UserLessonServiceClient
	User              ln.UserServiceClient
	Vocabulary        ln.VocabularyServiceClient
}

func NewLearnHandler(learnCon *grpc.ClientConn) *LearnHandler {
	return &LearnHandler{
		Exercises: 		   ln.NewQuestionServiceClient(learnCon),
		Language:          ln.NewLanguageServiceClient(learnCon),
        LessonS:           ln.NewLessonServiceClient(learnCon),
        UserProgress:      ln.NewUserProgressServiceClient(learnCon),
        UserLessonService: ln.NewUserLessonServiceClient(learnCon),
        User:              ln.NewUserServiceClient(learnCon),
	}
}

