package users

import (
	database "aftermath.link/repo/am-users-api/database"
	"aftermath.link/repo/am-users-api/database/errors"
	"aftermath.link/repo/logs"
	mongodriver "aftermath.link/repo/mongo-driver"
)

const (
	usersCollectionName = "users"
)

// FindDiscordUsersWithFilter finds a user by the given filter.
// If the user is not found, an error is returned.
// returned array is guaranteed to have at least one element.
func FindDiscordUsersWithFilter(filter, sort map[string]interface{}, limit int, target interface{}) error {
	var request mongodriver.FindRequest
	request.Query = filter
	request.Limit = limit
	request.Sort = sort

	response, err := database.Driver.FindInCollection(database.Name, usersCollectionName, request) // This is marshalled into json before being sent to the database
	if err != nil {
		return logs.Wrap(err, "database.FindInCollection failed")
	}

	if response.Count.Matched == 0 {
		return errors.ErrNoDocumentsFound
	}

	err = database.InterfaceToStruct(response.Documents, target)
	if err != nil {
		return logs.Wrap(err, "database.InterfaceToStruct failed")
	}

	return nil
}

// Create a new user in the database.
func AddDiscordUserRecord(payload ...interface{}) error {
	if len(payload) == 0 {
		return errors.ErrInvalidPayload
	}

	var request mongodriver.InsertRequest
	request.Documents = payload

	response, err := database.Driver.InsertToCollection(database.Name, usersCollectionName, request) // This is marshalled into json before being sent to the database
	if err != nil {
		return logs.Wrap(err, "database.InsertToCollection failed")
	}
	if response.Count.Inserted != 1 {
		return errors.ErrOperationCountMismatch
	}

	return nil
}

// UpdateDiscordUserRecord updates a user in the database using the given filter to find the user.
// Returns an error if the user is not found
func UpdateDiscordUserWithFilter(filter map[string]interface{}, update interface{}) error {
	var request mongodriver.UpdateRequest
	request.Query = filter
	request.Update = update

	response, err := database.Driver.UpdateInCollection(database.Name, usersCollectionName, request)
	if err != nil {
		return logs.Wrap(err, "database.UpdateInCollection failed")
	}

	if response.Count.Matched != 1 {
		return errors.ErrOperationCountMismatch
	}

	return nil
}
