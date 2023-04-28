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

type VideoRepositoryImpl struct {
	collection *mongo.Collection
}

func NewVideoRepository(client *mongo.Client) *VideoRepositoryImpl {
	return &VideoRepositoryImpl{collection: client.Database("videos").Collection("videos")}
}

func (r *VideoRepositoryImpl) FindOne(ctx context.Context, id primitive.ObjectID) (entity.Video, error) {
	var video entity.Video

	err := r.collection.FindOne(ctx, bson.M{"id": id}).Decode(&video)

	if err != nil {
		return entity.Video{}, err //TODO: nil
	}

	return video, nil
}

func (r *VideoRepositoryImpl) FindAll(ctx context.Context) ([]entity.Video, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var videos []entity.Video
	for cursor.Next(ctx) {
		var video entity.Video
		if err := cursor.Decode(&video); err != nil {
			return nil, err
		}
		videos = append(videos, video)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return videos, nil
}

func (r *VideoRepositoryImpl) FindTop10(ctx context.Context) ([]entity.Video, error) {
	sortStage := bson.D{{"$sort", bson.D{{"views", -1}}}}
	limitStage := bson.D{{"$limit", 10}}
	cursor, err := r.collection.Aggregate(ctx, mongo.Pipeline{sortStage, limitStage})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	var videos []entity.Video
	for cursor.Next(ctx) {
		var video entity.Video
		if err := cursor.Decode(&video); err != nil {
			return nil, err
		}
		videos = append(videos, video)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	return videos, nil
}

func (r *VideoRepositoryImpl) FindTagsRate(ctx context.Context) ([]map[string]interface{}, error) {
	unwindStage := bson.D{{"$unwind", "$tags"}}
	groupStage := bson.D{{"$group", bson.D{{"_id", "$tags"}, {"count", bson.D{{"$sum", 1}}}}}}
	sortStage := bson.D{{"$sort", bson.D{{"count", -1}}}}

	cursor, err := r.collection.Aggregate(ctx, mongo.Pipeline{unwindStage, groupStage, sortStage})
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

func (r *VideoRepositoryImpl) Add(ctx context.Context, video entity.Video) (primitive.ObjectID, error) {
	fmt.Println(video)

	res, err := r.collection.InsertOne(ctx, video)

	if err != nil {
		return primitive.NilObjectID, err
	}

	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, fmt.Errorf("failed to get ID of inserted video")
	}

	return id, nil
}
