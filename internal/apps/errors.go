package apps

import "errors"

var (
	ErrUnique     = errors.New("resource already exists")
	ErrNotExist   = errors.New("resource does not exist")
	ErrUnexpected = errors.New("unexpected error")
	ErrDataSource = errors.New("data source error")
)
