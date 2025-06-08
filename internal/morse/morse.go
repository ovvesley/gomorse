package morse

import (
	"strings"
)

func Encode(input string) string {
	alph := NewAlphabet()
	var result string
	for _, char := range input {
		charUpper := strings.ToUpper(string(char))
		if morse, ok := alph.literalToMorse[charUpper]; ok {
			result += morse + WHITE_SPACE_LITERAL
		} else if charUpper == WHITE_SPACE_LITERAL {
			result = strings.TrimSuffix(result, WHITE_SPACE_LITERAL) + strings.Repeat(WHITE_SPACE_LITERAL, 3)
		} else {
			result += UNKNOWN + WHITE_SPACE_LITERAL
		}
	}

	return strings.TrimSpace(result)
}
func Decode(input string) string {
	alph := NewAlphabet()
	var result string
	morseChars := strings.SplitSeq(input, WHITE_SPACE_LITERAL)
	for morseChar := range morseChars {
		if literal, ok := alph.morseToLiteral[morseChar]; ok {
			result += literal
		} else if morseChar == "" {
			result += WHITE_SPACE_LITERAL
		} else {
			result += UNKNOWN
		}
	}
	result = strings.Join(strings.Fields(result), WHITE_SPACE_LITERAL)
	return strings.TrimSpace(result)
}
