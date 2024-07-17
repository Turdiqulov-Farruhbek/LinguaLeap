package service

import (
	"context"
	"log"

	pb "learning/genprotos" 
	stg "learning/storage" 
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserProgressService struct {
	pb.UnimplementedUserProgressServiceServer
	repo stg.StorageInterface
}

func NewUserProgressService(repo stg.StorageInterface) *UserProgressService {
	return &UserProgressService{
		repo: repo,
	}
}

func (s *UserProgressService) AddUserProgress(ctx context.Context, req *pb.AddUserProgressRequest) (*pb.AddUserProgressResponse, error) {
	res, err := s.repo.UserProgress().AddUserProgress(ctx, req)
	if err != nil {
		log.Printf("Failed to add user progress: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to add user progress: %v", err)
	}
	return res, nil
}

func (s *UserProgressService) GetUserProgress(ctx context.Context, req *pb.GetUserProgressRequest) (*pb.GetUserProgressResponse, error) {
	res, err := s.repo.UserProgress().GetUserProgress(ctx, req)
	if err != nil {
		log.Printf("Failed to get user progress: %v", err)
		return nil, status.Errorf(codes.NotFound, "Failed to get user progress: %v", err)
	}
	return res, nil
}
