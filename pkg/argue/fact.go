package argue

import (
	"fmt"
	"strings"
)

// FactType represents the data type that a rule will
// accept.
type FactType int

// FactType Values
const (
	FactTypeBool = iota
	FactTypeString
	FactTypeInt
	FactTypeFloat
)

// Fact represents a rule that an argument must
// follow when parsing command-line arguments. Facts
// are akin to flags.
type Fact struct {
	Type       FactType
	Help       string
	FullName   string
	ShortName  byte
	Positional bool
	Required   bool
}

// NewFact returns a new fact with the given
// parameters.
func NewFact(t FactType, h string, fn string, sn byte, p bool, r bool) Fact {
	var fact Fact
	fact.Type = t
	fact.Help = h
	fact.FullName = strings.ToLower(fn)
	fact.ShortName = sn
	fact.Positional = false
	fact.Required = false
	return fact
}

// InfoString returns a string containing the
// information that will be printed out with the
// usage information.
func (f Fact) InfoString() string {
	if !f.Positional {
		if f.ShortName == 0 {
			return fmt.Sprintf("--%s\t\t%v", f.FullName, f.Help)
		}

		return fmt.Sprintf("-%s, --%s \t%v", string(f.ShortName), f.FullName, f.Help)
	}

	return ""
}

// SetHelp accepts a string and sets the Help
// property of the received fact to that string.
func (f *Fact) SetHelp(h string) {
	f.Help = h
}

// SetShortName accepts a byte and sets the ShortName
// property of the received fact to that byte.
func (f *Fact) SetShortName(b byte) {
	f.ShortName = b
}
