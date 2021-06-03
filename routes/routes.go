package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/y-yagi/go-api-template/handlers"
)

func New() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format:     "${cyan}[${ip}] ${blue}${status} ${blue}[${method}] ${white}${path} at ${white}${time}\n",
		TimeFormat: "2006-01-02 15:04:05 +09:00",
	}))

	api := app.Group("/api")
	api.Get("/books", handlers.GetBooks)
	api.Get("/books/:id", handlers.GetBook)
	api.Post("/books", handlers.CreateBook)
	api.Delete("/books/:id", handlers.DeleteBook)

	return app
}
