package argue

import (
	"regexp"
	"strings"
)

// StandardizeFactName accepts a string to be the
// name of a fact and returns a cleaned version of
// it that is compatible with a fact.
func StandardizeFactName(n string) string {
	replacer := strings.NewReplacer(" ", "-")
	reg := regexp.MustCompile("--+")

	n = replacer.Replace(strings.ToLower(n))
	n = reg.ReplaceAllString(n, "-")
	return n
}

// UpperFactName returns the uppercase equivalent
// of a fact name.
func UpperFactName(n string) string {
	replacer := strings.NewReplacer(" ", "", "-", "")
	name := replacer.Replace(n)
	return strings.ToUpper(name)
}
