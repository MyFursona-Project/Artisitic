package routes

import "github.com/gofiber/fiber/v2"

func HomeRoute(ctx *fiber.Ctx) error {
	ctx.SendString("Hello, and Welcome to the Artistically Api")
	return nil
}
