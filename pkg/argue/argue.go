// Package argue implements functions and types that
// assist in parsing command-line arguments.
package argue

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func getBinaryName() string {
	return os.Args[0][2:]
}

func determineShortName(a Argument, n string) byte {
	sn := string(n[0])
	if a.ContainsShortName(byte(sn[0])) || sn == "h" || sn == "v" {
		sn = strings.ToUpper(sn)
	}

	if a.ContainsShortName(byte(sn[0])) {
		return 0
	}

	return byte(sn[0])
}

func splitArguments(agmt Argument) (pm map[string]interface{}, fm map[string]interface{}) {
	// Regex expressions
	flagRegex := regexp.MustCompile(`^(-\S|--\S+)$`)

	// Maps
	positionalMap := make(map[string]interface{})
	flagMap := make(map[string]interface{})

	args := os.Args[1:]
	for _, a := range args {
		if flagRegex.MatchString(a) {
			flagMap[a] = "hey"
		}
	}

	fmt.Println(flagMap)

	return positionalMap, flagMap
}
