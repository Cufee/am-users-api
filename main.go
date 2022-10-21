package main

import (
	"github.com/byvko-dev/am-core/helpers/env"
	"github.com/byvko-dev/am-users-api/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	logs "github.com/sirupsen/logrus"
)

func main() {
	// Setup a server
	app := fiber.New()
	app.Use(logger.New())

	// Users
	v1 := app.Group("/v1")

	v1.Post("/", handlers.CreateUserHandler)                                // Create a user
	v1.Put("/:id/connections", handlers.UpdateUserConnectionsHandler)       // Update user connections
	v1.Put("/:id/customizations", handlers.UpdateUserCustomizationsHandler) // Update user customizations

	check := v1.Group("/check")
	check.Get("/:id", handlers.CheckUserByIDHandler)
	check.Get("/discord/:id", handlers.CheckUserByDiscordDHandler)
	check.Get("/wargaming/:id", handlers.CheckUserByWargamingIDHandler)

	logs.Fatal("Failed to start a server: %v", app.Listen(":"+env.MustGetString("PORT")))
}
