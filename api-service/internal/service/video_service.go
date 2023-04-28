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

type VideoServiceImpl struct {
	messageProducer sarama.SyncProducer
}

func (s VideoServiceImpl) Create(ctx context.Context, video entity.Video) error {
	message := dto.Message{
		RoutingKey: "video",
		Data:       video,
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

func (s VideoServiceImpl) GetAll(ctx context.Context) ([]map[string]interface{}, error) {

	host, exist := os.LookupEnv("DATA_SERVICE_HOST")
	if !exist {
		host = "http://localhost:8001/api/"
	}

	a := client.GetRequest(host + "videos")

	_, body, _ := a.Bytes()

	var data []map[string]interface{}

	err := json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s VideoServiceImpl) GetTop10Videos(ctx context.Context) ([]map[string]interface{}, error) {

	host, exist := os.LookupEnv("DATA_SERVICE_HOST")
	if !exist {
		host = "http://localhost:8001/api/"
	}

	a := client.GetRequest(host + "videos/top10")

	_, body, _ := a.Bytes()

	var data []map[string]interface{}

	err := json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s VideoServiceImpl) GetTagsRate(ctx context.Context) ([]map[string]interface{}, error) {

	host, exist := os.LookupEnv("DATA_SERVICE_HOST")
	if !exist {
		host = "http://localhost:8001/api/"
	}

	a := client.GetRequest(host + "videos/tags_rate")

	_, body, _ := a.Bytes()

	var data []map[string]interface{}

	err := json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func NewVideoServiceImpl(messageProducer sarama.SyncProducer) *VideoServiceImpl {
	return &VideoServiceImpl{messageProducer: messageProducer}
}
