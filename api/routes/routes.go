package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SpreadRoutes(app *fiber.App) {

	app.Get("/", HomeRoute)
}
