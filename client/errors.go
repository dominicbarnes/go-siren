package client

import "errors"

var (
	ErrInvalidMediaType   = errors.New("invalid media type")
	ErrInvalidSirenEntity = errors.New("invalid siren entity")
)
