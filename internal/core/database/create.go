package database

import (
	"errors"
	"time"

	"github.com/byvko-dev/am-core/mongodb/driver"
	"github.com/byvko-dev/am-types/users/v2"
	er "github.com/byvko-dev/am-users-api/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUserBan(ban users.UserBan) error {
	client, err := driver.NewClient()
	if err != nil {
		return er.ErrMongoFailedToConnect
	}

	_, err = client.InsertDocument(collectionBans, ban)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil
		}
		return er.ErrMongoOperationFailed
	}
	return nil
}

func CreateUserProfile(profile users.CompleteProfile) (string, error) {
	client, err := driver.NewClient()
	if err != nil {
		return "", er.ErrMongoFailedToConnect
	}

	profile.CreatedAt = time.Now()
	profile.UpdatedAt = time.Now()
	profile.LastSeen = time.Now()

	id, err := client.InsertDocument(collectionProfiles, profile)
	if err != nil {
		return "", er.ErrMongoOperationFailed
	}

	oid, ok := id.(primitive.ObjectID)
	if !ok {
		return "", er.ErrMongoOperationFailed
	}

	return oid.Hex(), nil
}
