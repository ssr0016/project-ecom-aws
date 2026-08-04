package api

import (
	"Go-Ecom-Aws/config"
	"Go-Ecom-Aws/internal/api/rest"
	"Go-Ecom-Aws/internal/api/rest/handlers"
	"Go-Ecom-Aws/internal/domain"
	"log"

	"github.com/gofiber/fiber/v2"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("database connection error %v\n", err)
	}

	log.Println("database connected")

	// run migration
	db.AutoMigrate(&domain.User{})

	rh := &rest.RestHandler{App: app, DB: db}

	setupRoutes(rh)

	app.Listen(config.ServerPort)
}

func setupRoutes(rh *rest.RestHandler) {
	// user handler
	handlers.SetupUserRoute(rh)

	// transaction handler

	// catalog handler
}
