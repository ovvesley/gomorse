package types_cli

import (
	"ovvesley/gomorse/internal/morse"
	"strings"
)

func IsValidMorseInput(input string) bool {
	input = strings.TrimSpace(input)
	if input == "" {
		return false
	}
	firstChar := input[0]
	if string(firstChar) == morse.DASH || string(firstChar) == morse.DOT {
		return true
	}
	return false
}

func IsValidLiteralInput(input string) bool {
	return !IsValidMorseInput(input) && input != ""
}
