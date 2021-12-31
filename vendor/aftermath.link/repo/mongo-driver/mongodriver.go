package mongodriver

import (
	"fmt"

	"aftermath.link/repo/logs"
	"aftermath.link/repo/mongo-driver/driver"
)

func (s *SimpleDriver) Setup(c SimpleConfig) {
	var d driver.Driver
	d.SetURI(c.URI)
	d.SetUser(c.User)
	d.SetPass(c.Password)
	s.Advanced = &d
}

func (s *SimpleDriver) preflightCheck() error {
	if s.Advanced == nil {
		return fmt.Errorf("globalDriver is nil")
	}
	return s.Advanced.Verify()
}

func (s *SimpleDriver) FindInCollection(databaseName, collection string, payload FindRequest) (*FindResponse, error) {
	err := s.preflightCheck()
	if err != nil {
		return nil, logs.Wrap(err, "preflightCheck failed")
	}

	var response FindResponse
	result, err := s.Advanced.FindMany(databaseName, collection, payload.Query, payload.Sort, payload.Limit, &response.Documents)
	if err != nil {
		return nil, logs.Wrap(err, "d.FindMany failed")
	}
	response.Count = Count(*result)
	return &response, nil
}

func (s *SimpleDriver) InsertToCollection(databaseName, collection string, payload InsertRequest) (*InsertResponse, error) {
	err := s.preflightCheck()
	if err != nil {
		return nil, logs.Wrap(err, "preflightCheck failed")
	}

	var response InsertResponse
	result, err := s.Advanced.Insert(databaseName, collection, payload.Documents...)
	if err != nil {
		return nil, logs.Wrap(err, "d.FindMany failed")
	}
	response.Count = Count(*result)
	return &response, nil
}

func (s *SimpleDriver) UpdateInCollection(databaseName, collection string, payload UpdateRequest) (*UpdateResponse, error) {
	err := s.preflightCheck()
	if err != nil {
		return nil, logs.Wrap(err, "preflightCheck failed")
	}

	var response UpdateResponse
	result, err := s.Advanced.UpdateOneWithUpsert(databaseName, collection, payload.Query, payload.Upsert, payload.Update)
	if err != nil {
		return nil, logs.Wrap(err, "d.FindMany failed")
	}
	response.Count = Count(*result)
	return &response, nil
}

func (s *SimpleDriver) DeleteInCollection(databaseName, collection string, payload DeleteRequest) (*DeleteResponse, error) {
	err := s.preflightCheck()
	if err != nil {
		return nil, logs.Wrap(err, "preflightCheck failed")
	}

	var response DeleteResponse
	result, err := s.Advanced.Delete(databaseName, collection, payload.Query)
	if err != nil {
		return nil, logs.Wrap(err, "d.FindMany failed")
	}
	response.Count = Count(*result)
	return &response, nil
}
