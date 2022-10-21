package logic

import (
	"github.com/byvko-dev/am-types/users/v2"
	"github.com/byvko-dev/am-users-api/internal/core/database"
)

func UpdateUserCustomizations(userId string, customizations users.Customizations) error {
	profile, err := database.GetUserProfileByID(userId)
	if err != nil {
		return err
	}

	profile.Customizations = customizations
	_, err = database.UpdateUserProfile(*profile)
	return err
}
