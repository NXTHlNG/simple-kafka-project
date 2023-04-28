package controller

import (
	"api-service/internal/dto"
	"api-service/internal/entity"
	"api-service/internal/response"
	"api-service/internal/service"
	"context"
	"github.com/gofiber/fiber/v2"
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
	var video dto.VideoDto
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&video); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	////use the validator library to validate required fields
	//if validationErr := validate.Struct(&task); validationErr != nil {
	//	return c.Status(http.StatusBadRequest).JSON(responses.TaskResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	//}

	err := v.service.Create(ctx, entity.VideoFromDto(&video))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.VideoResponse{Status: http.StatusInternalServerError, Message: "error"})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "success"})
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

	results, err := v.service.GetTop10Videos(ctx)

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
