// Package argue implements functions and types that
// assist in parsing command-line arguments.
package argue

import "os"

func getBinaryName() string {
	return os.Args[0][2:]
}
