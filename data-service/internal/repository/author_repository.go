package repository

import (
	"context"
	"data-service/internal/entity"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type AuthorRepositoryImpl struct {
	collection *mongo.Collection
}

func NewAuthorRepository(client *mongo.Client) *AuthorRepositoryImpl {
	return &AuthorRepositoryImpl{collection: client.Database("videos").Collection("authors")}
}

func (r *AuthorRepositoryImpl) FindOne(ctx context.Context, id primitive.ObjectID) (entity.Author, error) {
	var author entity.Author

	err := r.collection.FindOne(ctx, bson.M{"id": id}).Decode(&author)

	if err != nil {
		return entity.Author{}, err //TODO: nil
	}

	return author, nil
}

func (r *AuthorRepositoryImpl) FindAll(ctx context.Context) ([]entity.Author, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var authors []entity.Author
	for cursor.Next(ctx) {
		var author entity.Author
		if err := cursor.Decode(&author); err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return authors, nil
}

func (r *AuthorRepositoryImpl) FindTop10(ctx context.Context) ([]map[string]interface{}, error) {
	lookupStage := bson.D{{"$lookup", bson.D{
		{"from", "videos"},
		{"localField", "_id"},
		{"foreignField", "author_id"},
		{"as", "video"},
	}}}
	unwindStage := bson.D{{"$unwind", "$video"}}
	groupStage := bson.D{{"$group", bson.D{
		{"_id", "$_id"},
		{"name", bson.D{{"$first", "$name"}}},
		{"count", bson.D{{"$sum", "$video.views"}}},
	}}}
	sortStage := bson.D{{"$sort", bson.D{{"count", -1}, {"name", 1}}}}
	limitStage := bson.D{{"$limit", 10}}

	cursor, err := r.collection.Aggregate(ctx, mongo.Pipeline{lookupStage, unwindStage, groupStage, sortStage, limitStage})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	var results []map[string]interface{}
	for cursor.Next(ctx) {
		var result bson.M
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	return results, nil
}

func (r *AuthorRepositoryImpl) Add(ctx context.Context, author entity.Author) (primitive.ObjectID, error) {
	fmt.Println(author)

	res, err := r.collection.InsertOne(ctx, author)

	if err != nil {
		return primitive.NilObjectID, err
	}

	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, fmt.Errorf("failed to get ID of inserted video")
	}

	return id, nil
}
