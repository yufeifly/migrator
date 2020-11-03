package utils

import (
	"strconv"
	"strings"
)

// RenameService give the service a new name,
func RenameService(oldName string) string {
	var newName string
	dot := strings.Index(oldName, ".")
	newName = oldName[:dot+1] + adder(oldName[dot+1:])
	return newName
}

// adder service tail index add one
func adder(tail string) string {
	num, _ := strconv.Atoi(tail)
	num++
	return strconv.Itoa(num)
}
