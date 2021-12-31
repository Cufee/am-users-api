package handlers

import (
	"fmt"

	"aftermath.link/repo/am-users-api/users"
	"aftermath.link/repo/am-users-api/users/customization"
	"aftermath.link/repo/am-users-api/users/handlers/helpers"
	"aftermath.link/repo/logs"
)

func FindDiscordUserByID(id string) (*users.InternalUser, error) {
	u, err := helpers.FindUserFromDiscordID(id)
	if err != nil {
		return nil, logs.Wrap(err, "helpers.FindUserFromDiscordID failed")
	}
	return u, nil
}

func FindDiscordUserByVerifiedPlayerID(playerId string) (*users.InternalUser, error) {
	u, err := helpers.FindUserByFieldValues(2, helpers.KeyValue{K: "player.verifiedId", V: playerId})
	if err != nil {
		return nil, logs.Wrap(err, "helpers.FindUserByFieldValues failed")
	}

	if len(u) > 1 {
		return nil, fmt.Errorf("found more than one user with playerId %s", playerId)
	}

	return &u[0], nil
}

func FindDiscordUsersByBackground(backgroundImageURL string) (*[]users.InternalUser, error) {
	key := fmt.Sprintf("customizations.%s", customization.BackgroundOptionKey)
	u, err := helpers.FindUserByFieldValues(0, helpers.KeyValue{K: key, V: backgroundImageURL})
	if err != nil {
		return nil, logs.Wrap(err, "databaseusers.FindDiscordUsersWithFilter failed")
	}
	return &u, nil
}
