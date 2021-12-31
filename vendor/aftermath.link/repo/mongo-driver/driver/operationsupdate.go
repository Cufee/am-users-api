package driver

import (
	"fmt"

	"aftermath.link/repo/logs"
)

func (d *Driver) UpdateOne(database, collection string, filter map[string]interface{}, payload interface{}) (*QueryResult, error) {
	return d.updateWithUpsert(database, collection, filter, false, payload)
}

func (d *Driver) UpsertOne(database, collection string, filter map[string]interface{}, payload interface{}) (*QueryResult, error) {
	return d.updateWithUpsert(database, collection, filter, true, payload)
}

func (d *Driver) UpdateOneWithUpsert(database, collection string, filter map[string]interface{}, upsert bool, payload interface{}) (*QueryResult, error) {
	return d.updateWithUpsert(database, collection, filter, upsert, payload)
}

func (d *Driver) updateWithUpsert(database, collection string, filter map[string]interface{}, upsert bool, payload interface{}) (*QueryResult, error) {
	err := d.Verify()
	if err != nil {
		return nil, logs.Wrap(err, "driver.Verify failed")
	}

	o, err := d.newOperation(database, collection)
	if err != nil {
		return nil, logs.Wrap(err, "newDriver failed")
	}
	defer o.finish()

	if payload == nil {
		return nil, fmt.Errorf("payload is nil")
	}
	if filter == nil {
		return nil, fmt.Errorf("filter is nil")
	}

	r, err := o.updateDocument(filter, upsert, payload)
	if err != nil {
		return nil, err
	}
	var result QueryResult
	result.Updated = r.Updated + r.Inserted
	result.Matched = r.Matched
	return &result, nil
}
