package handlers

import (
	"github.com/byvko-dev/am-core/helpers/slices"
	api "github.com/byvko-dev/am-types/api/generic/v1"
	"github.com/byvko-dev/am-types/users/v2"
	er "github.com/byvko-dev/am-users-api/errors"
	"github.com/byvko-dev/am-users-api/internal/logic"
	"github.com/gofiber/fiber/v2"
)

func CreateUserHandler(c *fiber.Ctx) error {
	var response api.ResponseWithError
	var payload struct {
		Connections []users.ExternalProfileID `json:"connections"`
		Locale      string                    `json:"locale"`
	}
	if err := c.BodyParser(&payload); err != nil {
		response.Error.Message = er.ErrInvalidPayload.Error()
		response.Error.Context = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	if len(payload.Connections) == 0 {
		response.Error.Message = er.ErrInvalidPayload.Error()
		response.Error.Context = "failed to create user"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	for _, connection := range payload.Connections {
		profile, _ := logic.CheckUserByExternalID(connection.ExternalID, connection.Service)
		if profile.ID != "" {
			response.Error.Message = er.ErrConnectionAlreadyExists.Error()
			response.Error.Context = "failed to update user connections"
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}
	}

	profileID, err := logic.NewUser(payload.Connections, payload.Locale)
	if err != nil {
		response.Error.Message = err.Error()
		response.Error.Context = "failed to create new user"
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = profileID
	return c.Status(fiber.StatusCreated).JSON(response)
}

func UpdateUserConnectionsHandler(c *fiber.Ctx) error {
	var response api.ResponseWithError
	var payload []users.ExternalProfileID
	if err := c.BodyParser(&payload); err != nil {
		response.Error.Message = er.ErrInvalidPayload.Error()
		response.Error.Context = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

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

	for _, connection := range payload {
		connectionIndex := slices.Contains(users.ValidExternalServiceNames, connection.Service)
		if connectionIndex == -1 {
			response.Error.Message = er.ErrInvalidPayload.Error()
			response.Error.Context = "failed to update user connections"
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}

		if users.ValidExternalServices[connectionIndex].Unique {
			profile, _ := logic.CheckUserByExternalID(connection.ExternalID, connection.Service)
			if profile.ID != "" && profile.ID != check.ID {
				response.Error.Message = er.ErrConnectionAlreadyExists.Error()
				response.Error.Context = "failed to update user connections"
				return c.Status(fiber.StatusBadRequest).JSON(response)
			}
		}
		err := logic.UpdateUnverifiedConnection(userId, connection.Service, connection.ExternalID)
		if err != nil {
			response.Error.Message = err.Error()
			response.Error.Context = "failed to update user connections"
			return c.Status(fiber.StatusInternalServerError).JSON(response)
		}
	}

	return c.JSON(response)
}

func UpdateUserCustomizationsHandler(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}
