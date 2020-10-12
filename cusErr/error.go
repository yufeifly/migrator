package cusErr

import "errors"

var (
	ErrNotFound       = errors.New("no such record")
	ErrParamsNotValid = errors.New("params not valid")
)
