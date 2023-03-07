package common

import (
	"strings"
)

func EmptyString(str string) bool {
	if strings.TrimSpace(str) == "" {
		return true
	}
	return false
}
