package todoRouter

import (
	"github.com/gofiber/fiber/v2"
	todo "github.com/ryurim0109/study-go/cmd/todo"
)

func setupRoutes(micro *fiber.App) {


	// app.Get("/api/todo", todo.GetAllTodoList)

		micro.Route("/todo", func(router fiber.Router) {
		router.Post("/", todo.Create)
	})
}