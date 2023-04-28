package repository

import (
	"context"
	"data-service/internal/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type VideoRepository interface {
	Add(ctx context.Context, video entity.Video) (primitive.ObjectID, error)
	FindOne(ctx context.Context, id primitive.ObjectID) (entity.Video, error)
	FindAll(ctx context.Context) ([]entity.Video, error)
	FindTop10(ctx context.Context) ([]entity.Video, error)
	FindTagsRate(ctx context.Context) ([]map[string]interface{}, error)
}

type AuthorRepository interface {
	Add(ctx context.Context, author entity.Author) (primitive.ObjectID, error)
	FindOne(ctx context.Context, id primitive.ObjectID) (entity.Author, error)
	FindAll(ctx context.Context) ([]entity.Author, error)
	FindTop10(ctx context.Context) ([]map[string]interface{}, error)
}

//type ReviewsStorage interface {
//	Add(ctx context.Context, review entity.Review) error
//	GetAll(ctx context.Context) ([]entity.Review, error)
//}

type Repository struct {
	VideoRepository
	AuthorRepository
}

func NewRepository(client *mongo.Client) *Repository {
	return &Repository{
		VideoRepository:  NewVideoRepository(client),
		AuthorRepository: NewAuthorRepository(client),
	}
}
