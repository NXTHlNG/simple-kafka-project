package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Reply struct {
	AuthorID primitive.ObjectID `bson:"author_id,omitempty"`
	Content  string             `bson:"content,omitempty"`
}
