package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ReplyDto struct {
	AuthorID  primitive.ObjectID `json:"author_id" bson:"author_id"`
	Content   string             `json:"content" bson:"content"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
}
