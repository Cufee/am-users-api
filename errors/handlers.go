package er

import "errors"

var (
	ErrInvalidPayload          = errors.New("INVALID_PAYLOAD")
	ErrConnectionAlreadyExists = errors.New("CONNECTION_ALREADY_EXISTS")
)
