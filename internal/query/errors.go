package query

import "errors"

var (
	ErrEmptyAddress     = errors.New("address cannot be empty")
	ErrEmptyID          = errors.New("id cannot be empty")
	ErrNotConfigured    = errors.New("subgraph URL is not configured")
	ErrUnexpectedStatus = errors.New("non-200 status code")
)
