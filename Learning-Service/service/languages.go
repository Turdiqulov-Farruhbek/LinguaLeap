package service

import (
	"context"
	"log"

	pb "learning/genprotos" 
	stg "learning/storage" 
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LanguageService struct {
	repo stg.StorageInterface
	pb.UnimplementedLanguageServiceServer
}

func NewLanguageService(repo stg.StorageInterface) *LanguageService {
	return &LanguageService{
		repo: repo,
	}
}

func (s *LanguageService) CreateLanguage(ctx context.Context, req *pb.CreateLanguageRequest) (*pb.CreateLanguageResponse, error) {
	res, err := s.repo.Languages().CreateLanguage(ctx, req)
	if err != nil {
		log.Printf("Failed to create language: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to create language: %v", err)
	}
	return res, nil
}

func (s *LanguageService) GetLanguage(ctx context.Context, req *pb.GetLanguageRequest) (*pb.GetLanguageResponse, error) {
	res, err := s.repo.Languages().GetLanguage(ctx, req)
	if err != nil {
		log.Printf("Failed to get language: %v", err)
		return nil, status.Errorf(codes.NotFound, "Failed to get language: %v", err)
	}
	return res, nil
}

func (s *LanguageService) ListLanguages(ctx context.Context, req *pb.ListLanguagesRequest) (*pb.ListLanguagesResponse, error) {
	res, err := s.repo.Languages().ListLanguages(ctx, req)
	if err != nil {
		log.Printf("Failed to list languages: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to list languages: %v", err)
	}
	return res, nil
}

func (s *LanguageService) UpdateLanguage(ctx context.Context, req *pb.UpdateLanguageRequest) (*pb.UpdateLanguageResponse, error) {
	res, err := s.repo.Languages().UpdateLanguage(ctx, req)
	if err != nil {
		log.Printf("Failed to update language: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to update language: %v", err)
	}
	return res, nil
}

func (s *LanguageService) DeleteLanguage(ctx context.Context, req *pb.DeleteLanguageRequest) (*pb.DeleteLanguageResponse, error) {
	res, err := s.repo.Languages().DeleteLanguage(ctx, req)
	if err != nil {
		log.Printf("Failed to delete language: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to delete language: %v", err)
	}
	return res, nil
}
