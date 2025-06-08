package morse

import (
	"testing"
)

func TestNewAlphabet(t *testing.T) {
	alph := NewAlphabet()

	if alph == nil {
		t.Errorf("Expected NewAlphabet to return a non-nil value")
	}
}

func TestLiteralToMorse_A(t *testing.T) {
	alph := NewAlphabet()

	if alph.literalToMorse["A"] != ".-" {
		t.Errorf("Expected '.-' for literal 'A', got '%s'", alph.literalToMorse["A"])
	}
}

func TestLiteralToMorse_B(t *testing.T) {
	alph := NewAlphabet()

	if alph.literalToMorse["B"] != "-..." {
		t.Errorf("Expected '-...' for literal 'B', got '%s'", alph.literalToMorse["B"])
	}
}

func TestLiteralToMorse_Z(t *testing.T) {
	alph := NewAlphabet()

	if alph.literalToMorse["Z"] != "--.." {
		t.Errorf("Expected '--..' for literal 'Z', got '%s'", alph.literalToMorse["Z"])
	}
}

func TestLiteralToMorse_1(t *testing.T) {
	alph := NewAlphabet()

	if alph.literalToMorse["1"] != ".----" {
		t.Errorf("Expected '.----' for literal '1', got '%s'", alph.literalToMorse["1"])
	}
}

func TestLiteralToMorse_9(t *testing.T) {
	alph := NewAlphabet()

	if alph.literalToMorse["9"] != "----." {
		t.Errorf("Expected '----.' for literal '9', got '%s'", alph.literalToMorse["9"])
	}
}

func TestMorseToLiteral_DotDash(t *testing.T) {
	alph := NewAlphabet()

	if alph.morseToLiteral[".-"] != "A" {
		t.Errorf("Expected 'A' for morse '.-', got '%s'", alph.morseToLiteral[".-"])
	}
}

func TestMorseToLiteral_DashDotDotDot(t *testing.T) {
	alph := NewAlphabet()

	if alph.morseToLiteral["-..."] != "B" {
		t.Errorf("Expected 'B' for morse '-...', got '%s'", alph.morseToLiteral["-..."])
	}
}

func TestMorseToLiteral_DashDashDotDot(t *testing.T) {
	alph := NewAlphabet()

	if alph.morseToLiteral["--.."] != "Z" {
		t.Errorf("Expected 'Z' for morse '--..', got '%s'", alph.morseToLiteral["--.."])
	}
}

func TestMorseToLiteral_DotDashDashDashDash(t *testing.T) {
	alph := NewAlphabet()

	if alph.morseToLiteral[".----"] != "1" {
		t.Errorf("Expected '1' for morse '.----', got '%s'", alph.morseToLiteral[".----"])
	}
}

func TestMorseToLiteral_DashDashDashDashDot(t *testing.T) {
	alph := NewAlphabet()

	if alph.morseToLiteral["----."] != "9" {
		t.Errorf("Expected '9' for morse '----.', got '%s'", alph.morseToLiteral["----."])
	}
}

func TestUnknownLiteral(t *testing.T) {
	alph := NewAlphabet()

	if _, exists := alph.literalToMorse["?"]; exists {
		t.Errorf("Unexpected entry for unknown literal '?' in literalToMorse map")
	}
}

func TestUnknownMorse(t *testing.T) {
	alph := NewAlphabet()

	if _, exists := alph.morseToLiteral["......"]; exists {
		t.Errorf("Unexpected entry for unknown morse '......' in morseToLiteral map")
	}
}
