package postgres

import (
	"database/sql"

	pb "auth/genprotos/user"
)

type UserStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{db: db}
}

func (p *UserStorage) GetProfile(req *pb.GetProfileRequest) (*pb.GetProfileResponse, error) {
	var profile pb.GetProfileResponse
	query := `
		SELECT id, username, email, full_name, level, xp, created_at
		FROM users
		WHERE id = $1
	`
	err := p.db.QueryRow(query, req.Token).Scan(
		&profile.Id, &profile.Username, &profile.Email,
		&profile.FullName, &profile.Level, &profile.Xp, &profile.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (p *UserStorage) UpdateProfile(req *pb.UpdateProfileRequest) (*pb.UpdateProfileResponse, error) {
	query := `
		UPDATE users
		SET full_name = $1, avatar_url = $2, updated_at = now()
		WHERE id = $3
		RETURNING id, username, email, full_name, avatar_url, updated_at
	`
	var profile pb.UpdateProfileResponse
	err := p.db.QueryRow(query, req.FullName, req.AvatarUrl, req.Token).Scan(
		&profile.Id, &profile.Username, &profile.Email,
		&profile.FullName, &profile.AvatarUrl, &profile.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (p *UserStorage) ChangePassword(req *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {
	query := `
		UPDATE users
		SET password_hash = $1, updated_at = now()
		WHERE token = $2
	`
	_, err := p.db.Exec(query, req.NewPassword, req.Token)
	if err != nil {
		return nil, err
	}
	return &pb.ChangePasswordResponse{Message: "Password successfully changed"}, nil
}
