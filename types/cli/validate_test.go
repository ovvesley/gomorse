package types_cli

import (
	"ovvesley/gomorse/internal/morse"
	"testing"
)

func TestIsValidMorseInput_ValidMorseSequence(t *testing.T) {
	morse.DASH = "-"
	morse.DOT = "."

	input := "-.-"
	expected := true
	result := IsValidMorseInput(input)
	if result != expected {
		t.Errorf("IsValidMorseInput(%q) = %v; want %v", input, result, expected)
	}
}

func TestIsValidMorseInput_SOS(t *testing.T) {
	morse.DASH = "-"
	morse.DOT = "."

	input := "...---..."
	expected := true
	result := IsValidMorseInput(input)
	if result != expected {
		t.Errorf("IsValidMorseInput(%q) = %v; want %v", input, result, expected)
	}
}

func TestIsValidMorseInput_InvalidCharacters(t *testing.T) {
	morse.DASH = "-"
	morse.DOT = "."

	input := "AB"
	expected := false
	result := IsValidMorseInput(input)
	if result != expected {
		t.Errorf("IsValidMorseInput(%q) = %v; want %v", input, result, expected)
	}
}

func TestIsInvalidMorseInput_EmptyString(t *testing.T) {
	morse.DASH = "-"
	morse.DOT = "."

	input := ""
	expected := false
	result := IsValidMorseInput(input)
	if result != expected {
		t.Errorf("IsValidMorseInput(%q) = %v; want %v", input, result, expected)
	}
}

func TestIsValidMorseInput_OnlySpaces(t *testing.T) {
	morse.DASH = "-"
	morse.DOT = "."

	input := "   "
	expected := false
	result := IsValidMorseInput(input)
	if result != expected {
		t.Errorf("IsValidMorseInput(%q) = %v; want %v", input, result, expected)
	}
}

func TestIsValidMorseInput_LeadingSpaces(t *testing.T) {
	morse.DASH = "-"
	morse.DOT = "."

	input := "   -.-"
	expected := true
	result := IsValidMorseInput(input)
	if result != expected {
		t.Errorf("IsValidMorseInput(%q) = %v; want %v", input, result, expected)
	}
}

func TestIsValidLiteralInput_ValidLiteralInput(t *testing.T) {
	morse.DASH = "-"
	morse.DOT = "."

	input := "Hello"
	expected := true
	result := IsValidLiteralInput(input)
	if result != expected {
		t.Errorf("IsValidLiteralInput(%q) = %v; want %v", input, result, expected)
	}
}

func TestIsValidLiteralInput_InvalidMorseSequence(t *testing.T) {
	morse.DASH = "-"
	morse.DOT = "."

	input := "-.-"
	expected := false
	result := IsValidLiteralInput(input)
	if result != expected {
		t.Errorf("IsValidLiteralInput(%q) = %v; want %v", input, result, expected)
	}
}
