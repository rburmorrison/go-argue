// Package argue implements functions for parsing
// command-line arguments.
package argue

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"unicode"
)

// Error Definitions
var (
	// Lawyer
	ErrUnknownCommand = errors.New("argue: unknown command provided")
	ErrNoCommand      = errors.New("argue: no command specified")

	// Argument
	ErrExtraPositionals   = errors.New("argue: too many positional arguments provided")
	ErrMissingPositionals = errors.New("argue: not enough positional arguments provided")
	ErrMissingFlag        = errors.New("argue: a required flag is missing")
	ErrUnknownFlag        = errors.New("argue: dispute found an unknown flag while parsing")

	// Fact
	ErrWrongType   = errors.New("argue: fact was not able to set a value due to mismatched types")
	ErrNilValue    = errors.New("argue: nil was passed to a flag")
	ErrInvalidType = errors.New("argue: invalid type passed to GetFactType. " +
		"Options are *string, *bool, *int, *int64, *uint, *uint64, *float32, and *float64")
)

// Define regular expressions
var (
	flagReg = regexp.MustCompile(`^(-\S|--\S+)$`)
)

func binaryName() string {
	return filepath.Base(os.Args[0])
}

func breakCammelCase(s string) string {
	var broken string
	for i, r := range s {
		if i != 0 && unicode.IsUpper(rune(r)) {
			broken += "-" + string(r)
		} else {
			broken += string(r)
		}
	}

	return broken
}

func printError(msg string) {
	fmt.Printf("Error: %v\n", msg)
	fmt.Printf("Run \"%v --help\" to see usage information\n", os.Args[0])
	os.Exit(1)
}
