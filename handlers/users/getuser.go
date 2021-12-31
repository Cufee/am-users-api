package users

import (
	"aftermath.link/repo/am-users-api/handlers"
	"aftermath.link/repo/am-users-api/users"
	userhandlers "aftermath.link/repo/am-users-api/users/handlers"
	"aftermath.link/repo/logs"
	"github.com/gofiber/fiber/v2"
)

func FindUserByDiscordIDHandler(ctx *fiber.Ctx) error {
	var response handlers.ResponseJSON

	id := ctx.Params("id")
	if id == "" {
		logs.Warning("No ID provided")
		response.Error = "No ID provided"
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	user, err := userhandlers.FindDiscordUserByID(id)
	if err != nil {
		logs.Warning("Failed to find user: %v", err)
		response.Error = "Failed to find user"
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = getProfileWithBanCheck(user)
	return ctx.JSON(response)
}

func FindUserByPlayerIDHandler(ctx *fiber.Ctx) error {
	var response handlers.ResponseJSON

	id := ctx.Params("id")
	if id == "" {
		logs.Warning("No ID provided")
		response.Error = "No ID provided"
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	user, err := userhandlers.FindDiscordUserByVerifiedPlayerID(id)
	if err != nil {
		logs.Warning("Failed to find user: %v", err)
		response.Error = "Failed to find user"
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = getProfileWithBanCheck(user)
	return ctx.JSON(response)
}

func getProfileWithBanCheck(user *users.InternalUser) users.ExternalUser {
	profile := user.Export()
	profile.AddBanDetails()
	return profile
}
