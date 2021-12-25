package database

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"aftermath.link/repo/am-users-api/external"
	"aftermath.link/repo/am-users-api/external/database/errors"
	"aftermath.link/repo/logs"
)

type Opertation struct {
	Method string
	Path   string
}

var (
	backendURI          = "http://localhost:3000"
	operationFindMany   = Opertation{Method: "POST", Path: "findMany"}
	operationInseryMany = Opertation{Method: "PUT", Path: "insertMany"}
	operationUpdateMany = Opertation{Method: "POST", Path: "updateMany"}
	operationDeleteMany = Opertation{Method: "DELETE", Path: "deleteMany"}
)

func sendCollectionRequest(collection string, operation Opertation, payload interface{}, target interface{}) error {
	var url = fmt.Sprintf("%s/collections/%s/%s", backendURI, collection, operation.Path)
	return sendRequest(operation.Method, url, payload, target)
}

func sendRequest(url, method string, payload interface{}, target interface{}) error {
	var err error
	var bodyBytes []byte = nil

	headers, err := prepAuthHeader()
	if err != nil {
		return logs.Wrap(err, "prepAuthHeader failed")
	}

	if payload != nil {
		bodyBytes, err = json.Marshal(payload)
		if err != nil {
			return logs.Wrap(err, "json.Marshal failed")
		}
		headers["Content-Type"] = "application/json"
	}

	status, err := external.HTTPRequest(url, method, headers, bodyBytes, target)
	if err != nil {
		return logs.Wrap(err, "HTTPRequest failed")
	}
	if status != http.StatusOK {
		if status == http.StatusUnauthorized {
			return errors.ErrUnauthorized
		}
		return fmt.Errorf("HTTP status code: %v", status)
	}
	return nil
}

func prepAuthHeader() (map[string]string, error) {
	token := os.Getenv("API_TOKEN")
	if token == "" {
		logs.Critical("API_TOKEN not set")
		return nil, fmt.Errorf("API_TOKEN is not set")
	}

	return map[string]string{
		"Authorization": "Bearer " + token,
	}, nil
}

func FindInCollection(collection string, payload FindRequest) (*FindResponse, error) {
	var response FindResponse
	err := sendCollectionRequest(collection, operationFindMany, payload, &response)
	if err != nil {
		return nil, logs.Wrap(err, "sendCollectionRequest failed")
	}
	return &response, nil
}

func InsertToCollection(collection string, payload InsertRequest) (*InsertResponse, error) {
	var response InsertResponse
	err := sendCollectionRequest(collection, operationInseryMany, payload, &response)
	if err != nil {
		return nil, logs.Wrap(err, "sendCollectionRequest failed")
	}
	return &response, nil
}

func UpdateInCollection(collection string, payload UpdateRequest) (*UpdateResponse, error) {
	var response UpdateResponse
	err := sendCollectionRequest(collection, operationUpdateMany, payload, &response)
	if err != nil {
		return nil, logs.Wrap(err, "sendCollectionRequest failed")
	}
	return &response, nil
}

func DeleteInCollection(collection string, payload DeleteRequest) (*DeleteResponse, error) {
	var response DeleteResponse
	err := sendCollectionRequest(collection, operationDeleteMany, payload, &response)
	if err != nil {
		return nil, logs.Wrap(err, "sendCollectionRequest failed")
	}
	return &response, nil
}
