package service

import (
	"context"
	"log"

	pb "learning/genprotos" 
	stg "learning/storage" 
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type QuestionService struct {
	pb.UnimplementedQuestionServiceServer
	repo stg.StorageInterface
}

func NewQuestionService(repo stg.StorageInterface) *QuestionService {
	return &QuestionService{
		repo: repo,
	}
}

func (s *QuestionService) CreateQuestion(ctx context.Context, req *pb.CreateQuestionRequest) (*pb.CreateQuestionResponse, error) {
	res, err := s.repo.Exercises().CreateQuestion(ctx, req)
	if err != nil {
		log.Printf("Failed to create question: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to create question: %v", err)
	}
	return res, nil
}

func (s *QuestionService) GetQuestion(ctx context.Context, req *pb.GetQuestionRequest) (*pb.GetQuestionResponse, error) {
	res, err := s.repo.Exercises().GetQuestion(ctx, req)
	if err != nil {
		log.Printf("Failed to get question: %v", err)
		return nil, status.Errorf(codes.NotFound, "Failed to get question: %v", err)
	}
	return res, nil
}

func (s *QuestionService) ListQuestions(ctx context.Context, req *pb.ListQuestionsRequest) (*pb.ListQuestionsResponse, error) {
	res, err := s.repo.Exercises().ListQuestions(ctx, req)
	if err != nil {
		log.Printf("Failed to list questions: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to list questions: %v", err)
	}
	return res, nil
}

func (s *QuestionService) UpdateQuestion(ctx context.Context, req *pb.UpdateQuestionRequest) (*pb.UpdateQuestionResponse, error) {
	res, err := s.repo.Exercises().UpdateQuestion(ctx, req)
	if err != nil {
		log.Printf("Failed to update question: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to update question: %v", err)
	}
	return res, nil
}

func (s *QuestionService) DeleteQuestion(ctx context.Context, req *pb.DeleteQuestionRequest) (*pb.DeleteQuestionResponse, error) {
	res, err := s.repo.Exercises().DeleteQuestion(ctx, req)
	if err != nil {
		log.Printf("Failed to delete question: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to delete question: %v", err)
	}
	return res, nil
}
