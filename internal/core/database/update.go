package database

import (
	"context"
	"errors"
	"time"

	"github.com/byvko-dev/am-core/mongodb/driver"
	"github.com/byvko-dev/am-types/users/v2"
	er "github.com/byvko-dev/am-users-api/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateUserBan(ban users.UserBan, upsert bool) error {
	client, err := driver.NewClient()
	if err != nil {
		return er.ErrMongoFailedToConnect
	}

	filter := make(map[string]interface{})
	filter["_id"] = ban.ID

	update := make(map[string]interface{})
	update["$set"] = ban
	err = client.UpdateDocumentWithFilter(collectionBans, filter, update, upsert)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return er.ErrMongoNotFound
		}
		return er.ErrMongoOperationFailed
	}
	return nil
}

func UpdateUserProfile(profile users.CompleteProfile) (*users.CompleteProfile, error) {
	client, err := driver.NewClient()
	if err != nil {
		return nil, er.ErrMongoFailedToConnect
	}

	oid, _ := primitive.ObjectIDFromHex(profile.ID)
	if oid.IsZero() {
		return nil, er.ErrMongoInvalidID
	}

	profile.UpdatedAt = time.Now()

	filter := bson.M{}
	filter["_id"] = oid
	update := bson.M{}
	update["$set"] = profile

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = client.Raw(collectionProfiles).UpdateOne(ctx, filter, update)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, er.ErrUserNotFound
		}
		return nil, er.ErrMongoOperationFailed
	}

	err = client.Raw(collectionProfiles).FindOne(ctx, filter).Decode(&profile)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, er.ErrUserNotFound
		}
		return nil, er.ErrMongoOperationFailed
	}

	return &profile, nil
}

func UpdateUserConnection(connection users.ExternalConnection) error {
	client, err := driver.NewClient()
	if err != nil {
		return er.ErrMongoFailedToConnect
	}

	filter := make(map[string]interface{})
	filter["user_id"] = connection.UserID
	filter["service"] = connection.Service

	update := make(map[string]interface{})
	update["$set"] = connection

	err = client.UpdateDocumentWithFilter(collectionConnections, filter, update, true)
	if err != nil {
		return er.ErrMongoOperationFailed
	}
	return nil
}
