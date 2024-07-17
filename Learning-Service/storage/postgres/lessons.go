package postgres

import (
	"context"
	"time"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
	gen "learning/genprotos"
)

// Lesson struct provides the database connection
type Lesson struct {
	Db *pgx.Conn
}

// NewLesson creates a new instance of Lesson
func NewLessons(db *pgx.Conn) *Lesson {
	return &Lesson{Db: db}
}

// CreateLesson inserts a new lesson into the database
func (r *Lesson) CreateLesson(ctx context.Context, req *gen.CreateLessonRequest) (*gen.CreateLessonResponse, error) {
	lessonID := uuid.NewString()

	query := `
		INSERT INTO lessons (id, language_code, title, level, created_at)
		VALUES ($1, $2, $3, $4, NOW())
		RETURNING id`

	var id string
	err := r.Db.QueryRow(ctx, query,
		lessonID,
		req.LanguageCode,
		req.Title,
		req.Level,
	).Scan(&id)

	if err != nil {
		log.Error().Err(err).Msg("Error creating lesson")
		return nil, err
	}

	return &gen.CreateLessonResponse{
		Id: id,
	}, nil
}

// GetLesson retrieves a lesson by its ID
func (r *Lesson) GetLesson(ctx context.Context, req *gen.GetLessonRequest) (*gen.GetLessonResponse, error) {
	query := `
		SELECT id, language_code, title, level, created_at
		FROM lessons
		WHERE id = $1`

	lesson := &gen.Lesson{}
	var createdAt time.Time

	err := r.Db.QueryRow(ctx, query, req.Id).Scan(
		&lesson.Id,
		&lesson.LanguageCode,
		&lesson.Title,
		&lesson.Level,
		&createdAt,
	)

	if err != nil {
		log.Error().Err(err).Msg("Error retrieving lesson")
		return nil, err
	}

	lesson.CreatedAt = createdAt.Format(time.RFC3339)

	return &gen.GetLessonResponse{
		Lesson: lesson,
	}, nil
}

// ListLessons retrieves all lessons from the database
func (r *Lesson) ListLessons(ctx context.Context, req *gen.ListLessonsRequest) (*gen.ListLessonsResponse, error) {
	query := `
		SELECT id, language_code, title, level, created_at
		FROM lessons`

	rows, err := r.Db.Query(ctx, query)
	if err != nil {
		log.Error().Err(err).Msg("Error listing lessons")
		return nil, err
	}
	defer rows.Close()

	var lessons []*gen.Lesson

	for rows.Next() {
		lesson := &gen.Lesson{}
		var createdAt time.Time

		err := rows.Scan(
			&lesson.Id,
			&lesson.LanguageCode,
			&lesson.Title,
			&lesson.Level,
			&createdAt,
		)
		if err != nil {
			log.Error().Err(err).Msg("Error scanning lesson")
			return nil, err
		}

		lesson.CreatedAt = createdAt.Format(time.RFC3339)
		lessons = append(lessons, lesson)
	}

	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg("Error with rows in list lessons")
		return nil, err
	}

	return &gen.ListLessonsResponse{
		Lessons: lessons,
	}, nil
}

// UpdateLesson updates a lesson in the database
func (r *Lesson) UpdateLesson(ctx context.Context, req *gen.UpdateLessonRequest) (*gen.UpdateLessonResponse, error) {
	query := `
		UPDATE lessons
		SET language_code = $2, title = $3, level = $4, updated_at = NOW()
		WHERE id = $1
		RETURNING id, language_code, title, level, created_at, updated_at`

	lesson := &gen.Lesson{}
	var createdAt, updatedAt time.Time

	err := r.Db.QueryRow(ctx, query,
		req.Id,
		req.LanguageCode,
		req.Title,
		req.Level,
	).Scan(
		&lesson.Id,
		&lesson.LanguageCode,
		&lesson.Title,
		&lesson.Level,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		log.Error().Err(err).Msg("Error updating lesson")
		return nil, err
	}

	lesson.CreatedAt = createdAt.Format(time.RFC3339)
	lesson.UpdatedAt = updatedAt.Format(time.RFC3339)

	return &gen.UpdateLessonResponse{
		Lesson: lesson,
	}, nil
}

// DeleteLesson deletes a lesson from the database
func (r *Lesson) DeleteLesson(ctx context.Context, req *gen.DeleteLessonRequest) (*gen.DeleteLessonResponse, error) {
	query := `
		DELETE FROM lessons
		WHERE id = $1
		RETURNING id`

	var id string
	err := r.Db.QueryRow(ctx, query, req.Id).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("lesson not found")
		}
		log.Error().Err(err).Msg("Error deleting lesson")
		return nil, err
	}

	return &gen.DeleteLessonResponse{
		Id: id,
	}, nil
}
