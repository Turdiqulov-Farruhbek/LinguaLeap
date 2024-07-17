package postgres

import (
	"context"
	"fmt"
	"learning/config"
	stg "learning/storage"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StorageStruct struct {
	DB            *pgx.Conn
	Collection    *mongo.Collection
	Language_S    stg.LanguageInterface
	Lessons_S     stg.LessonsInterface
	UserProgres_S stg.UserProgresInterface
	Vocabulary_S  stg.VocabularyInterface
	Exercises_S   stg.ExercisesInterface
	User_S        stg.UserInterface
	UserLessons_S stg.UserLessonsInterface
}

func DBConn() (*StorageStruct, error) {
	var (
		db  *pgx.Conn
		err error
	)

// postgresDB connection
	cfg := config.Load()
	dbCon := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase)

	db, err = pgx.Connect(context.Background(), dbCon)
	if err != nil {
		slog.Warn("Unable to connect to database:", "eroRRR", err)
		return nil, err
	}
	err = db.Ping(context.Background())
	if err != nil {
		slog.Warn("Unable to ping database:", "eroRRR", err)
		return nil, err
	}


// MongoDB connection
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("MongoDB'ga ulanish muvaffaqiyatli bo'ldi!")

	collection := client.Database("learning_db").Collection("axercises")


	return &StorageStruct{
		DB: db,
		Collection: collection,
		Language_S:    NewLanguage(db),
        Lessons_S:     NewLessons(db),
        UserProgres_S: NewUserProgress(db),
        Vocabulary_S:  NewVocabulary(db),
        Exercises_S:   NewExercises(collection),
	}, nil
}



func (s *StorageStruct) Languages() stg.LanguageInterface {
	if s.Language_S == nil {
		s.Language_S = NewLanguage(s.DB)
	}

	return s.Language_S
}

func (s *StorageStruct) Lessons() stg.LessonsInterface {
	if s.Lessons_S == nil {
		s.Lessons_S = NewLessons(s.DB)
	}

	return s.Lessons_S
}

func (s *StorageStruct) UserProgress() stg.UserProgresInterface {
	if s.UserProgres_S == nil {
		s.UserProgres_S = NewUserProgress(s.DB)
	}

	return s.UserProgres_S
}

func (s *StorageStruct) Vocabularies() stg.VocabularyInterface {
	if s.Vocabulary_S == nil {
		s.Vocabulary_S = NewVocabulary(s.DB)
	}

	return s.Vocabulary_S
}

func (s *StorageStruct) Exercises() stg.ExercisesInterface {
	if s.Exercises_S == nil {
        s.Exercises_S = NewExercises(s.Collection)
    }

    return s.Exercises_S
}

func (s *StorageStruct) Users() stg.UserInterface {
	if s.User_S == nil {
        s.User_S = NewUser(s.DB)
    }

    return s.User_S
}

func (s *StorageStruct) UserLessons() stg.UserLessonsInterface {
	if s.UserLessons_S == nil {
        s.UserLessons_S = NewUserLessons(s.DB)
    }

    return s.UserLessons_S
}



