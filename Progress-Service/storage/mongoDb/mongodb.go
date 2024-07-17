package mongoDB

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	Db     *mongo.Collection
}

func NewMongoDBStorage() (*Storage, error) {
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

	collection := client.Database("learning_db").Collection("language")

	return &Storage{
		Db: collection,
	}, nil
}

