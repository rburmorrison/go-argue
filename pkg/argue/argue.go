// Package argue implements functions and types that
// assist in parsing command-line arguments.
package argue

import (
	"os"
	"strings"
)

func getBinaryName() string {
	return os.Args[0][2:]
}

func determineShortName(a Argument, n string) byte {
	sn := string(n[0])
	if a.ContainsShortName(byte(sn[0])) || sn == "h" {
		sn = strings.ToUpper(sn)
	}

	if a.ContainsShortName(byte(sn[0])) {
		return 0
	}

	return byte(sn[0])
}
