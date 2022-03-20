package server

import (
	"errors"

	"github.com/ssalamatov/gomaps/internal/server_error"
)

var ErrValidation = errors.New("validation failed")

func NewErrValidation(err error) error {
	return server_error.NewServerError(err, ErrValidation)
}
