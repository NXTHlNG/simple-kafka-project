package main

import (
	"api-service/internal/controller"
	"api-service/internal/routes"
	"api-service/internal/service"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{"broker:9092"}, config)
	if err != nil {
		log.Fatalf("Error creating Kafka producer: %s", err.Error())
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalf("Error closing Kafka producer: %s", err.Error())
		}
	}()

	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	fmt.Println("HERE1")

	service := service.NewService(producer)
	controller := controller.NewController(service)

	app := fiber.New()

	routes.VideoRoute(app, controller.VideoController)
	routes.AuthorRotute(app, controller.AuthorController)

	err = app.Listen(":8000")
	if err != nil {
		return
	}
}
