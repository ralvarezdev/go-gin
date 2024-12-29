package validator

import (
	"errors"
)

var (
	ErrNilHandler = errors.New("validator handler cannot be nil")
)
