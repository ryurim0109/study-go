package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	todo "github.com/ryurim0109/study-go/cmd/todo"
)



func status(c *fiber.Ctx) error {
	return c.SendString("Server is running! Send your request")
}



func setupRoutes(app *fiber.App) {

	app.Get("/", status)

	app.Get("/api/todo", todo.GetAllTodoList)
}

func main() {
	app := fiber.New()

	setupRoutes(app)

	// app.Get("/api/*", func(c *fiber.Ctx) error {
  //       msg := fmt.Sprintf("✋ %s", c.Params("*"))
  //       return c.SendString(msg) // => ✋ register
  //   })



	log.Fatal(app.Listen(":8004"))
}