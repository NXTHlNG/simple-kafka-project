package config

import (
	"context"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func NewMongoClient() *mongo.Client {
	mongodbUri, exist := os.LookupEnv("MONGODB_URI")
	if !exist {
		mongodbUri = "localhost:27017"
	}

	mongodbUsername, exist := os.LookupEnv("MONGODB_USERNAME")
	if !exist {
		mongodbUri = "root"
	}

	mongodbPassword, exist := os.LookupEnv("MONGODB_PASSWORD")
	if !exist {
		mongodbUri = "root"
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s", mongodbUri)).SetAuth(options.Credential{
		Username: mongodbUsername,
		Password: mongodbPassword,
	}))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")

	driver, err := mongodb.WithInstance(client, &mongodb.Config{
		DatabaseName: "videos",
	})

	if err != nil {
		logrus.Panicln(err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./db/migrations",
		"videos", driver)

	if err != nil {
		logrus.Panicln(err.Error())
	}

	err = m.Up()

	if err != nil && err != migrate.ErrNoChange {
		logrus.Panicln(err.Error())
	}

	return client
}
