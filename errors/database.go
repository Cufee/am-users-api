package er

import "errors"

var (
	ErrMongoOperationFailed = errors.New("MONGO_OPERATION_FAILED")
	ErrMongoFailedToConnect = errors.New("MONGO_FAILED_TO_CONNECT")
	ErrMongoInvalidID       = errors.New("MONGO_INVALID_ID")
	ErrMongoNotFound        = errors.New("MONGO_NOT_FOUND")
)
