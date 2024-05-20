package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	todoRouter "github.com/ryurim0109/study-go/router"
)

func main() {
	app := fiber.New()
	micro := fiber.New()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			main()
		}
	}()

	app.Use(logger.New(logger.Config{
		Format:     "[${ip}]:${port} ${status} - ${method} ${path} - ${time} \n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Asia/Seoul",
	}))

	app.Mount("/api", micro)

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://121.130.175.102:8000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowCredentials: true,
	}))

	todoRouter.SetupRoutes(micro)

	log.Fatal(app.Listen(":8004"))
}
