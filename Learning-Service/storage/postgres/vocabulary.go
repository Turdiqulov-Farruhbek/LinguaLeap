package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
	gen "learning/genprotos"
)

// Vocabulary struct provides the database connection
type Vocabulary struct {
	Db *pgx.Conn
}

// NewVocabulary creates a new instance of Vocabulary
func NewVocabulary(db *pgx.Conn) *Vocabulary {
	return &Vocabulary{Db: db}
}

// AddVocabulary inserts a new vocabulary into the database
func (r *Vocabulary) AddVocabulary(ctx context.Context, req *gen.AddVocabularyRequest) (*gen.AddVocabularyResponse, error) {
	vocabularyID := uuid.NewString()

	query := `
		INSERT INTO vocabularies (id, language_code, word, translation, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`

	var id string
	err := r.Db.QueryRow(ctx, query,
		vocabularyID,
		req.LanguageCode,
		req.Word,
		req.Translation,
		time.Now(),
	).Scan(&id)

	if err != nil {
		log.Error().Err(err).Msg("Error adding vocabulary")
		return nil, err
	}

	return &gen.AddVocabularyResponse{
		Id: id,
	}, nil
}

// GetVocabulary retrieves a vocabulary by its ID
func (r *Vocabulary) GetVocabulary(ctx context.Context, req *gen.GetVocabularyRequest) (*gen.GetVocabularyResponse, error) {
	query := `
		SELECT id, language_code, word, translation, created_at
		FROM vocabularies
		WHERE id = $1`

	vocab := &gen.Vocabulary{}
	var createdAt time.Time

	err := r.Db.QueryRow(ctx, query, req.Id).Scan(
		&vocab.Id,
		&vocab.LanguageCode,
		&vocab.Word,
		&vocab.Translation,
		&createdAt,
	)

	if err != nil {
		log.Error().Err(err).Msg("Error retrieving vocabulary")
		return nil, err
	}

	vocab.CreatedAt = createdAt.Format(time.RFC3339)

	return &gen.GetVocabularyResponse{
		Vocabulary: vocab,
	}, nil
}

// UpdateVocabulary updates a vocabulary in the database
func (r *Vocabulary) UpdateVocabulary(ctx context.Context, req *gen.UpdateVocabularyRequest) (*gen.UpdateVocabularyResponse, error) {
	query := `
		UPDATE vocabularies
		SET language_code = $2, word = $3, translation = $4, updated_at = NOW()
		WHERE id = $1
		RETURNING id, language_code, word, translation, created_at, updated_at`

	vocab := &gen.Vocabulary{}
	var createdAt, updatedAt time.Time

	err := r.Db.QueryRow(ctx, query,
		req.Id,
		req.LanguageCode,
		req.Word,
		req.Translation,
	).Scan(
		&vocab.Id,
		&vocab.LanguageCode,
		&vocab.Word,
		&vocab.Translation,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		log.Error().Err(err).Msg("Error updating vocabulary")
		return nil, err
	}

	vocab.CreatedAt = createdAt.Format(time.RFC3339)
	vocab.UpdatedAt = updatedAt.Format(time.RFC3339)

	return &gen.UpdateVocabularyResponse{
		Vocabulary: vocab,
	}, nil
}

// DeleteVocabulary deletes a vocabulary from the database
func (r *Vocabulary) DeleteVocabulary(ctx context.Context, req *gen.DeleteVocabularyRequest) (*gen.DeleteVocabularyResponse, error) {
	query := `
		DELETE FROM vocabularies
		WHERE id = $1
		RETURNING id`

	var id string
	err := r.Db.QueryRow(ctx, query, req.Id).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("vocabulary not found")
		}
		log.Error().Err(err).Msg("Error deleting vocabulary")
		return nil, err
	}

	return &gen.DeleteVocabularyResponse{
		Id: id,
	}, nil
}
