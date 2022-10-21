package logic

import (
	"github.com/byvko-dev/am-types/users/v2"
	"github.com/byvko-dev/am-users-api/internal/core/database"
)

func UpdateUnverifiedConnection(userId string, service users.ExternalService, externalId string) error {
	var connection users.ExternalConnection
	connection.UserID = userId
	connection.Service = service
	connection.ExternalID = externalId
	return database.UpdateUserConnection(connection)
}

func UpdateVerifiedConnection(connection users.ExternalConnection) error {
	return database.UpdateUserConnection(connection)
}
