package cusErr

import "errors"

var (
	ErrNotFound          = errors.New("no such record")
	ErrServiceNotFound   = errors.New("service not found")
	ErrParamsNotValid    = errors.New("params not valid")
	ErrTaskQueueNotFound = errors.New("no such task queue")
)
