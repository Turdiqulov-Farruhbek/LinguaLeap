package postgres

import (
	"context"
	"time"

	pb "learning/genprotos"

	"github.com/jackc/pgx/v5"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserLessonRepo struct {
	db *pgx.Conn
}

func NewUserLessons(db *pgx.Conn) *UserLessonRepo {
	return &UserLessonRepo{
		db: db,
	}
}

// AddUserLesson - Foydalanuvchi darslarini qo'shish
func (r *UserLessonRepo) AddUserLesson(ctx context.Context, req *pb.AddUserLessonRequest) (*pb.AddUserLessonResponse, error) {
	query := `
        INSERT INTO user_lessons (user_id, lesson_id, completed_at)
        VALUES ($1, $2, $3)
        RETURNING id
    `

	var id int32
	err := r.db.QueryRow(ctx, query, req.UserId, req.LessonId, req.CompletedAt.AsTime()).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &pb.AddUserLessonResponse{
		Id: id,
	}, nil
}

// GetUserLessons - Foydalanuvchi darslarini olish
func (r *UserLessonRepo) GetUserLessons(ctx context.Context, req *pb.GetUserLessonsRequest) (*pb.GetUserLessonsResponse, error) {
	query := `
        SELECT id, user_id, lesson_id, completed_at, created_at, updated_at
        FROM user_lessons
        WHERE user_id = $1
    `

	rows, err := r.db.Query(ctx, query, req.UserId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userLessons []*pb.UserLesson
	for rows.Next() {
		var userLesson pb.UserLesson
		var completedAt, createdAt, updatedAt time.Time

		err := rows.Scan(
			&userLesson.Id,
			&userLesson.UserId,
			&userLesson.LessonId,
			&completedAt,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, err
		}

		userLesson.CompletedAt = timestamppb.New(completedAt)
		userLesson.CreatedAt = timestamppb.New(createdAt)
		userLesson.UpdatedAt = timestamppb.New(updatedAt)

		userLessons = append(userLessons, &userLesson)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &pb.GetUserLessonsResponse{
		UserLessons: userLessons,
	}, nil
}
