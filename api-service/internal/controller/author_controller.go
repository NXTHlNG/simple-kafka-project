package controller

import (
	"api-service/internal/entity"
	"api-service/internal/response"
	"api-service/internal/service"
	"context"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"time"
)

type AuthorControllerImpl struct {
	service service.AuthorService
}

func NewAuthorControllerImpl(service service.AuthorService) *AuthorControllerImpl {
	return &AuthorControllerImpl{service: service}
}

func (v *AuthorControllerImpl) Create(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var author entity.Author
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&author); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	////use the validator library to validate required fields
	//if validationErr := validate.Struct(&task); validationErr != nil {
	//	return c.Status(http.StatusBadRequest).JSON(responses.TaskResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	//}

	err := v.service.Create(ctx, author)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.VideoResponse{Status: http.StatusInternalServerError, Message: "error"})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "success"})
}

func (v *AuthorControllerImpl) GetAllAuthors(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	results, err := v.service.GetAll(ctx)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(results)
}

func (v *AuthorControllerImpl) GetTop10Authors(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	results, err := v.service.GetTop10Authors(ctx)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(results)
}
