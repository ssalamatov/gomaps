package server

import "errors"

var ErrValidation = errors.New("validation failed")
var ErrDecodeBody = errors.New("body decode failed")
