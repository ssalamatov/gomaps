package server_error

import (
	"fmt"
)

type ServerError struct {
	InternalError error
	TypeError     error
}

func NewServerError(InternalError error, TypeError error) *ServerError {
	return &ServerError{InternalError: InternalError, TypeError: TypeError}
}

func (err *ServerError) Error() string {
	if err.InternalError != nil {
		return fmt.Sprintf("%s. %s", err.TypeError.Error(), err.InternalError.Error())
	}
	return fmt.Sprint(err.TypeError.Error())
}
