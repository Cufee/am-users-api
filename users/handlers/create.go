package handlers

import (
	"fmt"

	databaseusers "aftermath.link/repo/am-users-api/database/logic/users"
	"aftermath.link/repo/am-users-api/users"
	"aftermath.link/repo/logs"
)

func CreateNewDiscordUserRecord(user *users.InternalUser) error {
	ok := user.IsRecordComplete()
	if !ok {
		return fmt.Errorf("user record is not complete")
	}
	err := databaseusers.AddDiscordUserRecord(*user)
	if err != nil {
		return logs.Wrap(err, "databaseusers.AddDiscordUserRecord failed")
	}

	return nil
}

func CreateNewSimpleDiscordUserRecord(id, locale string) error {
	var user users.InternalUser
	user.ID = id
	user.Locale = locale

	err := CreateNewDiscordUserRecord(&user)
	if err != nil {
		return logs.Wrap(err, "CreateNewDiscordUserRecord failed")
	}

	return nil
}
