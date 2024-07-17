package service

import (
	"context"
	"log"

	pb "learning/genprotos" 
	stg "learning/storage" 
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LessonService struct {
	repo stg.StorageInterface
	pb.UnimplementedLessonServiceServer
}

func NewLessonService(repo stg.StorageInterface) *LessonService {
	return &LessonService{
		repo: repo,
	}
}

func (s *LessonService) CreateLesson(ctx context.Context, req *pb.CreateLessonRequest) (*pb.CreateLessonResponse, error) {
	res, err := s.repo.Lessons().CreateLesson(ctx, req)
	if err != nil {
		log.Printf("Failed to create lesson: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to create lesson: %v", err)
	}
	return res, nil
}

func (s *LessonService) GetLesson(ctx context.Context, req *pb.GetLessonRequest) (*pb.GetLessonResponse, error) {
	res, err := s.repo.Lessons().GetLesson(ctx, req)
	if err != nil {
		log.Printf("Failed to get lesson: %v", err)
		return nil, status.Errorf(codes.NotFound, "Failed to get lesson: %v", err)
	}
	return res, nil
}

func (s *LessonService) ListLessons(ctx context.Context, req *pb.ListLessonsRequest) (*pb.ListLessonsResponse, error) {
	res, err := s.repo.Lessons().ListLessons(ctx, req)
	if err != nil {
		log.Printf("Failed to list lessons: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to list lessons: %v", err)
	}
	return res, nil
}

func (s *LessonService) UpdateLesson(ctx context.Context, req *pb.UpdateLessonRequest) (*pb.UpdateLessonResponse, error) {
	res, err := s.repo.Lessons().UpdateLesson(ctx, req)
	if err != nil {
		log.Printf("Failed to update lesson: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to update lesson: %v", err)
	}
	return res, nil
}

func (s *LessonService) DeleteLesson(ctx context.Context, req *pb.DeleteLessonRequest) (*pb.DeleteLessonResponse, error) {
	res, err := s.repo.Lessons().DeleteLesson(ctx, req)
	if err != nil {
		log.Printf("Failed to delete lesson: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to delete lesson: %v", err)
	}
	return res, nil
}
