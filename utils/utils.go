package utils

import "os"

// IsDebugEnabled checks whether the debug flag is set or not.
func TargetNode() bool {
	return os.Getenv("DST") != ""
}
