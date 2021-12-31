package users

import (
	"aftermath.link/repo/am-users-api/handlers"
	userhandlers "aftermath.link/repo/am-users-api/users/handlers"
	"aftermath.link/repo/logs"
	"github.com/gofiber/fiber/v2"
)

func UpdatePlayerIFByDiscordIDHandler(ctx *fiber.Ctx) error {
	var response handlers.ResponseJSON

	id := ctx.Params("id")
	if id == "" {
		logs.Warning("No ID provided")
		response.Error = "No ID provided"
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	playerId := ctx.Params("playerId")
	if playerId == "" {
		logs.Warning("No player ID provided")
		response.Error = "No player ID provided"
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	u, err := userhandlers.SetPlayerIDByDiscordID(id, playerId)
	if err != nil {
		logs.Warning("Failed to find user: %v", err)
		response.Error = "Failed to find user"
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = getProfileWithBanCheck(u)
	return ctx.JSON(response)
}
