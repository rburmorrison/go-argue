package argue

import (
	"fmt"
	"reflect"
	"regexp"
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
func NewFact(h string, fn string, sn byte, p bool, r bool, v interface{}) Fact {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr {
		panic("argue: variables passed to a Fact must be pointers")
	}

	var t FactType
	ty := fmt.Sprintf("%s", reflect.TypeOf(v))
	switch ty {
	case "*bool":
		t = FactTypeBool
	case "*float64":
		t = FactTypeFloat
	case "*int":
		t = FactTypeInt
	case "*string":
		t = FactTypeString
	default:
		panic("argue: invalid type passed to NewFact. Options are *bool, *float64, *int, and *string")
	}

	// Standardize the name
	replacer := strings.NewReplacer(" ", "-")
	fn = replacer.Replace(strings.ToLower(fn))
	reg := regexp.MustCompile("--+")
	fn = reg.ReplaceAllString(fn, "-")

	var fact Fact
	fact.Type = t
	fact.Help = h
	fact.FullName = fn
	fact.ShortName = sn
	fact.SetPositional(p)
	fact.Required = false
	fact.Value = v
	return fact
}

// SetHelp accepts a string and sets the Help
// property of the received fact to that string.
func (f *Fact) SetHelp(h string) *Fact {
	f.Help = h
	return f
}

// SetShortName accepts a byte and sets the ShortName
// property of the received fact to that byte.
func (f *Fact) SetShortName(b byte) *Fact {
	f.ShortName = b
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

// UpperName returns the upper case full name that is
// typically used with positional argument names.
func (f *Fact) UpperName() string {
	replacer := strings.NewReplacer(" ", "", "-", "")
	name := replacer.Replace(f.FullName)
	return strings.ToUpper(name)
}