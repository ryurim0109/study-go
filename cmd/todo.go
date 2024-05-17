package CmdTodo

import (
	"github.com/gofiber/fiber/v2"
)

func GetAllTodoList(c *fiber.Ctx) error {
	return c.SendString("allTodoList")
}