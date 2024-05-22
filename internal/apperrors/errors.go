package apperrors

import "errors"

var (
	ErrAlreadyExists      = errors.New("resource already exists")
	ErrNotFound           = errors.New("resource not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrMalformedRequest   = errors.New("malformed request")
)
