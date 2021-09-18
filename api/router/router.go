package router

import (
	"artistically/api/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Setup(app *fiber.App) {
	api := app.Group("/v1", logger.New())

	api.Get("/art/:id", handler.GetArt)
}
