package postgres

import (
	"context"
	"time"

	gen "learning/genprotos"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// UserProgress struct provides the database connection
type UserProgress struct {
	Db *pgx.Conn
}

// NewUserProgress creates a new instance of UserProgress
func NewUserProgress(db *pgx.Conn) *UserProgress {
	return &UserProgress{Db: db}
}

// AddUserProgress inserts a new user progress into the database
func (r *UserProgress) AddUserProgress(ctx context.Context, req *gen.AddUserProgressRequest) (*gen.AddUserProgressResponse, error) {
	query := `
		INSERT INTO user_progress (user_id, language_id, lesson_id, exercise_id, vocabulary_id, completed_at)
		VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := r.Db.Exec(ctx, query,
		req.UserId,
		req.LanguageId,
		req.LessonId,
		req.ExerciseId,
		req.VocabularyId,
		req.CompletedAt.AsTime(),
	)

	if err != nil {
		log.Error().Err(err).Msg("Error adding user progress")
		return nil, err
	}

	return &gen.AddUserProgressResponse{}, nil
}

// GetUserProgress retrieves user progress by user ID
func (r *UserProgress) GetUserProgress(ctx context.Context, req *gen.GetUserProgressRequest) (*gen.GetUserProgressResponse, error) {
	query := `
		SELECT user_id, language_id, lesson_id, exercise_id, vocabulary_id, completed_at
		FROM user_progress
		WHERE user_id = $1`

	rows, err := r.Db.Query(ctx, query, req.UserId)
	if err != nil {
		log.Error().Err(err).Msg("Error retrieving user progress")
		return nil, err
	}
	defer rows.Close()

	var userProgressList []*gen.UserProgress

	for rows.Next() {
		up := &gen.UserProgress{}
		var completedAt time.Time

		err := rows.Scan(
			&up.UserId,
			&up.LanguageId,
			&up.LessonId,
			&up.ExerciseId,
			&up.VocabularyId,
			&completedAt,
		)
		if err != nil {
			log.Error().Err(err).Msg("Error scanning user progress")
			return nil, err
		}

		up.CompletedAt = timestamppb.New(completedAt)
		userProgressList = append(userProgressList, up)
	}

	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg("Error with rows in get user progress")
		return nil, err
	}

	return &gen.GetUserProgressResponse{
		UserProgressList: userProgressList,
	}, nil
}
