package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func MovieRoute(app *fiber.App) {
	app.Post("/movie", controllers.CreateMovie)
	app.Get("/movie/:movieId", controllers.GetMovie)
	app.Put("/movie/:movieId", controllers.EditMovie)
	app.Delete("/movie/:movieId", controllers.DeleteMovie)
	app.Get("/movies", controllers.GetAllMovies)
}
