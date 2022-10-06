package splitwise

import (
	"errors"
)

var (
	// ErrInvalidToken will be returned in case of receiving 401 from the service
	ErrInvalidToken = errors.New("invalid token")

	// ErrSplitwiseServer will be returned on 500 internal server errors
	ErrSplitwiseServer = errors.New("splitwise internal server error")
)
