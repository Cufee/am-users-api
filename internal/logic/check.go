package logic

import (
	"github.com/byvko-dev/am-types/users/v2"
	"github.com/byvko-dev/am-users-api/internal/core/database"
)

func CheckUserByExternalID(externalID interface{}, service users.ExternalService) (users.UserCheck, error) {
	profile, err := database.GetUserProfileByExternalID(externalID, service)
	if err != nil {
		return users.UserCheck{}, err
	}

	ban, err := database.GetUserBan(profile.ID)
	if err != nil {
		return users.UserCheck{}, err
	}

	connections, err := database.GetUserConnections(profile.ID)
	if err != nil {
		return users.UserCheck{}, err
	}

	return profile.ToUserCheck(ban, connections), nil
}

func CheckUserByID(id string) (users.UserCheck, error) {
	profile, err := database.GetUserProfileByID(id)
	if err != nil {
		return users.UserCheck{}, err
	}

	ban, err := database.GetUserBan(profile.ID)
	if err != nil {
		return users.UserCheck{}, err
	}

	connections, err := database.GetUserConnections(profile.ID)
	if err != nil {
		return users.UserCheck{}, err
	}

	return profile.ToUserCheck(ban, connections), nil
}
