package argue

import "strings"

// Argument represents the rules and information that
// argue must follow when parsing command-line
// arguments.
type Argument struct {
	Description    string
	Version        string
	PositonalFacts []*Fact
	FlagFacts      []*Fact
	ShowDesc       bool
	ShowVersion    bool
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

// InitialExists returns true if any facts within the
// argument has the initial passed (excluding 0),
// false otherwise.
func (a Argument) InitialExists(i byte) bool {
	fs := a.Facts()

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
	return append(a.PositonalFacts, a.FlagFacts...)
}
