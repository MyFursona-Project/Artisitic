package main

import (
	"log"

	"artistically/api/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	api := fiber.New()
	api.Use(cors.New())
	api.Use(logger.New())
	router.Setup(api)
	log.Fatal(api.Listen(":3000"))
}
