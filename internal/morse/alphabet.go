package morse

type AlphabetMap map[string]string

var WHITE_SPACE_LITERAL = " "
var WHITE_SPACE_MORSE = "   "
var DOT = "."
var DASH = "-"
var UNKNOWN = "?"

var alphabetMap = AlphabetMap{
	".-":    "A",
	"-...":  "B",
	"-.-.":  "C",
	"-..":   "D",
	".":     "E",
	"..-.":  "F",
	"--.":   "G",
	"....":  "H",
	"..":    "I",
	".---":  "J",
	"-.-":   "K",
	".-..":  "L",
	"--":    "M",
	"-.":    "N",
	"---":   "O",
	".--.":  "P",
	"--.-":  "Q",
	".-.":   "R",
	"...":   "S",
	"-":     "T",
	"..-":   "U",
	"...-":  "V",
	".--":   "W",
	"-..-":  "X",
	"-.--":  "Y",
	"--..":  "Z",
	"-----": "0",
	".----": "1",
	"..---": "2",
	"...--": "3",
	"....-": "4",
	".....": "5",
	"-....": "6",
	"--...": "7",
	"---..": "8",
	"----.": "9",
}

type Alphabet struct {
	literalToMorse map[string]string
	morseToLiteral map[string]string
}

func NewAlphabet() *Alphabet {
	alph := Alphabet{
		literalToMorse: make(map[string]string),
		morseToLiteral: make(map[string]string),
	}

	for morse, literal := range alphabetMap {
		alph.literalToMorse[literal] = morse
		alph.morseToLiteral[morse] = literal
	}
	return &alph
}
