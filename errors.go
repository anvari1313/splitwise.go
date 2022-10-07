package splitwise

import (
	"errors"
)

var (
	// ErrInvalidToken will be returned in case of receiving 401 from the service
	ErrInvalidToken = errors.New("invalid token")

	// ErrPermissionDenied will be returned in case of receiving 403 from the service
	ErrPermissionDenied = errors.New("invalid API Request: you do not have permission to perform that action")

	// ErrRecordNotFound will be returned in case of receiving 404 from the service
	ErrRecordNotFound = errors.New("invalid API Request: record not found")

	// ErrSplitwiseServer will be returned on 500 internal server errors
	ErrSplitwiseServer = errors.New("splitwise internal server error")
)
