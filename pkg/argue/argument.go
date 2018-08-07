package argue

import (
	"sort"
)

// Argument represents the rules and information that
// argue must follow when parsing command-line
// arguments.
type Argument struct {
	Description string
	Version     string
	Facts       []*Fact
	ShowDesc    bool
	ShowVersion bool
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

// AddFact adds a new fact to the argument with the
// given parameters.
func (agmt *Argument) AddFact(ft FactType, name string, help string, v interface{}) *Fact {
	fact := NewFact(ft, help, name, determineShortName(*agmt, name), false, false, v)
	agmt.Facts = append(agmt.Facts, &fact)
	agmt.SortFacts()
	return &fact
}

// NameInFlagFacts accepts a name and checks if that
// name exists within the existing facts of the
// received argument. This checks both the short and
// long names.
func (agmt Argument) NameInFlagFacts(name string) (*Fact, bool) {
	for _, f := range agmt.FlagFacts() {
		if name == f.FullName || name == string(f.ShortName) {
			return f, true
		}
	}
	return &Fact{}, false
}

// NameInPositonalFacts accepts a name and checks if
// that name exists within the existing facts of the
// received argument. This checks both the short and
// long names.
func (agmt Argument) NameInPositonalFacts(name string) (*Fact, bool) {
	for _, f := range agmt.PositionalFacts() {
		if name == f.FullName || name == string(f.ShortName) {
			return f, true
		}
	}
	return &Fact{}, false
}

// Propose will parse the command line arguments and
// determine if they align with the facts in the
// received Argument type. The values of facts will
// be populated accordingly. Propose accepts a
// boolean that determines if the proposed argument
// must succeed or not. If set, the usage will be
// written to the standard output and the program
// will exit with error code 1.
func (agmt Argument) Propose() bool {
	ps, fm := splitArguments(agmt)

	// Check if all required positional arguments are present
	if agmt.NumRequiredPositional() > 0 {
		pfs := agmt.PositionalFacts()
		lowest := 0
		var lastPositional *Fact
		for i, pf := range pfs {
			if pf.Required {
				lastPositional = pf
				lowest = i
			}
		}

		if len(ps) <= lowest {
			agmt.PrintError("positional argument " + lastPositional.UpperName() + " is missing")
		}
	}

	// Check if all required flags are present.
	for _, f := range agmt.RequiredFacts() {
		if f.Positional {
			continue
		}

		var contains bool
		for k := range fm {
			if k == f.FullName || byte(k[0]) == f.ShortName {
				contains = true
			}
		}

		if !contains {
			agmt.PrintError("required flag --" + f.FullName + " is missing")
		}
	}

	for k, v := range fm {
		f, _ := agmt.NameInFlagFacts(k)
		setFactValue(agmt, f, v)
	}

	counter := 0
	for _, v := range ps {
		facts := agmt.PositionalFacts()
		f := facts[counter]
		setFactValue(agmt, f, v)
		counter++
	}

	return true
}

// SortFacts sorts the facts in an argument by fact
// type.
func (agmt *Argument) SortFacts() {
	sort.Slice(agmt.Facts, func(i, j int) bool {
		return agmt.Facts[i].ShortName < agmt.Facts[j].ShortName
	})
}

// ContainsShortName iterates through the facts of
// the received argument and returns true if any of
// the facts contain the short name provided, false
// otherwise.
func (agmt Argument) ContainsShortName(b byte) bool {
	for _, f := range agmt.Facts {
		if f.ShortName == b {
			return true
		}
	}
	return false
}

// NumPositional returns the number of positional
// facts within the received argument.
func (agmt Argument) NumPositional() int {
	var count int
	for _, f := range agmt.Facts {
		if f.Positional {
			count++
		}
	}
	return count
}

// PositionalFacts returns a slice of all the
// positional facts in the received argument.
func (agmt Argument) PositionalFacts() []*Fact {
	var facts []*Fact
	for _, f := range agmt.Facts {
		if f.Positional {
			facts = append(facts, f)
		}
	}
	return facts
}

// NumFlags returns the number of optional facts
// within the received argument.
func (agmt Argument) NumFlags() int {
	var count int
	for _, f := range agmt.Facts {
		if !f.Positional {
			count++
		}
	}
	return count
}

// FlagFacts returns a slice of all the facts that
// are a flag in the received argument.
func (agmt Argument) FlagFacts() []*Fact {
	var facts []*Fact
	for _, f := range agmt.Facts {
		if !f.Positional {
			facts = append(facts, f)
		}
	}
	return facts
}

// NumRequired returns the number of required facts
// within the received argument.
func (agmt Argument) NumRequired() int {
	var count int
	for _, f := range agmt.Facts {
		if f.Required {
			count++
		}
	}
	return count
}

// NumRequiredPositional returns the number of
// positional facts that are also required within
// the received argument.
func (agmt Argument) NumRequiredPositional() int {
	var count int
	for _, f := range agmt.Facts {
		if f.Required && f.Positional {
			count++
		}
	}
	return count
}

// RequiredFacts returns a slice of all the facts
// that are required in the received argument.
func (agmt Argument) RequiredFacts() []*Fact {
	var facts []*Fact
	for _, f := range agmt.Facts {
		if f.Required {
			facts = append(facts, f)
		}
	}
	return facts
}

// RequiredPositionalFacts returns a slice of all the
// facts that are required and positional in the
// received argument.
func (agmt Argument) RequiredPositionalFacts() []*Fact {
	var facts []*Fact
	for _, f := range agmt.Facts {
		if f.Required && f.Positional {
			facts = append(facts, f)
		}
	}
	return facts
}
