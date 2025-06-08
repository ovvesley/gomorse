package cli

import (
	"fmt"
	"os"
	types_cli "ovvesley/gomorse/types/cli"
)

var handlers = map[string]func([]string) (*types_cli.Args, error){
	"-m":        HandleMorseInput,
	"--morse":   HandleMorseInput,
	"-l":        HandleLiteralInput,
	"--literal": HandleLiteralInput,
	"-v":        HandleVersion,
	"--version": HandleVersion,
	"-h":        HandleHelp,
	"--help":    HandleHelp,
}

func getHandler(stringArg string) func([]string) (*types_cli.Args, error) {

	if handler, exists := handlers[stringArg]; exists {
		return handler
	}

	if morseIsValid(&stringArg) {
		return HandleMorseInput
	}
	if literalInputIsValid(&stringArg) {
		return HandleLiteralInput
	}

	return HandleDefaultCase

}

func ParseArgs() (*types_cli.Args, error) {
	if len(os.Args) < 2 {
		return HandleDefaultCase(os.Args)
	}

	osArgs := os.Args[1:]
	firstArg := osArgs[0]

	myHandlers := getHandler(firstArg)

	return myHandlers(osArgs)
}

func morseIsValid(_ *string) bool {
	if len(os.Args) < 2 {
		os.Args = append(os.Args, "")
	}

	osArgsInput := os.Args[1]

	morseIsValid := types_cli.IsValidMorseInput(osArgsInput)
	return morseIsValid
}

func literalInputIsValid(_ *string) bool {
	if len(os.Args) < 2 {
		os.Args = append(os.Args, "")
	}

	osArgsInput := os.Args[1]

	literalInputIsValid := types_cli.IsValidLiteralInput(osArgsInput)
	return literalInputIsValid
}

func HandleDefaultCase(osArgs []string) (*types_cli.Args, error) {
	return HandleHelp(osArgs)
}

func HandleMorseInput(osArgs []string) (*types_cli.Args, error) {
	return types_cli.NewArgsMorseInput(osArgs[len(osArgs)-1]), nil
}

func HandleLiteralInput(osArgs []string) (*types_cli.Args, error) {
	return types_cli.NewArgsLiteralInput(osArgs[len(osArgs)-1]), nil
}

func HandleHelp(osArgs []string) (*types_cli.Args, error) {
	fmt.Print(types_cli.Name)
	fmt.Print(types_cli.Example)
	os.Exit(0)
	return &types_cli.Args{}, nil
}

func HandleVersion(osArgs []string) (*types_cli.Args, error) {
	fmt.Println(types_cli.Version, "-", types_cli.License, "-", types_cli.Copyright)
	os.Exit(0)
	return &types_cli.Args{}, nil
}
