package argue

import (
	"errors"
	"regexp"
	"sort"
	"strings"
)

// Error Definitions
var (
	ErrMismatchedPositionals = errors.New("argue: too many positional arguments provided")
)

// Argument represents the rules and information that
// argue must follow when parsing command-line
// arguments.
type Argument struct {
	Description     string
	Version         string
	PositionalFacts []*Fact
	FlagFacts       []*Fact
	ShowDesc        bool
	ShowVersion     bool
}

// SplitArguments splits command-line arguments into
// their "positional" and "flag" categories. They are
// returned in that order. The passed arguments
// should not include the call to the binary.
func (a Argument) SplitArguments(arguments []string) ([]string, map[string]interface{}) {
	// Define regular expressions
	flagReg := regexp.MustCompile(`^-\S+$`)

	// Define structures to return
	var positionalSlice []string
	var flagMap = make(map[string]interface{})
	for len(arguments) > 0 {
		arg := arguments[0]
		if !flagReg.MatchString(arg) {
			positionalSlice = append(positionalSlice, arg)

			// Remove this argument from the total list
			arguments = arguments[1:]
		} else {
			// If the argument is not a defined fact, treat it as
			// a boolean flag, otherwise get the fact that this
			// argument represents
			f, nameExists := a.DressedNameExists(arg)
			if !nameExists {
				f2, initialExists := a.DressedInitialExists(arg)
				if !initialExists {
					flagMap[arg] = true

					// Remove this argument from the total list and
					// restart the loop
					arguments = arguments[1:]
					continue
				} else {
					f = f2
				}
			}

			// Treat boolean arguments specially since they do
			// not require a value
			if f.Type == FactTypeBool {
				flagMap[arg] = true

				// Remove this argument from the total list
				arguments = arguments[1:]
			} else {
				// If the next argument is another flag, or if this
				// is the last argument, assign it a nil value and
				// continue on
				if len(arguments) <= 1 || flagReg.MatchString(arguments[1]) {
					flagMap[arg] = nil

					// Remove this argument from the total list and
					// restart the loop
					arguments = arguments[1:]
					continue
				}

				// Treat the next argument as the value to this one
				flagMap[arg] = arguments[1]

				// Remove this argument and the next from the total
				// list
				arguments = arguments[2:]
			}
		}
	}

	return positionalSlice, flagMap
}

// NewArgument accepts a description and will return
// a new Argument with that description and default
// values. NewArgument also sets ShowDesc and
// ShowVersion to true.
func NewArgument(desc string, version string) Argument {
	var agmt Argument
	agmt.Description = desc
	agmt.Version = version
	agmt.ShowDesc = true
	agmt.ShowVersion = true
	return agmt
}

// NewEmptyArgument returns a new argument without a
// description or version, and sets ShowDesc and
// ShowVersion to false.
func NewEmptyArgument() Argument {
	agmt := NewArgument("", "")
	agmt.ShowDesc = false
	agmt.ShowVersion = false
	return agmt
}

// AddFlagFact creates a new flag fact based on a
// name, help description, and a reference to a
// variable to place the parsed contents.
func (a *Argument) AddFlagFact(name string, help string, v interface{}) *Fact {
	name = StandardizeFactName(name)
	if _, ok := a.NameExists(name); ok {
		panic("argument: name already exits within this argument")
	}

	fact := NewFact(help, name, a.GenerateInitial(name), false, false, v)
	a.FlagFacts = append(a.FlagFacts, &fact)
	a.SortFlagFacts()
	return &fact
}

// AddPositionalFact creates a new positional fact
// based on a name, help description, and a reference
// to a variable to place the parsed contents.
func (a *Argument) AddPositionalFact(name string, help string, v interface{}) *Fact {
	name = StandardizeFactName(name)
	if _, ok := a.NameExists(name); ok {
		panic("argument: name already exits within this argument")
	}

	fact := NewFact(help, name, a.GenerateInitial(name), true, false, v)
	a.PositionalFacts = append(a.PositionalFacts, &fact)
	return &fact
}

// SortFlagFacts sorts the flag facts in an argument
// by fact type.
func (a *Argument) SortFlagFacts() {
	sort.Slice(a.FlagFacts, func(i, j int) bool {
		return a.FlagFacts[i].Initial < a.FlagFacts[j].Initial
	})
}

// DressedNameExists returns true if any facts within
// the argument has the dressed name passed, false
// otherwise.
func (a Argument) DressedNameExists(dn string) (*Fact, bool) {
	fs := a.Facts()

	for _, f := range fs {
		if f.DressedName() == dn {
			return f, true
		}
	}

	return nil, false
}

// DressedInitialExists returns true if any flag
// facts within the argument have the dressed initial
// passed, false otherwise.
func (a Argument) DressedInitialExists(di string) (*Fact, bool) {
	fs := a.FlagFacts

	for _, f := range fs {
		if f.DressedInitial() == di {
			return f, true
		}
	}

	return nil, false
}

// NameExists returns true if any facts within the
// argument has the name passed, false otherwise.
func (a Argument) NameExists(n string) (*Fact, bool) {
	fs := a.Facts()

	for _, f := range fs {
		if f.Name == n {
			return f, true
		}
	}

	return nil, false
}

// InitialExists returns true if any flag facts
// within the argument have the initial passed
// (excluding 0), false otherwise.
func (a Argument) InitialExists(i byte) (*Fact, bool) {
	fs := a.FlagFacts

	for _, f := range fs {
		if f.Initial == i {
			return f, true
		}
	}

	return nil, false
}

// GenerateInitial accepts a name and returns an
// appropriate inital for it based on existing fact
// initials. If all reasonable initials are taken,
// 0 will be returned instead.
func (a Argument) GenerateInitial(n string) byte {
	// Check if first letter is valid, if not, try it's
	// upper case version
	fl := n[0]
	if _, ok := a.InitialExists(fl); ok || fl == byte("h"[0]) ||
		(a.ShowVersion && fl == byte("v"[0])) {
		fl = byte(strings.ToUpper(string(fl))[0])
	}

	// If unavailable, return 0
	if _, ok := a.InitialExists(fl); ok {
		return 0
	}

	return fl
}

// Facts returns all the facts of an arguments,
// positional and otherwise.
func (a Argument) Facts() []*Fact {
	return append(a.PositionalFacts, a.FlagFacts...)
}
