package main

import (
	"context"
	"data-service/internal/config"
	"data-service/internal/controller"
	"data-service/internal/entity"
	"data-service/internal/repository"
	"data-service/internal/routes"
	"data-service/internal/service"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/gofiber/fiber/v2"
)

type ConsumerMessage struct {
	RoutingKey string `json:"routingKey"`
}

type ConsumerVideoMessage struct {
	RoutingKey string       `json:"routingKey"`
	Video      entity.Video `json:"data"`
}

type ConsumerAuthorMessage struct {
	RoutingKey string        `json:"routingKey"`
	Author     entity.Author `json:"data"`
}

func main() {
	client := config.NewMongoClient()
	repository := repository.NewRepository(client)
	service := service.NewService(repository)
	controller := controller.NewController(service)

	app := fiber.New()

	routes.VideoRoute(app, controller.VideoController)
	routes.AuthorRoute(app, controller.AuthorController)

	go func() {
		err := app.Listen(":8001")
		if err != nil {
			return
		}
	}()

	//Set up configuration for the consumer
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	//Create new consumer
	consumer, err := sarama.NewConsumer([]string{"broker:9092"}, config)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			panic(err)
		}
	}()

	//Subscribe to the topic
	topic := "videos"
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			panic(err)
		}
	}()

	// Start consuming messages
	for message := range partitionConsumer.Messages() {
		var consumerMessage ConsumerMessage

		fmt.Printf("Message value: %s\n", string(message.Value))

		err = json.Unmarshal(message.Value, &consumerMessage)
		if err != nil {
			fmt.Printf("Catalog consumer handler error: %s", err)
			continue
		}

		if consumerMessage.RoutingKey == "video" {
			fmt.Println("ПРИШЛО ВИДЕО")
			var videoMessage ConsumerVideoMessage
			err = json.Unmarshal(message.Value, &videoMessage)
			fmt.Println(videoMessage)
			fmt.Println(videoMessage.Video)
			service.VideoService.Create(context.TODO(), videoMessage.Video)
		}

		if consumerMessage.RoutingKey == "author" {
			fmt.Println("ПРИШЁЛ АВТОР")
			var authorMessage ConsumerAuthorMessage
			err = json.Unmarshal(message.Value, &authorMessage)
			fmt.Println(authorMessage)
			fmt.Println(authorMessage.Author)
			service.AuthorService.Create(context.TODO(), authorMessage.Author)
		}

	}

	//	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, sarama.NewTestConfig())
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	defer func() {
	//		if err := consumer.Close(); err != nil {
	//			log.Fatalln(err)
	//		}
	//	}()
	//
	//	partitionConsumer, err := consumer.ConsumePartition("my_topic", 0, OffsetNewest)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	defer func() {
	//		if err := partitionConsumer.Close(); err != nil {
	//			log.Fatalln(err)
	//		}
	//	}()
	//
	//	// Trap SIGINT to trigger a shutdown.
	//	signals := make(chan os.Signal, 1)
	//	signal.Notify(signals, os.Interrupt)
	//
	//	consumed := 0
	//ConsumerLoop:
	//	for {
	//		select {
	//		case msg := <-partitionConsumer.Messages():
	//			log.Printf("Consumed message offset %d\n", msg.Offset)
	//			consumed++
	//		case <-signals:
	//			break ConsumerLoop
	//		}
	//	}

}
