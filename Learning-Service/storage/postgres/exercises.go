package postgres

import (
	"context"

	pb "learning/genprotos" // Protobufdan generatsiya qilingan Go paketiga yo'lni o'zgartiring

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// QuestionRepo - savollar uchun repository
type QuestionRepo struct {
	dbmongo *mongo.Collection
}

// NewQuestionRepo - yangi savollar repository sini yaratish
func NewExercises(db *mongo.Collection) *QuestionRepo {
	return &QuestionRepo{
		dbmongo: db,
	}
}

// CreateQuestion - savol yaratish metodi
func (r *QuestionRepo) CreateQuestion(ctx context.Context, req *pb.CreateQuestionRequest) (*pb.CreateQuestionResponse, error) {

	question := &pb.Question{
		Id:            "", // ID avtomatik yaratiladi
		Type:          req.Type,
		LanguageCode:  req.LanguageCode,
		Difficulty:    req.Difficulty,
		Question:      req.Question,
		Options:       req.Options,
		CorrectAnswer: req.CorrectAnswer,
		Explanation:   req.Explanation,
		CreatedAt:     timestamppb.Now(),
	}

	result, err := r.dbmongo.InsertOne(ctx, question)
	if err != nil {
		return nil, err
	}

	question.Id = result.InsertedID.(string)
	return &pb.CreateQuestionResponse{Id: question.Id}, nil
}

// GetQuestion - savolni olish metodi
func (r *QuestionRepo) GetQuestion(ctx context.Context, req *pb.GetQuestionRequest) (*pb.GetQuestionResponse, error) {

	filter := bson.M{"_id": req.Id}
	var question pb.Question
	err := r.dbmongo.FindOne(ctx, filter).Decode(&question)
	if err != nil {
		return nil, err
	}

	return &pb.GetQuestionResponse{Question: &question}, nil
}

// ListQuestions - barcha savollar ro'yxatini olish metodi
func (r *QuestionRepo) ListQuestions(ctx context.Context, req *pb.ListQuestionsRequest) (*pb.ListQuestionsResponse, error) {

	cursor, err := r.dbmongo.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var questions []*pb.Question
	for cursor.Next(ctx) {
		var question pb.Question
		if err := cursor.Decode(&question); err != nil {
			return nil, err
		}
		questions = append(questions, &question)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &pb.ListQuestionsResponse{Questions: questions}, nil
}

// UpdateQuestion - savolni yangilash metodi
func (r *QuestionRepo) UpdateQuestion(ctx context.Context, req *pb.UpdateQuestionRequest) (*pb.UpdateQuestionResponse, error) {
	
	filter := bson.M{"_id": req.Id}
	update := bson.M{
		"$set": bson.M{
			"type":           req.Type,
			"language_code":  req.LanguageCode,
			"difficulty":     req.Difficulty,
			"question":       req.Question,
			"options":        req.Options,
			"correct_answer": req.CorrectAnswer,
			"explanation":    req.Explanation,
		},
	}

	_, err := r.dbmongo.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	var updatedQuestion pb.Question
	err = r.dbmongo.FindOne(ctx, filter).Decode(&updatedQuestion)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateQuestionResponse{Question: &updatedQuestion}, nil
}

// DeleteQuestion - savolni o'chirish metodi
func (r *QuestionRepo) DeleteQuestion(ctx context.Context, req *pb.DeleteQuestionRequest) (*pb.DeleteQuestionResponse, error) {

	filter := bson.M{"_id": req.Id}

	_, err := r.dbmongo.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteQuestionResponse{Id: req.Id}, nil
}
