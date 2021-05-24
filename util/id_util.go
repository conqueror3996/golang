package util

import (
	"strings"

	"github.com/google/uuid"
)

// UUIDV1 ...
func UUIDV1() string {
	return uuid.New().String()
}

// UUIDV2 ...
func UUIDV2() string {
	uuidString := uuid.New().String()
	uuidDelHyphenString := strings.Replace(uuidString, "-", "", -1)
	return uuidDelHyphenString
}
