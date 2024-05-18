package todo

import (
	"github.com/gofiber/fiber/v2"
	database "github.com/ryurim0109/study-go/db/mariadb"
	
)

func GetAllTodoList(c *fiber.Ctx) error {
	db := database.DBConn
	db.Find(&edel_todo)
	return c.SendString("allTodoList")
}