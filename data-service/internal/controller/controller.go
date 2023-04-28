package controller

import (
	"data-service/internal/service"
	"github.com/gofiber/fiber/v2"
)

type VideoController interface {
	Create(c *fiber.Ctx) error
	GetVideo(c *fiber.Ctx) error
	GetAllVideos(c *fiber.Ctx) error
	GetTop10Videos(c *fiber.Ctx) error
	GetTagsRate(c *fiber.Ctx) error
}

type AuthorController interface {
	Create(c *fiber.Ctx) error
	GetAuthor(c *fiber.Ctx) error
	GetAllAuthors(c *fiber.Ctx) error
	GetTop10Authors(c *fiber.Ctx) error
}

type Controller struct {
	VideoController
	AuthorController
}

func NewController(service *service.Service) *Controller {
	return &Controller{
		VideoController:  NewVideoControllerImpl(service.VideoService),
		AuthorController: NewAuthorControllerImpl(service.AuthorService),
	}
}
