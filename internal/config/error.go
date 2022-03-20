package config

import (
	"errors"

	"github.com/ssalamatov/gomaps/internal/server_error"
)

var ErrConfigNotFound = errors.New("config not found")
var ErrMissingEnv = errors.New("missing env var")
var ErrConfigParseFailed = errors.New("parsing failed")

func NewErrConfigNotFound(err error) error {
	return server_error.NewServerError(err, ErrConfigNotFound)
}

func NewErrMissingEnv(err error) error {
	return server_error.NewServerError(err, ErrMissingEnv)
}

func NewErrConfigParseFailed(err error) error {
	return server_error.NewServerError(err, ErrConfigParseFailed)
}
