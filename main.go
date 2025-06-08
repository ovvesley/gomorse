package main

import (
	"fmt"
	"os"
	"ovvesley/gomorse/cli"
)

func main() {
	args, err := cli.ParseArgs()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	args.Handle()
}
