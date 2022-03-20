package config

import "errors"

var ErrConfigNotFound = errors.New("config not found")
var ErrMissingEnv = errors.New("missing env var")
var ErrConfigParseFailed = errors.New("parsing failed")
