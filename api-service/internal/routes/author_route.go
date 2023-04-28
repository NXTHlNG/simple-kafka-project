package routes

import (
	"api-service/internal/controller"
	"github.com/gofiber/fiber/v2"
)

func AuthorRotute(app *fiber.App, controller controller.AuthorController) {
	app.Get("/api/authors/", controller.GetAllAuthors)
	app.Get("/api/authors/top10", controller.GetTop10Authors)
	app.Post("/api/author", controller.Create)
}
