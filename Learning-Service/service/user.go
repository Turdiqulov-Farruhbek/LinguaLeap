package service

import (
	"context"
	"log"

	pb "learning/genprotos"
	stg "learning/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	repo stg.StorageInterface
}

func NewUserService(repo stg.StorageInterface) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	res, err := s.repo.Users().CreateUser(ctx, req)
	if err != nil {
		log.Printf("Failed to create user: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to create user: %v", err)
	}
	return res, nil
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	res, err := s.repo.Users().GetUser(ctx, req)
	if err != nil {
		log.Printf("Failed to get user: %v", err)
		return nil, status.Errorf(codes.NotFound, "Failed to get user: %v", err)
	}
	return res, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	res, err := s.repo.Users().UpdateUser(ctx, req)
	if err != nil {
		log.Printf("Failed to update user: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to update user: %v", err)
	}
	return res, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	res, err := s.repo.Users().DeleteUser(ctx, req)
	if err != nil {
		log.Printf("Failed to delete user: %v", err)
		return nil, status.Errorf(codes.NotFound, "Failed to delete user: %v", err)
	}
	return res, nil
}

func (s *UserService) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	res, err := s.repo.Users().ListUsers(ctx, req)
	if err != nil {
		log.Printf("Failed to list users: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to list users: %v", err)
	}
	return res, nil
}
