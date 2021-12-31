package driver

import "go.mongodb.org/mongo-driver/bson"

func mapToBSON(filter map[string]interface{}) (document bson.D) {
	for k, v := range filter {
		document = append(document, bson.E{Key: k, Value: v})
	}
	return document
}
