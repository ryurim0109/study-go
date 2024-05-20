package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	todoRouter "github.com/ryurim0109/study-go/router"
)





func main() {
	app := fiber.New()
	micro := fiber.New()

	app.Mount("/api", micro)
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://121.130.175.102:8000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowCredentials: true,
	}))

	todoRouter.SetupRoutes(micro)



	log.Fatal(app.Listen(":8004"))
}
	// app.Get("/api/*", func(c *fiber.Ctx) error {
  //       msg := fmt.Sprintf("✋ %s", c.Params("*"))
  //       return c.SendString(msg) // => ✋ register
  //   })



	
