package main

import (
	"backend/configs"
	"backend/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	configs.ConnectDB()

	routes.MovieRoute(app)
	app.Listen(":" + configs.ApiPort())
}
