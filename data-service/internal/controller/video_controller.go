package controller

import (
	"context"
	"data-service/internal/entity"
	"data-service/internal/service"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

type VideoControllerImpl struct {
	service service.VideoService
}

func NewVideoControllerImpl(service service.VideoService) *VideoControllerImpl {
	return &VideoControllerImpl{service: service}
}

func (v *VideoControllerImpl) Create(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var video entity.Video
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&video); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	////use the validator library to validate required fields
	//if validationErr := validate.Struct(&task); validationErr != nil {
	//	return c.Status(http.StatusBadRequest).JSON(responses.TaskResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	//}

	insertedId, err := v.service.Create(ctx, video)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(insertedId)
}

func (v *VideoControllerImpl) GetVideo(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	taskId := c.Params("id")
	defer cancel()

	id, _ := primitive.ObjectIDFromHex(taskId)

	res, err := v.service.Get(ctx, id)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(res)
}

func (v *VideoControllerImpl) GetAllVideos(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	results, err := v.service.GetAll(ctx)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(results)
}

func (v *VideoControllerImpl) GetTop10Videos(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	results, err := v.service.GetTop10(ctx)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(results)
}

func (v *VideoControllerImpl) GetTagsRate(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	results, err := v.service.GetTagsRate(ctx)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(results)
}
