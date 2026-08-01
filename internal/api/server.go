package api

import (
	"Go-Ecom-Aws/config"
	"Go-Ecom-Aws/internal/api/rest"
	"Go-Ecom-Aws/internal/api/rest/handlers"

	"github.com/gofiber/fiber/v2"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	rh := &rest.RestHandler{App: app}

	setupRoutes(rh)

	app.Listen(config.ServerPort)
}

func setupRoutes(rh *rest.RestHandler) {
	// user handler
	handlers.SetupUserRoute(rh)

	// transaction handler

	// catalog handler
}
