package controllers

import (
	"backend/configs"
	"backend/models"
	"backend/responses"
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var movieCollection *mongo.Collection = configs.GetCollection(configs.DB, "movies")

func GetAllMovies(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var movies []models.Movie
	defer cancel()

	results, err := movieCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.MovieResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var dMovies models.Movie
		if err = results.Decode(&dMovies); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.MovieResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data: &fiber.Map{
					"data": err.Error(),
				},
			})
		}

		movies = append(movies, dMovies)
	}

	return c.Status(http.StatusOK).JSON(responses.MovieResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    &fiber.Map{"data": movies}},
	)
}

func CreateMovie(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var movie models.Movie
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&movie); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.MovieResponse{
			Status:  http.StatusBadRequest,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()},
		})
	}

	newMovie := models.Movie{
		Id:    primitive.NewObjectID(),
		Title: movie.Title,
		Genre: movie.Genre,
		Year:  movie.Year,
	}

	result, err := movieCollection.InsertOne(ctx, newMovie)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.MovieResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()},
		})
	}

	return c.Status(http.StatusCreated).JSON(responses.MovieResponse{
		Status:  http.StatusCreated,
		Message: "success",
		Data:    &fiber.Map{"data": result},
	})
}

func EditMovie(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	movieId := c.Params("movieId")
	var movie models.Movie
	defer cancel()

	objId, nil := primitive.ObjectIDFromHex(movieId)

	//validate the request body
	if err := c.BodyParser(&movie); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.MovieResponse{
			Status:  http.StatusBadRequest,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()},
		})
	}

	update := bson.M{
		"title": movie.Title,
		"genre": movie.Genre,
		"year":  movie.Year,
	}

	result, err := movieCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.MovieResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()},
		})
	}
	//get updated user details
	var updatedMovie models.Movie

	if result.MatchedCount == 1 {
		err := movieCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedMovie)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.MovieResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()},
			})
		}
	}

	return c.Status(http.StatusOK).JSON(responses.MovieResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    &fiber.Map{"data": updatedMovie},
	})
}

func GetMovie(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	movieId := c.Params("movieId")
	var movie models.Movie
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(movieId)

	err := movieCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&movie)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.MovieResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()},
		})
	}

	return c.Status(http.StatusOK).JSON(responses.MovieResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    &fiber.Map{"data": movie},
	})
}

func DeleteMovie(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	movieId := c.Params("movieId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(movieId)

	result, err := movieCollection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.MovieResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()},
		})
	}

	if result.DeletedCount < 1 {
		return c.Status(http.StatusNotFound).JSON(
			responses.MovieResponse{
				Status:  http.StatusNotFound,
				Message: "error",
				Data:    &fiber.Map{"data": "User with specified ID not found!"},
			},
		)
	}

	return c.Status(http.StatusOK).JSON(
		responses.MovieResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    &fiber.Map{"data": "User successfully deleted!"},
		},
	)
}
