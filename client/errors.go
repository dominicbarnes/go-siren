package client

import "errors"

var (
	// ErrInvalidSirenEntity is used when a response body can not be decoded as a siren entity.
	ErrInvalidSirenEntity = errors.New("invalid siren entity")
)
