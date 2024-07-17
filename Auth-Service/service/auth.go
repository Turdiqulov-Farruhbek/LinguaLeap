package service

import (
	"context"
	"log"

	pb "auth/genprotos/auth"
	s "auth/storage"
)

type AuthService struct {
	stg s.InitRoot
	pb.UnimplementedAuthServiceServer
}

func NewAuthService(stg s.InitRoot) *AuthService {
	return &AuthService{stg: stg}
}

func (a *AuthService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	res, err := a.stg.Auth().Register(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return res, nil
}

func (a *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	res, err := a.stg.Auth().Login(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return res, nil
}

func (a *AuthService) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	res, err := a.stg.Auth().Logout(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return res, nil
}

func (a *AuthService) ResetPassword(ctx context.Context, req *pb.ResetPasswordRequest) (*pb.ResetPasswordResponse, error) {
	res, err := a.stg.Auth().ResetPassword(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return res, nil
}
