package driver

import (
	"fmt"

	"aftermath.link/repo/logs"
)

func (d *Driver) Delete(database, collection string, filter map[string]interface{}) (*QueryResult, error) {
	err := d.Verify()
	if err != nil {
		return nil, logs.Wrap(err, "driver.Verify failed")
	}

	o, err := d.newOperation(database, collection)
	if err != nil {
		return nil, logs.Wrap(err, "newDriver failed")
	}
	defer o.finish()

	if filter == nil {
		return nil, fmt.Errorf("filter is nil")
	}

	r, err := o.deleteManyDocuments(filter)
	if err != nil {
		return nil, err
	}
	var result QueryResult
	result.Deleted = r.Deleted
	return &result, nil
}
