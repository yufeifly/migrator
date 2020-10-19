package utils

import (
	"strconv"
	"strings"
)

func MakeNameForService(oldName string) string {
	var newName string
	dot := strings.Index(oldName, ".")
	newName = oldName[:dot+2] + adder(oldName[dot+2:])
	return newName
}

func adder(tail string) string {
	num, _ := strconv.Atoi(tail)
	num++
	return strconv.Itoa(num)
}
