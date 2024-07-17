package postgres

import (
	"context"
	"database/sql"
	"fmt"

	pb "learning/genprotos" // Protobuf faylini import qilish

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/lib/pq"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type userRepo struct {
	db *pgx.Conn
}

func NewUser(db *pgx.Conn) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	id := uuid.New().String()

	query := `
		INSERT INTO users (id, username, email, password_hash, full_name, native_language, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		RETURNING id
	`

	_, err := r.db.Exec(ctx, query, id, req.Username, req.Email, req.Password, req.FullName, req.NativeLanguage)
	if err != nil {
		return nil, fmt.Errorf("could not execute query: %v", err)
	}

	return &pb.CreateUserResponse{Id: id}, nil
}

func (r *userRepo) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	query := `
		SELECT id, username, email, password_hash, full_name, native_language, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	row := r.db.QueryRow(ctx, query, req.UserId)

	var user pb.User
	var createdAt, updatedAt pq.NullTime

	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.PasswordHash, &user.FullName, &user.NativeLanguage, &createdAt, &updatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("could not execute query: %v", err)
	}

	if createdAt.Valid {
		user.CreatedAt = timestamppb.New(createdAt.Time)
	}
	if updatedAt.Valid {
		user.UpdatedAt = timestamppb.New(updatedAt.Time)
	}

	return &pb.GetUserResponse{User: &user}, nil
}

func (r *userRepo) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	query := `
		UPDATE users
		SET username = $1, email = $2, password_hash = $3, full_name = $4, native_language = $5, updated_at = CURRENT_TIMESTAMP
		WHERE id = $6
		RETURNING id, username, email, password_hash, full_name, native_language, created_at, updated_at
	`

	row := r.db.QueryRow(ctx, query, req.Username, req.Email, req.Password, req.FullName, req.NativeLanguage, req.Id)

	var user pb.User
	var createdAt, updatedAt pq.NullTime

	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.PasswordHash, &user.FullName, &user.NativeLanguage, &createdAt, &updatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("could not execute query: %v", err)
	}

	if createdAt.Valid {
		user.CreatedAt = timestamppb.New(createdAt.Time)
	}
	if updatedAt.Valid {
		user.UpdatedAt = timestamppb.New(updatedAt.Time)
	}

	return &pb.UpdateUserResponse{User: &user}, nil
}

func (r *userRepo) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	query := `
		DELETE FROM users
		WHERE id = $1
		RETURNING id
	`

	var id string
	err := r.db.QueryRow(ctx, query, req.Id).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("could not execute query: %v", err)
	}

	return &pb.DeleteUserResponse{Id: id}, nil
}

func (r *userRepo) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	query := `
		SELECT id, username, email, password_hash, full_name, native_language, created_at, updated_at
		FROM users
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("could not execute query: %v", err)
	}
	defer rows.Close()

	var users []*pb.User

	for rows.Next() {
		var user pb.User
		var createdAt, updatedAt pq.NullTime

		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.PasswordHash, &user.FullName, &user.NativeLanguage, &createdAt, &updatedAt)
		if err != nil {
			return nil, fmt.Errorf("could not scan row: %v", err)
		}

		if createdAt.Valid {
			user.CreatedAt = timestamppb.New(createdAt.Time)
		}
		if updatedAt.Valid {
			user.UpdatedAt = timestamppb.New(updatedAt.Time)
		}

		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}

	return &pb.ListUsersResponse{Users: users}, nil
}
