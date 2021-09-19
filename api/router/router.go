package router

import (
	"artistically/api/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Setup(app *fiber.App) {
	api := app.Group("/v1", logger.New())

	api.Get("/", handler.Index)

	art := api.Group("/art")

	art.Get("/:id", handler.GetArtByID)
	art.Post("/art", handler.PostArt)
	art.Get("/arts/:id", handler.GetUserArtsById)
}
