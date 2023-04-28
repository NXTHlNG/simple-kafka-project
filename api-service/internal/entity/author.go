package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Author struct {
	Id     primitive.ObjectID `json:"id" bson:"_id"  binding:"required"`
	Name   string             `json:"name" bson:"name" binding:"required"`
	Email  string             `json:"email" bson:"hours" binding:"required"`
	Avatar string             `json:"avatar" bson:"cost" binding:"required"`
}
