package service

import (
	"api-service/internal/client"
	"api-service/internal/dto"
	"api-service/internal/entity"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
)

type AuthorServiceImpl struct {
	messageProducer sarama.SyncProducer
}

func (s AuthorServiceImpl) Create(ctx context.Context, author entity.Author) error {
	message := dto.Message{
		RoutingKey: "author",
		Data:       author,
	}

	jsonMessage, err := json.Marshal(message)
	if err != nil {
		log.Fatal(err)
	}

	msg := &sarama.ProducerMessage{
		Topic: "videos",
		Value: sarama.StringEncoder(jsonMessage),
	}
	partition, offset, err := s.messageProducer.SendMessage(msg)
	if err != nil {
		return err
	}
	fmt.Printf("Message sent to partition %d at offset %d\n", partition, offset)
	return nil
}

func (s AuthorServiceImpl) GetAll(ctx context.Context) ([]map[string]interface{}, error) {

	host, exist := os.LookupEnv("DATA_SERVICE_HOST")
	if !exist {
		host = "http://localhost:8001/api/"
	}

	a := client.GetRequest(host + "authors")

	_, body, _ := a.Bytes()

	var data []map[string]interface{}

	err := json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s AuthorServiceImpl) GetTop10Authors(ctx context.Context) ([]map[string]interface{}, error) {

	host, exist := os.LookupEnv("DATA_SERVICE_HOST")
	if !exist {
		host = "http://localhost:8001/api/"
	}

	a := client.GetRequest(host + "authors/top10")

	_, body, _ := a.Bytes()

	var data []map[string]interface{}

	err := json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func NewAuthorServiceImpl(messageProducer sarama.SyncProducer) *AuthorServiceImpl {
	return &AuthorServiceImpl{messageProducer: messageProducer}
}
