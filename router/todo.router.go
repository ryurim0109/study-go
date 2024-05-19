package todoRouter

import (
	"github.com/gofiber/fiber/v2"
	todo "github.com/ryurim0109/study-go/cmd/todo"
)

func SetupRoutes(micro *fiber.App) {


	// app.Get("/api/todo", todo.GetAllTodoList)

		micro.Route("/todo", func(router fiber.Router) {
		router.Get("/", todo.Get)
		router.Post("/", todo.Create)
	})
}