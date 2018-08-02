package argue

import "fmt"

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

// InfoString returns a string containing the
// information that will be printed out with the
// usage information.
func (f Fact) InfoString() string {
	return fmt.Sprintf("-%s, --%s \t%v", string(f.ShortName), f.FullName, f.Help)
}

// SetHelp accepts a help string and sets the Help
// property of the received fact to that string.
func (f *Fact) SetHelp(h string) {
	f.Help = h
}
