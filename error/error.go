package errorHandler

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func SendJSONError(c *fiber.Ctx, message string) error {
	log.Fatalf("error: %v\n", message)
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": message})
}
