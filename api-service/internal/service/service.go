package service

import (
	"api-service/internal/entity"
	"context"
	"github.com/Shopify/sarama"
)

type VideoService interface {
	Create(ctx context.Context, film entity.Video) error
	GetAll(ctx context.Context) ([]map[string]interface{}, error)
	GetTagsRate(ctx context.Context) ([]map[string]interface{}, error)
	GetTop10Videos(ctx context.Context) ([]map[string]interface{}, error)
}

type AuthorService interface {
	Create(ctx context.Context, film entity.Author) error
	GetAll(ctx context.Context) ([]map[string]interface{}, error)
	GetTop10Authors(ctx context.Context) ([]map[string]interface{}, error)
}

type Service struct {
	VideoService
	AuthorService
}

func NewService(messageProducer sarama.SyncProducer) *Service {
	return &Service{
		VideoService:  NewVideoServiceImpl(messageProducer),
		AuthorService: NewAuthorServiceImpl(messageProducer),
	}
}
