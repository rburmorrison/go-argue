// Package argue implements functions and types that
// assist in parsing command-line arguments.
package argue

import (
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func getBinaryName() string {
	return filepath.Base(os.Args[0])
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

// splitArguments splits command-line arguments into
// their "positional" and "flag" categories. They are
// returned in that order.
func splitArguments(agmt Argument) (map[string]interface{}, map[string]interface{}) {
	// Regex expressions
	flagRegex := regexp.MustCompile(`^(-\S|--\S+)$`)

	// Maps
	positionalMap := make(map[string]interface{})
	flagMap := make(map[string]interface{})

	args := os.Args[1:]
	for i, a := range args {
		sn := a[1:]
		ln := a[2:]

		// Handle default help flag
		if sn == "h" || ln == "help" {
			agmt.PrintUsage()
			os.Exit(0)
		}

		// Handle version flag if implemented
		if agmt.ShowVersion && sn == "v" || ln == "version" {
			agmt.PrintVersion()
			os.Exit(0)
		}

		// Handle flags
		if flagRegex.MatchString(a) {
			name := sn
			if strings.HasPrefix(a, "--") {
				name = ln
			}

			if _, ok := agmt.NameInFlagFacts(name); !ok {
				agmt.PrintError("unknown flag " + a + " provided")
			}

			if f, ok := agmt.NameInFlagFacts(name); ok {
				var val interface{}
				if f.Type != FactTypeBool {
					if len(args)-1 <= i {
						agmt.PrintError("no value supplied to --" + f.FullName)
					} else {
						val = args[i+1]
					}
				} else {
					val = true
				}

				flagMap[name] = val
			}
		}
	}

	// TODO: parse positional arguments
	return positionalMap, flagMap
}

func setFactValue(agmt Argument, f *Fact, v interface{}) {
	val := reflect.ValueOf(f.Value).Elem()
	switch f.Type {
	case FactTypeBool:
		b, ok := v.(bool)
		if !ok {
			agmt.PrintError("--" + f.FullName + " requires a boolean value")
		}
		val.SetBool(b)
	case FactTypeString:
		s, ok := v.(string)
		if !ok {
			agmt.PrintError("--" + f.FullName + " requires a string value")
		}
		val.SetString(s)
	case FactTypeInt:
		s := v.(string)
		i, err := strconv.Atoi(s)
		if err != nil {
			agmt.PrintError("--" + f.FullName + " requires an integer value")
		}
		val.SetInt(int64(i))
	case FactTypeFloat:
		s := v.(string)
		fl, err := strconv.ParseFloat(s, 64)
		if err != nil {
			agmt.PrintError("--" + f.FullName + " requires a float value")
		}
		val.SetFloat(fl)
	}
}
