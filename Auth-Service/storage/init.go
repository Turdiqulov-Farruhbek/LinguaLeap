package postgres

import (
	pbAuth "auth/genprotos/auth"
	pbUser "auth/genprotos/user"
)

type InitRoot interface {
	Auth() Auth
	User() User
}

type Auth interface {
	Register(req *pbAuth.RegisterRequest) (*pbAuth.RegisterResponse, error)
	Login(req *pbAuth.LoginRequest) (*pbAuth.LoginResponse, error)
	Logout(req *pbAuth.LogoutRequest) (*pbAuth.LogoutResponse, error)
	ResetPassword(req *pbAuth.ResetPasswordRequest) (*pbAuth.ResetPasswordResponse, error)
}

type User interface {
	GetProfile(req *pbUser.GetProfileRequest) (*pbUser.GetProfileResponse, error)
	UpdateProfile(req *pbUser.UpdateProfileRequest) (*pbUser.UpdateProfileResponse, error)
	ChangePassword(req *pbUser.ChangePasswordRequest) (*pbUser.ChangePasswordResponse, error)
}
