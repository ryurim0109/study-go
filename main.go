package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	todoRouter "github.com/ryurim0109/study-go/router"
	logScheduler "github.com/ryurim0109/study-go/scheduler"
)

func main() {
	// Fiber 설정
	app := fiber.New()
	micro := fiber.New()

	app.Use(logger.New(logger.Config{
		Format:     "[${ip}]:${port} ${status} - ${method} ${path} - ${time} |${latency}\n",
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

	logScheduler.SecondCron()

	// Fiber 서버 시작
	log.Fatal(app.Listen(":8004"))
}
