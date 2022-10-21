package database

import (
	"context"
	"errors"
	"time"

	"github.com/byvko-dev/am-core/errors/database"
	"github.com/byvko-dev/am-core/logs"
	"github.com/byvko-dev/am-core/mongodb/driver"
	"github.com/byvko-dev/am-types/users/v2"
	er "github.com/byvko-dev/am-users-api/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserProfileByID(id string) (*users.CompleteProfile, error) {
	client, err := driver.NewClient()
	if err != nil {
		return nil, er.ErrMongoFailedToConnect
	}

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, er.ErrMongoInvalidID
	}

	filter := make(map[string]interface{})
	filter["_id"] = oid

	var profile users.CompleteProfile
	err = client.GetDocumentWithFilter(collectionProfiles, filter, &profile)
	if err != nil {
		defer logs.Error("GetUserProfileByID(%v).GetDocumentWithFilter: %v", id, err.Error())
		if errors.Is(err, database.ErrDocumentNotFound) {
			return nil, er.ErrUserNotFound
		}
		return nil, er.ErrMongoOperationFailed
	}
	return &profile, nil
}

func GetUserBan(id string) (*users.UserBan, error) {
	client, err := driver.NewClient()
	if err != nil {
		return nil, er.ErrMongoFailedToConnect
	}

	filter := make(map[string]interface{})
	filter["user_id"] = id
	filter["lifted"] = false
	filter["expiration"] = bson.M{"$gt": time.Now()}

	var ban users.UserBan
	err = client.GetDocumentWithFilter(collectionBans, filter, &ban)
	if err != nil {
		defer logs.Error("GetUserBan(%v).GetDocumentWithFilter: %v", id, err.Error())
		if errors.Is(err, database.ErrDocumentNotFound) {
			return nil, nil
		}
		return nil, er.ErrMongoOperationFailed
	}
	return &ban, nil
}

func GetUserProfileByExternalID(id interface{}, service users.ExternalService) (*users.CompleteProfile, error) {
	client, err := driver.NewClient()
	if err != nil {
		return nil, er.ErrMongoFailedToConnect
	}

	filter := make(map[string]interface{})
	filter["service"] = service
	filter["external_id"] = id

	var connection users.ExternalConnection
	err = client.GetDocumentWithFilter(collectionConnections, filter, &connection)
	if err != nil {
		defer logs.Error("GetUserProfileByExternalID(%v, %v).GetDocumentWithFilter: %v", id, service, err.Error())
		if errors.Is(err, database.ErrDocumentNotFound) {
			return nil, er.ErrUserNotFound
		}
		return nil, er.ErrMongoOperationFailed
	}

	return GetUserProfileByID(connection.UserID)
}

func GetUserConnections(id string) ([]users.ExternalConnection, error) {
	client, err := driver.NewClient()
	if err != nil {
		return nil, er.ErrMongoFailedToConnect
	}

	filter := bson.M{}
	filter["user_id"] = id

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var connections []users.ExternalConnection
	cur, err := client.Raw(collectionConnections).Find(ctx, filter)
	if err != nil {
		defer logs.Error("GetUserConnections(%v).Find: %v", id, err.Error())
		if errors.Is(err, mongo.ErrNoDocuments) {
			// No connections found
			return nil, nil
		}
		return nil, er.ErrMongoOperationFailed
	}

	err = cur.All(ctx, &connections)
	if err != nil {
		defer logs.Error("GetUserConnections(%v).cur.All: %v", id, err.Error())
		return nil, er.ErrMongoOperationFailed
	}

	return connections, nil
}
