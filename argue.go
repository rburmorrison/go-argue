// Package argue implements functions for parsing
// command-line arguments.
package argue

import (
	"os"
	"path/filepath"
	"unicode"
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
