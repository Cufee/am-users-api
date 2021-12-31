package banrecords

import (
	"aftermath.link/repo/am-users-api/database"
	"aftermath.link/repo/am-users-api/database/errors"
	"aftermath.link/repo/logs"
	mongodriver "aftermath.link/repo/mongo-driver"
)

const banRecordsCollectionName = "banrecords"

// FindBanRecordWithFilter finds a ban record by the given filter.
// If the ban record is not found, no error is returned.
// returned array is guaranteed to have at least one element.
func FindBanRecordWithFilter(filter, sort map[string]interface{}, limit int, target interface{}) error {
	var request mongodriver.FindRequest
	request.Query = filter
	request.Limit = limit
	request.Sort = sort

	response, err := database.Driver.FindInCollection(database.Name, banRecordsCollectionName, request)
	if err != nil {
		return logs.Wrap(err, "mongodriver.FindInCollection failed")
	}

	if response.Count.Matched == 0 {
		return nil
	}

	err = database.InterfaceToStruct(response.Documents, target)
	if err != nil {
		return logs.Wrap(err, "mongodriver.InterfaceToStruct failed")
	}

	return nil
}

// Create a new user in the mongodriver.
func AddNewBanRecord(payload interface{}) error {
	var request mongodriver.InsertRequest
	request.Documents = []interface{}{payload}

	response, err := database.Driver.InsertToCollection(database.Name, banRecordsCollectionName, request)
	if err != nil {
		return logs.Wrap(err, "mongodriver.InsertToCollection failed")
	}

	if response.Count.Inserted != 1 {
		return errors.ErrOperationCountMismatch
	}

	return nil
}
