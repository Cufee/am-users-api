package driver

import (
	"fmt"

	"aftermath.link/repo/logs"
)

func (d *Driver) Insert(database, collection string, payload ...interface{}) (*QueryResult, error) {
	err := d.Verify()
	if err != nil {
		return nil, logs.Wrap(err, "driver.Verify failed")
	}

	o, err := d.newOperation(database, collection)
	if err != nil {
		return nil, logs.Wrap(err, "newDriver failed")
	}
	defer o.finish()

	if len(payload) == 0 {
		return nil, fmt.Errorf("payload is nil")
	}

	r, err := o.insertManyDocuments(payload)
	if err != nil {
		return nil, err
	}
	var result QueryResult
	result.Inserted = r.Inserted
	return &result, nil
}
