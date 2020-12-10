package utils

import "os"

// IsDebugEnabled checks whether the debug flag is set or not.
func IsDSTNode() bool {
	return os.Getenv("DST") != ""
}
