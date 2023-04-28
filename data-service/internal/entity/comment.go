package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Comment struct {
	AuthorID primitive.ObjectID `bson:"author_id,omitempty"`
	Content  string             `bson:"content,omitempty"`
	Replies  []Reply            `bson:"replies,omitempty"`
}
