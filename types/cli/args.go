package types_cli

import (
	"fmt"
	"ovvesley/gomorse/internal/morse"
)

type Args struct {
	morseInput   string
	literalInput string
	Handle       func() error
}

func (a *Args) MorseInput() string {
	return a.morseInput
}

func (a *Args) LiteralInput() string {
	return a.literalInput
}

func NewArgsMorseInput(input string) *Args {
	return &Args{
		morseInput: input,
		Handle: func() error {
			decodedMessage := morse.Decode(input)
			if decodedMessage == "" {
				return fmt.Errorf("invalid morse input: %s", input)
			}
			fmt.Println(decodedMessage)
			return nil
		},
	}
}

func NewArgsLiteralInput(input string) *Args {
	return &Args{
		literalInput: input,
		Handle: func() error {
			encodedMessage := morse.Encode(input)
			if encodedMessage == "" {
				return fmt.Errorf("invalid literal input: %s", input)
			}
			fmt.Println(encodedMessage)
			return nil
		},
	}
}
