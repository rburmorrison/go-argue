package argue

import (
	"reflect"
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
	Value      interface{}
}

// NewFact returns a new fact with the given
// parameters.
func NewFact(t FactType, h string, fn string, sn byte, p bool, r bool, v interface{}) Fact {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr {
		panic("argue: variables passed to a fact must be pointers")
	}

	replacer := strings.NewReplacer(" ", "-")

	var fact Fact
	fact.Type = t
	fact.Help = h
	fact.FullName = replacer.Replace(strings.ToLower(fn))
	fact.ShortName = sn
	fact.SetPositional(p)
	fact.Required = false
	return fact
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

// SetPositional accepts a bool and sets the
// Positional property of the received fact to that
// bool.
func (f *Fact) SetPositional(p bool) {
	if p && f.Type == FactTypeBool {
		panic("argue: a fact of type bool can not be positional")
	}

	f.Positional = p
}

// SetRequired accepts a bool and sets the Required
// property of the received fact to that bool.
func (f *Fact) SetRequired(r bool) {
	f.Required = r
}
