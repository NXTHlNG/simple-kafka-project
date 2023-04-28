package service

import (
	"context"
	"data-service/internal/entity"
	"data-service/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type VideoService interface {
	Create(ctx context.Context, video entity.Video) (primitive.ObjectID, error)
	Get(ctx context.Context, id primitive.ObjectID) (entity.Video, error)
	GetAll(ctx context.Context) ([]entity.Video, error)
	GetTop10(ctx context.Context) ([]entity.Video, error)
	GetTagsRate(ctx context.Context) ([]map[string]interface{}, error)
}

type AuthorService interface {
	Create(ctx context.Context, author entity.Author) (primitive.ObjectID, error)
	Get(ctx context.Context, id primitive.ObjectID) (entity.Author, error)
	GetAll(ctx context.Context) ([]entity.Author, error)
	GetTop10(ctx context.Context) ([]map[string]interface{}, error)
}

type Service struct {
	VideoService
	AuthorService
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		VideoService:  NewVideoServiceImpl(repository.VideoRepository),
		AuthorService: NewAuthorServiceImpl(repository.AuthorRepository),
	}
}
