package service

import (
	"context"
	"log"

	pb "learning/genprotos"
	stg "learning/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type VocabularyService struct {
	pb.UnimplementedVocabularyServiceServer
	repo stg.StorageInterface
}

func NewVocabularyService(repo stg.StorageInterface) *VocabularyService {
	return &VocabularyService{
		repo: repo,
	}
}

func (s *VocabularyService) AddVocabulary(ctx context.Context, req *pb.AddVocabularyRequest) (*pb.AddVocabularyResponse, error) {
	res, err := s.repo.Vocabularies().AddVocabulary(ctx, req)
	if err != nil {
		log.Printf("Failed to add vocabulary: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to add vocabulary: %v", err)
	}
	return res, nil
}

func (s *VocabularyService) GetVocabulary(ctx context.Context, req *pb.GetVocabularyRequest) (*pb.GetVocabularyResponse, error) {
	res, err := s.repo.Vocabularies().GetVocabulary(ctx, req)
	if err != nil {
		log.Printf("Failed to get vocabulary: %v", err)
		return nil, status.Errorf(codes.NotFound, "Failed to get vocabulary: %v", err)
	}
	return res, nil
}

func (s *VocabularyService) UpdateVocabulary(ctx context.Context, req *pb.UpdateVocabularyRequest) (*pb.UpdateVocabularyResponse, error) {
	res, err := s.repo.Vocabularies().UpdateVocabulary(ctx, req)
	if err != nil {
		log.Printf("Failed to update vocabulary: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to update vocabulary: %v", err)
	}
	return res, nil
}

func (s *VocabularyService) DeleteVocabulary(ctx context.Context, req *pb.DeleteVocabularyRequest) (*pb.DeleteVocabularyResponse, error) {
	res, err := s.repo.Vocabularies().DeleteVocabulary(ctx, req)
	if err != nil {
		log.Printf("Failed to delete vocabulary: %v", err)
		return nil, status.Errorf(codes.NotFound, "Failed to delete vocabulary: %v", err)
	}
	return res, nil
}
