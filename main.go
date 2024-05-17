package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)



func status(c *fiber.Ctx) error {
	return c.SendString("Server is running! Send your request")
}

func setupRoutes(app *fiber.App) {

	app.Get("/", status)

	app.Get("/api/bookmark", todo.GetAllTodoList)
}

func main() {
	app := fiber.New()

	// app.Get("/api/*", func(c *fiber.Ctx) error {
  //       msg := fmt.Sprintf("✋ %s", c.Params("*"))
  //       return c.SendString(msg) // => ✋ register
  //   })



	log.Fatal(app.Listen(":8004"))
}