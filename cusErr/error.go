package cusErr

import "errors"

var (
	// ErrNotFound ...
	ErrNotFound = errors.New("no such record")
	// ErrServiceNotFound ...
	ErrServiceNotFound = errors.New("service not found")
	// ErrParamsNotValid ...
	ErrParamsNotValid = errors.New("params not valid")
	// ErrTaskQueueNotFound ...
	ErrTaskQueueNotFound = errors.New("no such task queue")
)
