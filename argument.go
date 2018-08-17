package argue

import (
	"sort"
	"strings"
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
	if a.NameExists(name) {
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
	if a.NameExists(name) {
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

// NameExists returns true if any facts within the
// argument has the name passed, false otherwise.
func (a Argument) NameExists(n string) bool {
	fs := a.Facts()

	for _, f := range fs {
		if f.Name == n {
			return true
		}
	}

	return false
}

// InitialExists returns true if any flag facts
// within the argument have the initial passed
// (excluding 0), false otherwise.
func (a Argument) InitialExists(i byte) bool {
	fs := a.FlagFacts

	for _, f := range fs {
		if f.Initial == i {
			return true
		}
	}

	return false
}

// GenerateInitial accepts a name and returns an
// appropriate inital for it based on existing fact
// initials. If all reasonable initials are taken,
// 0 will be returned instead.
func (a Argument) GenerateInitial(n string) byte {
	// Check if first letter is valid, if not, try it's
	// upper case version
	fl := n[0]
	if a.InitialExists(fl) || fl == byte("h"[0]) ||
		(a.ShowVersion && fl == byte("v"[0])) {
		fl = byte(strings.ToUpper(string(fl))[0])
	}

	// If unavailable, return 0
	if a.InitialExists(fl) {
		return 0
	}

	return fl
}

// Facts returns all the facts of an arguments,
// positional and otherwise.
func (a Argument) Facts() []*Fact {
	return append(a.PositionalFacts, a.FlagFacts...)
}
