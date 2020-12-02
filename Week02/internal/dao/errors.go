package dao

import (
	"errors"
)

var (
	recordNotFound = errors.New("record not found")
)

func IsNotFoundRecord(err error) bool {
	if errors.Is(err, recordNotFound) {
		return true
	}
	return false
}
