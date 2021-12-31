package driver

import (
	"context"
	"fmt"
	"time"

	"aftermath.link/repo/logs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c *Driver) SetURI(uri string) error {
	if uri == "" {
		return fmt.Errorf("uri cannot be blank")
	}
	c.uri = uri
	return nil
}

func (c *Driver) SetUser(user string) error {
	if user == "" {
		return fmt.Errorf("user cannot be blank")
	}
	c.user = user
	return nil
}

func (c *Driver) SetPass(pass string) error {
	if pass == "" {
		return fmt.Errorf("pass cannot be blank")
	}
	c.pass = pass
	return nil
}

func (c *Driver) Verify() error {
	if c.uri == "" {
		return fmt.Errorf("uri cannot be blank")
	}
	if c.user == "" {
		return fmt.Errorf("user cannot be blank")
	}
	if c.pass == "" {
		return fmt.Errorf("pass cannot be blank")
	}
	return nil
}

func (c *Driver) newOperation(database, collection string) (*Operation, error) {
	if database == "" {
		return nil, fmt.Errorf("invalid database name")
	}
	if collection == "" {
		return nil, fmt.Errorf("invalid collection name")
	}

	client, err := c.connect()
	if err != nil {
		return nil, logs.Wrap(err, "Connect failed")
	}

	var o Operation
	o.ctx = context.Background()
	o.client = client
	o.finish = func() {
		o.client.Disconnect(o.ctx)
	}

	db := o.client.Database(database)
	if db == nil {
		return nil, fmt.Errorf("database is nil")
	}
	o.collection = db.Collection(collection)
	if o.collection == nil {
		return nil, fmt.Errorf("collection is nil")
	}
	return &o, nil
}

func (d *Driver) connect() (*mongo.Client, error) {
	credential := options.Credential{
		Username: d.user,
		Password: d.pass,
	}

	clientOpts := options.Client().ApplyURI(d.uri).SetAuth(credential)
	client, err := mongo.NewClient(clientOpts)
	if err != nil {
		return nil, logs.Wrap(err, "mongo.NewClient failed")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, logs.Wrap(err, "client.Connect failed")
	}

	// Check connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, logs.Wrap(err, "client.Ping failed")
	}
	return client, nil
}
