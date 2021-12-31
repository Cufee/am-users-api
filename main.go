package main

import (
	"aftermath.link/repo/am-users-api/database"
	usershandlers "aftermath.link/repo/am-users-api/handlers/users"
	"aftermath.link/repo/logs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Setup database driver
	err := database.InitDriver()
	if err != nil {
		logs.Fatal("InitDriver failed: %v", err)
	}

	// Setup a server
	app := fiber.New()
	app.Use(logger.New())

	// Users
	users := app.Group("/users")

	// From background image - this is a low priority endpoint as it will be used for bans
	manageusers := users.Group("/manage")                 // Admin endpoint to manage users and do some odd lookups
	manageusers.Get("/background/:url", dummyHandlerfunc) // Find many users from background URL

	// From Discord ID
	discordusers := users.Group("/discord/:id")
	discordusers.Get("/", usershandlers.FindUserByDiscordIDHandler) // Find user from discord ID
	discordusers.Patch("/set-player/:playerId", usershandlers.UpdatePlayerIFByDiscordIDHandler)
	discordusers.Delete("/customizations", dummyHandlerfunc) // Delete customization
	discordusers.Patch("/customizations", dummyHandlerfunc)  // Update customization
	discordusers.Put("/customizations", dummyHandlerfunc)    //

	// From Player ID
	playerusers := users.Group("/player")
	playerusers.Get("/:id", usershandlers.FindUserByPlayerIDHandler) // Find user from player ID

	// Post
	users.Post("/discord", usershandlers.NewUserHandler) // Create a new User from ID and Locale

	logs.Fatal("Failed to start a server: %v", app.Listen(":3001"))
}

func dummyHandlerfunc(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}
