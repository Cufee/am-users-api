package users

import (
	"aftermath.link/repo/am-users-api/external/database"
	"aftermath.link/repo/am-users-api/external/database/errors"
	"aftermath.link/repo/logs"
	"github.com/mitchellh/mapstructure"
)

const (
	usersCollectionName = "users"
)

func FindUsersByDiscordID(id string, target interface{}) error {
	var request database.FindRequest
	request.Query = make(map[string]interface{})
	request.Query["discordId"] = id

	response, err := database.FindInCollection(usersCollectionName, request)
	if err != nil {
		return logs.Wrap(err, "database.FindInCollection failed")
	}

	if response.Count.Matched == 0 {
		return errors.ErrNoDocumentsFound
	}

	return mapstructure.Decode(response.Documents, &target)
}

func AddDiscordUserRecord(payload interface{}) error {
	var request database.InsertRequest
	request.Documents = make([]map[string]interface{}, 1)
	err := mapstructure.Decode(payload, &request.Documents[0])
	if err != nil {
		return logs.Wrap(err, "mapstructure.Decode failed")
	}

	response, err := database.InsertToCollection(usersCollectionName, request)
	if err != nil {
		return logs.Wrap(err, "database.FindInCollection failed")
	}

	if response.Count.Inserted != 1 {
		return errors.ErrOperationCountMismatch
	}

	return nil
}
