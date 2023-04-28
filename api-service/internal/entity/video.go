package entity

import (
	"api-service/internal/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Video struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Description string             `bson:"description,omitempty"`
	Views       int64              `bson:"views,omitempty"`
	Likes       int64              `bson:"likes,omitempty"`
	Dislikes    int64              `bson:"dislikes,omitempty"`
	Tags        []string           `bson:"tags,omitempty"`
	AuthorID    primitive.ObjectID `bson:"author_id,omitempty"`
	UploadedAt  primitive.DateTime `bson:"uploadedAt,omitempty"`
	Duration    int64              `bson:"duration,omitempty"`
	Comments    []Comment          `bson:"comments,omitempty"`
}

func VideoFromDto(dto *dto.VideoDto) Video {
	var v Video

	v.ID = dto.ID
	v.Title = dto.Title
	v.Description = dto.Description
	v.Views = dto.Views
	v.Likes = dto.Likes
	v.Dislikes = dto.Dislikes
	v.Tags = dto.Tags
	v.AuthorID = dto.AuthorID
	v.UploadedAt = primitive.NewDateTimeFromTime(dto.UploadedAt)
	v.Duration = dto.Duration

	comments := make([]Comment, len(dto.Comments))
	for i, c := range dto.Comments {
		comments[i] = Comment{
			AuthorID:  c.AuthorID,
			Content:   c.Content,
			CreatedAt: primitive.NewDateTimeFromTime(c.CreatedAt),
			Replies:   make([]Reply, len(c.Replies)),
		}

		for j, r := range c.Replies {
			comments[i].Replies[j] = Reply{
				AuthorID:  r.AuthorID,
				Content:   r.Content,
				CreatedAt: primitive.NewDateTimeFromTime(r.CreatedAt),
			}
		}
	}

	v.Comments = comments

	return v
}
