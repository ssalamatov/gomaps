package postgresql

import (
	"errors"

	"github.com/ssalamatov/gomaps/internal/server_error"
)

var ErrConnect = errors.New("failed postgresql connection")
var ErrQueryFailed = errors.New("failed query")
var ErrNotFound = errors.New("not found")

func NewErrConnect(err error) error {
	return server_error.NewServerError(err, ErrConnect)
}

func NewErrQueryFailed(err error) error {
	return server_error.NewServerError(err, ErrQueryFailed)
}

func NewErrNotFound(err error) error {
	return server_error.NewServerError(err, ErrNotFound)
}
