package service

import (
	"context"
	"data-service/internal/entity"
	"data-service/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type VideoServiceImpl struct {
	repository repository.VideoRepository
}

func (v VideoServiceImpl) Create(ctx context.Context, video entity.Video) (primitive.ObjectID, error) {
	return v.repository.Add(ctx, video)
}

func (v VideoServiceImpl) Get(ctx context.Context, id primitive.ObjectID) (entity.Video, error) {
	return v.repository.FindOne(ctx, id)
}

func (v VideoServiceImpl) GetAll(ctx context.Context) ([]entity.Video, error) {
	return v.repository.FindAll(ctx)
}

func (v VideoServiceImpl) GetTop10(ctx context.Context) ([]entity.Video, error) {
	return v.repository.FindTop10(ctx)
}

func (v VideoServiceImpl) GetTagsRate(ctx context.Context) ([]map[string]interface{}, error) {
	return v.repository.FindTagsRate(ctx)
}

func NewVideoServiceImpl(repository repository.VideoRepository) *VideoServiceImpl {
	return &VideoServiceImpl{repository: repository}
}
