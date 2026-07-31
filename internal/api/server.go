package api

import (
	"Go-Ecom-Aws/config"

	"github.com/gofiber/fiber/v2"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	app.Listen(config.ServerPort)
}
