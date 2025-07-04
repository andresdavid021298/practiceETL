package component

import "strings"

func ValidStringIsBlank(text string) bool {
	return strings.TrimSpace(text) == ""
}

func ValidStringIsNotBlank(text string) bool {
	return !ValidStringIsBlank(text)
}
