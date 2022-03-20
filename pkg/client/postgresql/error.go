package postgresql

import "errors"

var ErrConnect = errors.New("failed postgresql connection")
var ErrSqlQueryFailed = errors.New("sql failed query")
var ErrSqlScanFailed = errors.New("sql scan query")
var ErrNotFound = errors.New("not found")
