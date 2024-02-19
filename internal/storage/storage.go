package storage

import "errors"

var (
	ErrSourceURLNotFound   = errors.New("source wallet url not found")
	ErrTargetURLNotFound   = errors.New("target wallet url not found")
	ErrRequiredURLNotFound = errors.New("required wallet url not found")
	ErrNotEnoughBalance    = errors.New("not enough balance on source wallet")
)
