package errors

import "errors"

var (
	ErrNoDocumentsFound       = errors.New("no documents found")
	ErrOperationCountMismatch = errors.New("operation count mismatch")
	ErrUnauthorized           = errors.New("unauthorized")
	ErrBadStatusCode          = errors.New("bad status code")
	ErrInvalidPayload         = errors.New("invalid payload")
)
