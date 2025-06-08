package morse

import (
	"testing"
)

func TestEncodeSOS(t *testing.T) {
	input := "SOS"
	expected := "... --- ..."
	result := Encode(input)
	if result != expected {
		t.Errorf("Encode(%q) = %q; want %q", input, result, expected)
	}
}

func TestEncodeHelloWorld(t *testing.T) {
	input := "HELLO WORLD"
	expected := ".... . .-.. .-.. ---   .-- --- .-. .-.. -.."
	result := Encode(input)
	if result != expected {
		t.Errorf("Encode(%q) = %q; want %q", input, result, expected)
	}
}

func TestEncodeNumbers(t *testing.T) {
	input := "123"
	expected := ".---- ..--- ...--"
	result := Encode(input)
	if result != expected {
		t.Errorf("Encode(%q) = %q; want %q", input, result, expected)
	}
}

func TestEncodeEmptyString(t *testing.T) {
	input := ""
	expected := ""
	result := Encode(input)
	if result != expected {
		t.Errorf("Encode(%q) = %q; want %q", input, result, expected)
	}
}

func TestEncodeUnknownCharacter(t *testing.T) {
	input := "UNKNOWN!"
	expected := "..- -. -.- -. --- .-- -. ?"
	result := Encode(input)
	if result != expected {
		t.Errorf("Encode(%q) = %q; want %q", input, result, expected)
	}
}

func TestEncodeSingleCharacter(t *testing.T) {
	input := "A"
	expected := ".-"
	result := Encode(input)
	if result != expected {
		t.Errorf("Encode(%q) = %q; want %q", input, result, expected)
	}
}

func TestEncodeSpace(t *testing.T) {
	input := " "
	expected := ""
	result := Encode(input)
	if result != expected {
		t.Errorf("Encode(%q) = %q; want %q", input, result, expected)
	}
}

func TestEncodeInvalidCharacter(t *testing.T) {
	input := "@"
	expected := "?"
	result := Encode(input)
	if result != expected {
		t.Errorf("Encode(%q) = %q; want %q", input, result, expected)
	}
}

func TestDecodeSOS(t *testing.T) {
	input := "... --- ..."
	expected := "SOS"
	result := Decode(input)
	if result != expected {
		t.Errorf("Decode(%q) = %q; want %q", input, result, expected)
	}
}

func TestDecodeHelloWorld(t *testing.T) {
	input := ".... . .-.. .-.. ---   .-- --- .-. .-.. -.."
	expected := "HELLO WORLD"
	result := Decode(input)
	if result != expected {
		t.Errorf("Decode(%q) = %q; want %q", input, result, expected)
	}
}

func TestDecodeNumbers(t *testing.T) {
	input := ".---- ..--- ...--"
	expected := "123"
	result := Decode(input)
	if result != expected {
		t.Errorf("Decode(%q) = %q; want %q", input, result, expected)
	}
}

func TestDecodeEmptyString(t *testing.T) {
	input := ""
	expected := ""
	result := Decode(input)
	if result != expected {
		t.Errorf("Decode(%q) = %q; want %q", input, result, expected)
	}
}

func TestDecodeUnknownCharacter(t *testing.T) {
	input := "..- -. -.- -. --- .-- -. ?"
	expected := "UNKNOWN?"
	result := Decode(input)
	if result != expected {
		t.Errorf("Decode(%q) = %q; want %q", input, result, expected)
	}
}

func TestDecodeSingleCharacter(t *testing.T) {
	input := ".-"
	expected := "A"
	result := Decode(input)
	if result != expected {
		t.Errorf("Decode(%q) = %q; want %q", input, result, expected)
	}
}

func TestDecodeSpace(t *testing.T) {
	input := "/"
	expected := "?"
	result := Decode(input)
	if result != expected {
		t.Errorf("Decode(%q) = %q; want %q", input, result, expected)
	}
}

func TestDecodeInvalidCharacter(t *testing.T) {
	input := "?"
	expected := "?"
	result := Decode(input)
	if result != expected {
		t.Errorf("Decode(%q) = %q; want %q", input, result, expected)
	}
}
