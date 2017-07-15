package client

import "errors"

var (
	// ErrInvalidMediaType is used when an incorrect media type is detected by the client.
	ErrInvalidMediaType = errors.New("invalid media type")

	// ErrInvalidSirenEntity is used when a response body can not be decoded as a siren entity.
	ErrInvalidSirenEntity = errors.New("invalid siren entity")
)
