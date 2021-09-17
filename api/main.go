package main

import (
	"artistically/api/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	routes.SpreadRoutes(app)

	log.Fatal(app.Listen(":3000"))

}
