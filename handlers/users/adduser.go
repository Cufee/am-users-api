package users

import (
	"aftermath.link/repo/am-users-api/handlers"
	userhandlers "aftermath.link/repo/am-users-api/users/handlers"
	"aftermath.link/repo/logs"
	"github.com/gofiber/fiber/v2"
)

type NewUserPostRequest struct {
	ID     string `json:"id"`
	Locale string `json:"locale"`
}

func NewUserHandler(ctx *fiber.Ctx) error {
	var response handlers.ResponseJSON

	var user NewUserPostRequest
	if err := ctx.BodyParser(&user); err != nil {
		logs.Warning("Failed to parse body: %v", err)
		response.Error = "Failed to parse body"
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	if err := userhandlers.CreateNewSimpleDiscordUserRecord(user.ID, user.Locale); err != nil {
		logs.Warning("Failed to create new user record: %v", err)
		response.Error = "Failed to create new user record"
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	return ctx.Status(fiber.StatusCreated).JSON(response) // Response is empty here
}
