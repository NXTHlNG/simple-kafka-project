package routes

import (
	"data-service/internal/controller"
	"github.com/gofiber/fiber/v2"
)

func AuthorRoute(app *fiber.App, controller controller.AuthorController) {
	app.Get("/api/author/:id?", controller.GetAuthor)
	app.Get("/api/authors/", controller.GetAllAuthors)
	app.Get("/api/authors/top10", controller.GetTop10Authors)
	app.Post("/api/author", controller.Create)
}
