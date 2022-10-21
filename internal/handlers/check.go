package handlers

import (
	api "github.com/byvko-dev/am-types/api/generic/v1"
	"github.com/byvko-dev/am-types/users/v2"
	er "github.com/byvko-dev/am-users-api/errors"
	"github.com/byvko-dev/am-users-api/internal/logic"
	"github.com/gofiber/fiber/v2"
)

func CheckUserByIDHandler(c *fiber.Ctx) error {
	var response api.ResponseWithError
	userId := c.Params("id")
	if userId == "" {
		response.Error.Message = er.ErrUserNotFound.Error()
		response.Error.Context = "failed to parse user id"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	check, err := logic.CheckUserByID(userId)
	if err != nil {
		response.Error.Message = err.Error()
		response.Error.Context = "failed to check user by id"
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = check
	return c.JSON(response)
}

func CheckUserByDiscordDHandler(c *fiber.Ctx) error {
	var response api.ResponseWithError
	id := c.Params("id")
	if id == "" {
		response.Error.Message = er.ErrUserNotFound.Error()
		response.Error.Context = "failed to parse user id"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	check, err := logic.CheckUserByExternalID(id, users.ExternalServiceDiscord.Name)
	if err != nil {
		response.Error.Message = err.Error()
		response.Error.Context = "failed to check user by discord id"
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = check
	return c.JSON(response)
}

func CheckUserByWargamingIDHandler(c *fiber.Ctx) error {
	var response api.ResponseWithError
	id := c.Params("id")
	if id == "" {
		response.Error.Message = er.ErrUserNotFound.Error()
		response.Error.Context = "failed to parse user id"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	check, err := logic.CheckUserByExternalID(id, users.ExternalServiceWargaming.Name)
	if err != nil {
		response.Error.Message = err.Error()
		response.Error.Context = "failed to check user by wargaming id"
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = check
	return c.JSON(response)
}
