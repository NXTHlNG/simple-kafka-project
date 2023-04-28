package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type VideoDto struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Views       int64              `json:"views" bson:"views"`
	Likes       int64              `json:"likes" bson:"likes"`
	Dislikes    int64              `json:"dislikes" bson:"dislikes"`
	Tags        []string           `json:"tags" bson:"tags"`
	AuthorID    primitive.ObjectID `json:"author_id" bson:"author_id"`
	UploadedAt  time.Time          `json:"uploadedAt" bson:"uploadedAt"`
	Duration    int64              `json:"duration" bson:"duration"`
	Comments    []CommentDto       `json:"comments" bson:"comments"`
}
