package errorHandler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func SendJSONError(c *fiber.Ctx, message string) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": message})
		}
	}()

	panic(message)

}
