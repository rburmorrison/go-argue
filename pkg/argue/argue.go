// Package argue implements functions and types that
// assist in parsing command-line arguments.
package argue

import (
	"os"
	"reflect"
	"regexp"
	"strconv"
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
		if flagRegex.MatchString(a) {
			name := a[1:]
			if strings.HasPrefix(a, "--") {
				name = a[2:]
			}

			// TODO: adjust to only look at flag facts
			if f, ok := agmt.NameInFlagFacts(name); ok {
				var val interface{}
				if f.Type != FactTypeBool {
					if len(args)-1 <= i {
						panic("argue: no value supplied to --" + f.FullName)
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

func setFactValue(f *Fact, v interface{}) {
	val := reflect.ValueOf(f.Value).Elem()
	switch f.Type {
	case FactTypeBool:
		b, ok := v.(bool)
		if !ok {
			panic("argue: invalid type passed to bool flag")
		}
		val.SetBool(b)
	case FactTypeString:
		s, ok := v.(string)
		if !ok {
			panic("argue: invalid type passed to string flag")
		}
		val.SetString(s)
	case FactTypeInt:
		s := v.(string)
		i, err := strconv.Atoi(s)
		if err != nil {
			panic("argue: invalid type passed to int flag")
		}
		val.SetInt(int64(i))
	case FactTypeFloat:
		s := v.(string)
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			panic("argue: invalid type passed to float flag")
		}
		val.SetFloat(f)
	}
}
