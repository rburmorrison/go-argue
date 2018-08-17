package argue

import (
	"github.com/rburmorrison/go-argue/internal/mirror"
)

// Fact represents a rule that an argument must
// follow when parsing command-line arguments. Facts
// are akin to flags.
type Fact struct {
	Type       FactType
	Help       string
	Name       string
	Initial    byte
	Positional bool
	Required   bool
	Value      interface{}
}

// NewFact returns a new fact with the given
// parameters. Parameters in order are: help, name,
// initial, positional, required, and value.
func NewFact(h string, n string, i byte, p bool, r bool, v interface{}) Fact {
	if !mirror.IsPointer(v) {
		panic("argue: variables passed to a Fact must be pointers")
	}

	// Extract type from value and panic if it is not
	// compatible
	t, err := GetFactType(v)
	if err != nil {
		panic(err)
	}

	var fact Fact
	fact.Type = t
	fact.Help = h
	fact.Name = StandardizeFactName(n)
	fact.Initial = i
	fact.SetPositional(p)
	fact.Required = r
	fact.Value = v
	return fact
}

// DressedName returns the standardized name of a
// fact prefixed with two hyphens.
func (f Fact) DressedName() string {
	return "--" + StandardizeFactName(f.Name)
}

// DressedInitial returns the standardized initial of
// a fact prefixed with one hyphen.
func (f Fact) DressedInitial() string {
	return "-" + string(f.Initial)
}

// SetHelp accepts a string and sets the Help
// property of the received fact to that string.
func (f *Fact) SetHelp(h string) *Fact {
	f.Help = h
	return f
}

// SetInitial accepts a byte and sets the Initial
// property of the received fact to that byte.
func (f *Fact) SetInitial(b byte) *Fact {
	f.Initial = b
	return f
}

// SetPositional accepts a bool and sets the
// Positional property of the received fact to that
// bool.
func (f *Fact) SetPositional(p bool) *Fact {
	if p && f.Type == FactTypeBool {
		panic("argue: a fact of type bool can not be positional")
	}

	f.Positional = p
	return f
}

// SetRequired accepts a bool and sets the Required
// property of the received fact to that bool.
func (f *Fact) SetRequired(r bool) *Fact {
	f.Required = r
	return f
}
