package service

import (
	"context"
	"log"

	pb "auth/genprotos/user"
	s "auth/storage"
)

type UserService struct {
	stg s.InitRoot
	pb.UnimplementedUserServiceServer
}

func NewUserService(stg s.InitRoot) *UserService {
	return &UserService{stg: stg}
}

func (u *UserService) GetProfile(ctx context.Context, req *pb.GetProfileRequest) (*pb.GetProfileResponse, error) {
	profile, err := u.stg.User().GetProfile(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return profile, nil
}

func (u *UserService) UpdateProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.UpdateProfileResponse, error) {
	res, err := u.stg.User().UpdateProfile(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return res, nil
}

func (u *UserService) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {
	res, err := u.stg.User().ChangePassword(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return res, nil
}
