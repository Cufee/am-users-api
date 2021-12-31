package driver

import (
	"fmt"

	"aftermath.link/repo/logs"
)

func (d *Driver) FindOne(database, collection string, filter, sort map[string]interface{}, target interface{}) (*QueryResult, error) {
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
	if sort == nil {
		sort = make(map[string]interface{})
	}

	r, err := o.getDocument(filter, sort, target)
	if err != nil {
		return nil, err
	}

	var result QueryResult
	result.Matched = r.Matched
	return &result, nil
}

func (d *Driver) FindMany(database, collection string, filter, sort map[string]interface{}, limit int, target interface{}) (*QueryResult, error) {
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
	if sort == nil {
		sort = make(map[string]interface{})
	}

	r, err := o.getManyDocuments(filter, sort, limit, target)
	if err != nil {
		return nil, err
	}

	var result QueryResult
	result.Matched = r.Matched
	return &result, nil
}
