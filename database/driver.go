package database

import (
	"encoding/json"
	"fmt"
	"os"

	"aftermath.link/repo/logs"
	mongodriver "aftermath.link/repo/mongo-driver"
)

const Name = "am-users-api" // Database name

var Driver *mongodriver.SimpleDriver

func InitDriver() error {
	Driver = &mongodriver.SimpleDriver{}
	config, err := driverConfigFromEnv()
	if err != nil {
		return logs.Wrap(err, "driverConfigFromEnv failed")
	}
	Driver.Setup(config)
	return Driver.Advanced.Verify()
}

func InterfaceToStruct(in interface{}, out interface{}) error {
	b, err := json.Marshal(in)
	if err != nil {
		return logs.Wrap(err, "json.Marshal failed")
	}
	err = json.Unmarshal(b, out)
	if err != nil {
		return logs.Wrap(err, "json.Unmarshal failed")
	}
	return nil
}

func driverConfigFromEnv() (mongodriver.SimpleConfig, error) {
	var err error
	var config mongodriver.SimpleConfig
	config.URI, err = getMongoURI()
	if err != nil {
		return config, logs.Wrap(err, "getMongoURI failed")
	}
	config.User, err = getMongoUser()
	if err != nil {
		return config, logs.Wrap(err, "getMongoUser failed")
	}
	config.Password, err = getMongoPass()
	if err != nil {
		return config, logs.Wrap(err, "getMongoPass failed")
	}
	return config, nil
}

func getMongoURI() (string, error) {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		return "", fmt.Errorf("MONGO_URI cannot be blank")
	}
	port := os.Getenv("MONGO_PORT")
	if port == "" {
		return "", fmt.Errorf("MONGO_PORT cannot be blank")
	}

	return uri + ":" + port, nil
}

func getMongoUser() (string, error) {
	uri := os.Getenv("MONGO_USER")
	if uri == "" {
		return "", fmt.Errorf("MONGO_URI cannot be blank")
	}
	return uri, nil
}

func getMongoPass() (string, error) {
	uri := os.Getenv("MONGO_PASS")
	if uri == "" {
		return "", fmt.Errorf("MONGO_URI cannot be blank")
	}
	return uri, nil
}
