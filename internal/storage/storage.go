package storage

import "errors"

var (
	ErrUrlNotFound = errors.New("URL not found")
	ErrUrlExist    = errors.New("URL exists")
)
