package helpers

import (
	databaseusers "aftermath.link/repo/am-users-api/database/logic/users"
	"aftermath.link/repo/am-users-api/users"
	"aftermath.link/repo/logs"
)

type KeyValue struct {
	K string
	V interface{}
}

func FindUserFromDiscordID(discordID string) (*users.InternalUser, error) {
	u, err := FindUserByFieldValues(1, KeyValue{K: "_id", V: discordID})
	if err != nil {
		return nil, logs.Wrap(err, "FindUserByFieldValues failed")
	}
	return &u[0], nil
}

func FindUserByFieldValues(limit int, pairs ...KeyValue) ([]users.InternalUser, error) {
	var u []users.InternalUser
	filter := make(map[string]interface{})
	for _, pair := range pairs {
		filter[pair.K] = pair.V
	}

	err := databaseusers.FindDiscordUsersWithFilter(filter, nil, limit, &u) // Returns at least 1 element or err
	if err != nil {
		return nil, logs.Wrap(err, "databaseusers.FindDiscordUsersWithFilter failed")
	}

	return u, nil
}
