package driver

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Driver struct {
	uri  string
	user string
	pass string
}

type Operation struct {
	ctx        context.Context
	collection *mongo.Collection
	client     *mongo.Client
	finish     func()
}

type OperationResult struct {
	Inserted int
	Updated  int
	Deleted  int
	Matched  int
}
type QueryResult struct {
	Inserted int `json:"inserted"`
	Updated  int `json:"updated"`
	Deleted  int `json:"deleted"`
	Matched  int `json:"matched"`
}

func (q *QueryResult) Merge(r *QueryResult) {
	q.Inserted += r.Inserted
	q.Updated += r.Updated
	q.Deleted += r.Deleted
	q.Matched += r.Matched
}
