package main

import (
	"os"

	"github.com/Spirit55555/rdap"
)

func main() {
	exitCode := rdap.RunCLI(os.Args[1:], os.Stdout, os.Stderr, rdap.CLIOptions{})

	os.Exit(exitCode)
}
