package routes

import (
	"data-service/internal/controller"
	"github.com/gofiber/fiber/v2"
)

func VideoRoute(app *fiber.App, controller controller.VideoController) {
	app.Get("/api/video/:id?", controller.GetVideo)
	app.Get("/api/videos/", controller.GetAllVideos)
	app.Get("/api/videos/top10", controller.GetTop10Videos)
	app.Get("/api/videos/tags_rate", controller.GetTagsRate)
	app.Post("/api/video", controller.Create)
}
