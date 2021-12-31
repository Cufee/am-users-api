package handlers

import (
	"fmt"
	"strconv"

	"aftermath.link/repo/am-users-api/users"
	"aftermath.link/repo/am-users-api/users/customization"
	"aftermath.link/repo/am-users-api/users/handlers/helpers"
	"aftermath.link/repo/am-users-api/users/players"
	"aftermath.link/repo/logs"
)

func SetBackgroundByDiscordID(id string, url string) (*users.InternalUser, error) {
	u, err := helpers.FindUserFromDiscordID(id)
	if err != nil {
		return nil, logs.Wrap(err, "helpers.FindUserFromDiscordID failed")
	}
	addBackgroundURL(u, url)

	// Check if the record is still valid
	ok := u.IsRecordComplete()
	if !ok {
		return nil, fmt.Errorf("user record is not complete")
	}

	err = helpers.UpdateUserByDiscordID(u.ID, *u)
	if err != nil {
		return nil, logs.Wrap(err, "dhelpers.UpdateUserByDiscordID failed")
	}
	return u, nil
}

func addBackgroundURL(u *users.InternalUser, url string) {
	background := customization.NewBackgroundOption(url)
	u.UniqueCustomizations[background.Key] = background
}

func SetPlayerIDByDiscordID(id, playerId string) (*users.InternalUser, error) {
	u, err := helpers.FindUserFromDiscordID(id)
	if err != nil {
		return nil, logs.Wrap(err, "helpers.FindUserFromDiscordID failed")
	}

	err = setPlayerID(u, playerId)
	if err != nil {
		return nil, logs.Wrap(err, "setPlayerID failed")
	}

	// Check if the record is still valid
	ok := u.IsRecordComplete()
	if !ok {
		return nil, fmt.Errorf("user record is not complete")
	}

	err = helpers.UpdateUserByDiscordID(u.ID, *u)
	if err != nil {
		return nil, logs.Wrap(err, "databaseusers.UpdateDiscordUserWithFilter failed")
	}
	return u, nil
}

func setPlayerID(u *users.InternalUser, playerId string) error {
	id, err := strconv.Atoi(playerId)
	if err != nil {
		return logs.Wrap(err, "strconv.Atoi failed")
	}

	if u.Player == nil {
		u.Player = &players.InternalProfile{}
	}
	u.Player.DefaultID = id

	return nil
}
