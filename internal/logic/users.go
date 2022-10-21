package logic

import (
	"github.com/byvko-dev/am-core/helpers/slices"
	"github.com/byvko-dev/am-core/helpers/strings"
	"github.com/byvko-dev/am-types/users/v2"
	"github.com/byvko-dev/am-users-api/internal/core/database"
)

func NewUser(connections []users.ExternalProfileID, locale string) (string, error) {
	var newProfile users.CompleteProfile
	newProfile.Features.SetDefaults()
	newProfile.Locale = strings.Or(locale, "en")

	id, err := database.CreateUserProfile(newProfile)
	if err != nil {
		return "", err
	}

	for _, connection := range connections {
		if slices.Contains(users.ValidExternalServiceNames, connection.Service) > -1 {
			err := UpdateUnverifiedConnection(id, connection.Service, connection.ExternalID)
			if err != nil {
				return "", err
			}
		}
	}

	return id, nil
}
