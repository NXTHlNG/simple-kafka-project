package service

import (
	"context"
	"data-service/internal/entity"
	"data-service/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthorServiceImpl struct {
	repository repository.AuthorRepository
}

func (v AuthorServiceImpl) Create(ctx context.Context, author entity.Author) (primitive.ObjectID, error) {
	return v.repository.Add(ctx, author)
}

func (v AuthorServiceImpl) Get(ctx context.Context, id primitive.ObjectID) (entity.Author, error) {
	return v.repository.FindOne(ctx, id)
}

func (v AuthorServiceImpl) GetAll(ctx context.Context) ([]entity.Author, error) {
	return v.repository.FindAll(ctx)
}

func (v AuthorServiceImpl) GetTop10(ctx context.Context) ([]map[string]interface{}, error) {
	return v.repository.FindTop10(ctx)
}

func NewAuthorServiceImpl(repository repository.AuthorRepository) *AuthorServiceImpl {
	return &AuthorServiceImpl{repository: repository}
}
