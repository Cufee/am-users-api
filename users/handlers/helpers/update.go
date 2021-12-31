package helpers

import (
	"aftermath.link/repo/logs"

	databaseusers "aftermath.link/repo/am-users-api/database/logic/users"
	"aftermath.link/repo/am-users-api/users"
)

func UpdateUserByDiscordID(discordID string, update users.InternalUser) error {
	filter := make(map[string]interface{})
	filter["_id"] = discordID
	err := databaseusers.UpdateDiscordUserWithFilter(filter, update)
	if err != nil {
		return logs.Wrap(err, "databaseusers.UpdateDiscordUserWithFilter failed")
	}
	return nil
}
