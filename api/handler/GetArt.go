package handler

import "github.com/gofiber/fiber/v2"

func GetArt(ctx *fiber.Ctx) error {

	return ctx.SendString("Fiber Is Working")
}
