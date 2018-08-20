// Package argue implements functions for parsing
// command-line arguments.
package argue

import (
	"errors"
	"os"
	"path/filepath"
	"unicode"
)

// Error Definitions
var (
	ErrExtraPositionals   = errors.New("argue: too many positional arguments provided")
	ErrMissingPositionals = errors.New("argue: not enough positional arguments provided")
	ErrUnknownFlag        = errors.New("argue: dispute found an unknown flag while parsing")
	ErrWrongType          = errors.New("argue: fact was not able to set a value due to mismatched types")
	ErrNilValue           = errors.New("argue: nil was passed to a flag")
	ErrInvalidType        = errors.New("argue: invalid type passed to GetFactType. " +
		"Options are *string, *bool, *int, *int64, *uint, *uint64, *float32, and *float64")
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
