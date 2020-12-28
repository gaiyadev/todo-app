package main

import (
	"fiber-todo/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
)


func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(logger.New())
	app.Use(cors.New())

	app.Static("/", "./public")
	//app.Use(limiter.New())

	//Adding api/todo prefix to the endpoints

	//Landing page
	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	api := app.Group("/api/v1")
	api.Get("/", handler.AllTodo)

	api.Post("/add", handler.CreateTodo)

	api.Get("/:id", handler.GetOne)
	api.Delete("/:id", handler.DeleteTodo)
//api.Patch("/:id", handler.UpdateTodo)

		//Error handling
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Sorry. something went  wrong!")
	})


	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}