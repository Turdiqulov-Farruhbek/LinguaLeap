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

// Language struct provides the database connection
type Language struct {
	Db *pgx.Conn
}

// NewLanguage creates a new instance of Language
func NewLanguage(db *pgx.Conn) *Language {
	return &Language{Db: db}
}

// CreateLanguage inserts a new language into the database
func (r *Language) CreateLanguage(ctx context.Context, req *gen.CreateLanguageRequest) (*gen.CreateLanguageResponse, error) {
	languageID := uuid.NewString()

	query := `
		INSERT INTO languages (id, language_code, name, emoji)
		VALUES ($1, $2, $3, $4)
		RETURNING id, language_code, name, emoji, created_at`

	langRes := &gen.Language{}
	var createdAt time.Time

	err := r.Db.QueryRow(ctx, query,
		languageID,
		req.LanguageCode,
		req.Name,
		req.Emoji,
	).Scan(
		&langRes.Id,
		&langRes.LanguageCode,
		&langRes.Name,
		&langRes.Emoji,
		&createdAt,
	)

	if err != nil {
		log.Error().Err(err).Msg("Error creating language")
		return nil, err
	}

	langRes.CreatedAt = createdAt.Format(time.RFC3339)

	return &gen.CreateLanguageResponse{
		Language: langRes,
	}, nil
}

// GetLanguage retrieves a language by its ID
func (r *Language) GetLanguage(ctx context.Context, req *gen.GetLanguageRequest) (*gen.GetLanguageResponse, error) {
	query := `
		SELECT id, language_code, name, emoji, created_at
		FROM languages
		WHERE id = $1`

	langRes := &gen.Language{}
	var createdAt time.Time

	err := r.Db.QueryRow(ctx, query, req.Id).Scan(
		&langRes.Id,
		&langRes.LanguageCode,
		&langRes.Name,
		&langRes.Emoji,
		&createdAt,
	)

	if err != nil {
		log.Error().Err(err).Msg("Error retrieving language")
		return nil, err
	}

	langRes.CreatedAt = createdAt.Format(time.RFC3339)

	return &gen.GetLanguageResponse{
		Language: langRes,
	}, nil
}

// ListLanguages retrieves all languages from the database
func (r *Language) ListLanguages(ctx context.Context, req *gen.ListLanguagesRequest) (*gen.ListLanguagesResponse, error) {
	query := `
		SELECT id, language_code, name, emoji, created_at
		FROM languages`

	rows, err := r.Db.Query(ctx, query)
	if err != nil {
		log.Error().Err(err).Msg("Error listing languages")
		return nil, err
	}
	defer rows.Close()

	var languages []*gen.Language

	for rows.Next() {
		lang := &gen.Language{}
		var createdAt time.Time

		err := rows.Scan(
			&lang.Id,
			&lang.LanguageCode,
			&lang.Name,
			&lang.Emoji,
			&createdAt,
		)
		if err != nil {
			log.Error().Err(err).Msg("Error scanning language")
			return nil, err
		}

		lang.CreatedAt = createdAt.Format(time.RFC3339)
		languages = append(languages, lang)
	}

	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg("Error with rows in list languages")
		return nil, err
	}

	return &gen.ListLanguagesResponse{
		Languages: languages,
	}, nil
}

// UpdateLanguage updates a language in the database
func (r *Language) UpdateLanguage(ctx context.Context, req *gen.UpdateLanguageRequest) (*gen.UpdateLanguageResponse, error) {
	query := `
		UPDATE languages
		SET language_code = $2, name = $3, emoji = $4, updated_at = NOW()
		WHERE id = $1
		RETURNING id, language_code, name, emoji, created_at, updated_at`

	langRes := &gen.Language{}
	var createdAt, updatedAt time.Time

	err := r.Db.QueryRow(ctx, query,
		req.Id,
		req.LanguageCode,
		req.Name,
		req.Emoji,
	).Scan(
		&langRes.Id,
		&langRes.LanguageCode,
		&langRes.Name,
		&langRes.Emoji,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		log.Error().Err(err).Msg("Error updating language")
		return nil, err
	}

	langRes.CreatedAt = createdAt.Format(time.RFC3339)
	langRes.UpdatedAt = updatedAt.Format(time.RFC3339)

	return &gen.UpdateLanguageResponse{
		Language: langRes,
	}, nil
}

// DeleteLanguage deletes a language from the database
func (r *Language) DeleteLanguage(ctx context.Context, req *gen.DeleteLanguageRequest) (*gen.DeleteLanguageResponse, error) {
	query := `
		DELETE FROM languages
		WHERE id = $1
		RETURNING id`

	var id string
	err := r.Db.QueryRow(ctx, query, req.Id).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("language not found")
		}
		log.Error().Err(err).Msg("Error deleting language")
		return nil, err
	}

	return &gen.DeleteLanguageResponse{
		Id: id,
	}, nil
}
