package service

import (
	"context"
	"log"

	pb "learning/genprotos"
	stg "learning/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserLessonService struct {
	repo stg.StorageInterface
	pb.UnimplementedUserLessonServiceServer
}

func NewUserLessonService(repo stg.StorageInterface) *UserLessonService {
	return &UserLessonService{
		repo: repo,
	}
}

func (s *UserLessonService) AddUserLesson(ctx context.Context, req *pb.AddUserLessonRequest) (*pb.AddUserLessonResponse, error) {
	res, err := s.repo.UserLessons().AddUserLesson(ctx, req)
	if err != nil {
		log.Printf("Failed to add user lesson: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to add user lesson: %v", err)
	}
	return res, nil
}

func (s *UserLessonService) GetUserLessons(ctx context.Context, req *pb.GetUserLessonsRequest) (*pb.GetUserLessonsResponse, error) {
	res, err := s.repo.UserLessons().GetUserLessons(ctx, req)
	if err != nil {
		log.Printf("Failed to get user lessons: %v", err)
		return nil, status.Errorf(codes.NotFound, "Failed to get user lessons: %v", err)
	}
	return res, nil
}
