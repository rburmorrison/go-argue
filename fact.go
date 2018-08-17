package argue

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

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

// SetValue accepts a value and attempts to assign
// the Value property of the received fact it's
// parsed value. An error will be returned if that
// is not possible.
func (f *Fact) SetValue(v interface{}) error {
	val := reflect.ValueOf(f.Value).Elem()
	switch f.Type {
	case FactTypeString:
		s, ok := v.(string)
		if !ok {
			return errors.New("requires a string value")
		}
		val.SetString(s)
	case FactTypeBool:
		b, ok := v.(bool)
		if !ok {
			return errors.New("requires a boolean value")
		}
		val.SetBool(b)
	case FactTypeInt:
		fallthrough
	case FactTypeInt64:
		s := v.(string)
		i, err := strconv.Atoi(s)
		if err != nil {
			return errors.New("requires an integer value")
		}
		val.SetInt(int64(i))
	case FactTypeUInt:
		fallthrough
	case FactTypeUInt64:
		s := v.(string)
		i, err := strconv.Atoi(s)
		if err != nil {
			return errors.New("requires an non-negative integer value")
		}
		val.SetUint(uint64(i))
	case FactTypeFloat32:
		s := v.(string)
		fl, err := strconv.ParseFloat(s, 32)
		if err != nil {
			return errors.New("requires a 32-bit float value")
		}
		val.SetFloat(fl)
	case FactTypeFloat64:
		s := v.(string)
		fl, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return errors.New("requires a 64-bit float value")
		}
		val.SetFloat(fl)
	}

	return nil
}

// usageHeader returns a string to be used with
// the argument.PrintUsage function.
func (f Fact) usageHeader() string {
	if f.Positional {
		return fmt.Sprintf("%s", UpperFactName(f.Name))
	}

	var s string
	if f.Initial == 0 {
		s = fmt.Sprintf("%s", f.DressedName())
	} else {
		s = fmt.Sprintf("%s %s", f.DressedInitial(), f.DressedName())
	}

	if f.Type != FactTypeBool {
		valueString := " VALUE"
		s += valueString
	}

	return s
}
