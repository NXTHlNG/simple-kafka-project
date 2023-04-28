package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type AuthorDto struct {
	Id     primitive.ObjectID `json:"id" bson:"_id"  binding:"required"`
	Name   string             `json:"name" bson:"name" binding:"required"`
	Email  string             `json:"email" bson:"email" binding:"required"`
	Avatar string             `json:"avatar" bson:"avatar" binding:"required"`
}
