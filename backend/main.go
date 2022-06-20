package main

import (
	"backend/configs"
	"backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	configs.ConnectDB()

	routes.MovieRoute(app)
	app.Listen(":" + configs.ApiPort())
}
