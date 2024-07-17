package storage

import (
	"context"
	gen "learning/genprotos"

)

type StorageInterface interface {
	Languages() LanguageInterface
	Lessons() LessonsInterface
	UserProgress() UserProgresInterface
	Vocabularies() VocabularyInterface
	Exercises() ExercisesInterface
	Users() UserInterface
	UserLessons() UserLessonsInterface
}

type LanguageInterface interface {
	GetLanguage(ctx context.Context, req *gen.GetLanguageRequest) (*gen.GetLanguageResponse, error)
    ListLanguages(ctx context.Context, req *gen.ListLanguagesRequest) (*gen.ListLanguagesResponse, error)
	CreateLanguage(ctx context.Context, req *gen.CreateLanguageRequest) (*gen.CreateLanguageResponse, error)
	UpdateLanguage(ctx context.Context, req *gen.UpdateLanguageRequest) (*gen.UpdateLanguageResponse, error)
	DeleteLanguage(ctx context.Context, req *gen.DeleteLanguageRequest) (*gen.DeleteLanguageResponse, error)
}


type LessonsInterface interface {
	GetLesson(ctx context.Context, req *gen.GetLessonRequest) (*gen.GetLessonResponse, error)
	ListLessons(ctx context.Context, req *gen.ListLessonsRequest) (*gen.ListLessonsResponse, error)
	CreateLesson(ctx context.Context, req *gen.CreateLessonRequest) (*gen.CreateLessonResponse, error)
	UpdateLesson(ctx context.Context, req *gen.UpdateLessonRequest) (*gen.UpdateLessonResponse, error)
	DeleteLesson(ctx context.Context, req *gen.DeleteLessonRequest) (*gen.DeleteLessonResponse, error)
}


type UserProgresInterface interface {
	GetUserProgress(ctx context.Context, req *gen.GetUserProgressRequest) (*gen.GetUserProgressResponse, error)
	AddUserProgress(ctx context.Context, req *gen.AddUserProgressRequest) (*gen.AddUserProgressResponse, error)
}


type VocabularyInterface interface {
	AddVocabulary(ctx context.Context, req *gen.AddVocabularyRequest) (*gen.AddVocabularyResponse, error)
    GetVocabulary(ctx context.Context, req *gen.GetVocabularyRequest) (*gen.GetVocabularyResponse, error)
	UpdateVocabulary(ctx context.Context, req *gen.UpdateVocabularyRequest) (*gen.UpdateVocabularyResponse, error)
	DeleteVocabulary(ctx context.Context, req *gen.DeleteVocabularyRequest) (*gen.DeleteVocabularyResponse, error)
}


type ExercisesInterface interface {
    CreateQuestion(ctx context.Context, req *gen.CreateQuestionRequest) (*gen.CreateQuestionResponse, error)
	GetQuestion(ctx context.Context, req *gen.GetQuestionRequest) (*gen.GetQuestionResponse, error)
	ListQuestions(ctx context.Context, req *gen.ListQuestionsRequest) (*gen.ListQuestionsResponse, error)
	UpdateQuestion(ctx context.Context, req *gen.UpdateQuestionRequest) (*gen.UpdateQuestionResponse, error)
	DeleteQuestion(ctx context.Context, req *gen.DeleteQuestionRequest) (*gen.DeleteQuestionResponse, error)
}


type UserInterface interface {
	GetUser(ctx context.Context, req *gen.GetUserRequest) (*gen.GetUserResponse, error)
	CreateUser(ctx context.Context, req *gen.CreateUserRequest) (*gen.CreateUserResponse, error)
	UpdateUser(ctx context.Context, req *gen.UpdateUserRequest) (*gen.UpdateUserResponse, error)
	DeleteUser(ctx context.Context, req *gen.DeleteUserRequest) (*gen.DeleteUserResponse, error)
	ListUsers(ctx context.Context, req *gen.ListUsersRequest) (*gen.ListUsersResponse, error)
}


type UserLessonsInterface interface {
    GetUserLessons(ctx context.Context, req *gen.GetUserLessonsRequest) (*gen.GetUserLessonsResponse, error)
    AddUserLesson(ctx context.Context, req *gen.AddUserLessonRequest) (*gen.AddUserLessonResponse, error)
}