package driver

import (
	"fmt"

	"aftermath.link/repo/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (o *Operation) getDocument(filters map[string]interface{}, sort map[string]interface{}, target interface{}) (r OperationResult, err error) {
	// Set sort options
	opts := options.FindOne()
	sortInterface := mapToBSON(sort)

	if len(sortInterface) > 0 {
		opts.SetSort(sort)
	}

	// Compile query
	filter := mapToBSON(filters)

	// Find and decode
	err = o.collection.FindOne(o.ctx, filter, opts).Decode(target)
	r.Matched = 1
	return r, err
}

func (o *Operation) getManyDocuments(filters map[string]interface{}, sort map[string]interface{}, limit int, target interface{}) (r OperationResult, err error) {
	// Set sort options
	opts := options.Find()
	opts.SetLimit(int64(limit))
	var sortInterface bson.D
	for k, v := range sort {
		sortInterface = append(sortInterface, bson.E{Key: k, Value: v})
	}
	if len(sortInterface) > 0 {
		opts.SetSort(sort)
	}

	// Compile query
	filter := mapToBSON(filters)

	// Find and decode
	cur, err := o.collection.Find(o.ctx, filter, opts)
	if err != nil {
		return r, err
	}
	r.Matched = int(cur.RemainingBatchLength())
	return r, cur.All(o.ctx, target)
}

func (o *Operation) updateDocumentFields(filters map[string]interface{}, payload map[string]interface{}) (r OperationResult, err error) {
	// Set upsert
	opts := options.Update()

	// Compile query
	filter := mapToBSON(filters)

	// Update
	update := mapToBSON(payload)

	// Update
	result, err := o.collection.UpdateOne(o.ctx, filter, bson.M{"$set": update}, opts)
	if err != nil {
		return r, logs.Wrap(err, "UpdateOne (field) failed")
	}
	r.Inserted = int(result.UpsertedCount)
	r.Updated = int(result.ModifiedCount)
	r.Matched = int(result.MatchedCount)
	return r, nil
}

func (o *Operation) updateDocument(filters map[string]interface{}, upsert bool, payload interface{}) (r OperationResult, err error) {
	// Set upsert
	opts := options.Update()
	opts.SetUpsert(upsert)

	// Compile query
	filter := mapToBSON(filters)

	// Update
	result, err := o.collection.UpdateOne(o.ctx, filter, bson.M{"$set": payload}, opts)
	if err != nil {
		return r, logs.Wrap(err, "UpdateOne failed")
	}
	r.Inserted = int(result.UpsertedCount)
	r.Updated = int(result.ModifiedCount)
	r.Matched = int(result.MatchedCount)
	return r, nil
}

func (o *Operation) insertManyDocuments(payload []interface{}) (r OperationResult, err error) {
	result, err := o.collection.InsertMany(o.ctx, payload)
	if err != nil {
		return r, logs.Wrap(err, "insertManyDocuments failed")
	}
	r.Inserted = len(result.InsertedIDs)
	return r, nil
}

func (o *Operation) deleteManyDocuments(filters map[string]interface{}) (r OperationResult, err error) {
	if len(filters) == 0 {
		return r, fmt.Errorf("filter is required")
	}

	// Compile query
	filter := mapToBSON(filters)

	// Update
	result, err := o.collection.DeleteMany(o.ctx, filter)
	if err != nil {
		return r, logs.Wrap(err, "DeleteMany failed")
	}
	r.Deleted = int(result.DeletedCount)
	r.Matched = int(result.DeletedCount)
	return r, nil
}

func (o *Operation) bulkWrite(models []mongo.WriteModel) (OperationResult, error) {
	r, err := o.collection.BulkWrite(o.ctx, models)

	var result OperationResult
	result.Inserted = int(r.InsertedCount) + int(r.UpsertedCount)
	result.Updated = int(r.ModifiedCount)
	result.Deleted = int(r.DeletedCount)
	result.Matched = int(r.MatchedCount)
	return result, err
}
