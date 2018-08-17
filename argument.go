package argue

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

// Error Definitions
var (
	ErrExtraPositionals   = errors.New("argue: too many positional arguments provided")
	ErrMissingPositionals = errors.New("argue: not enough positional arguments provided")
	ErrUnknownFlag        = errors.New("argue: dispute found an unknown flag while parsing")
	ErrWrongType          = errors.New("argue: fact was not able to set a value due to mismatched types")
	ErrNilValue           = errors.New("argue: nil was passed to a flag")
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

// Dispute uses the facts in the received argument
// to parse the passed arguments. An error will be
// returned if Dispute fails to parse the arguments
// as it expects to. Optionally, setting "strict" to
// true will automatically print an error message to
// the console and exit the program on failing.
func (a Argument) Dispute(arguments []string, strict bool) error {
	ps, fm := a.SplitArguments(arguments)

	// Handle printing help and version if they exist
	for k := range fm {
		if k == "-h" || k == "--help" {
			a.PrintUsage()
			os.Exit(0)
		}

		if a.ShowVersion && (k == "-v" || k == "--version") {
			a.PrintVersion()
			os.Exit(0)
		}
	}

	// Check for unknown flags
	for k := range fm {
		_, nok := a.DressedNameExists(k)
		_, iok := a.DressedInitialExists(k)
		if !nok && !iok {
			if strict {
				a.PrintError("unknown flag " + k + " provided")
			}

			return ErrUnknownFlag
		}
	}

	// Check for extra positional arguments
	if len(ps) > len(a.PositionalFacts) {
		if strict {
			a.PrintError("too many positional arguments provided")
		}

		return ErrExtraPositionals
	}

	// Make sure the last required argument is satisfied
	lastPos := -1
	for i, f := range a.PositionalFacts {
		if f.Required {
			lastPos = i
		}
	}

	if lastPos > -1 && len(ps) < lastPos+1 {
		if strict {
			if lastPos+1 == 1 {
				a.PrintError(fmt.Sprintf("expected %d positional argument, but got %d", lastPos+1, len(ps)))
			} else {
				a.PrintError(fmt.Sprintf("expected %d positional arguments, but got %d", lastPos+1, len(ps)))
			}
		}

		return ErrMissingPositionals
	}

	// Set values for positional facts
	for i, s := range ps {
		fact := a.PositionalFacts[i]
		err := fact.SetValue(s)
		if err != nil {
			if strict {
				a.PrintError("positional argument " + UpperFactName(fact.Name) + " " + err.Error())
			}

			return ErrWrongType
		}
	}

	// Set values for flag facts
	for k, v := range fm {
		// Check if the value provided was nil
		if v == nil {
			if strict {
				a.PrintError("no value was provided for " + k)
			}

			return ErrNilValue
		}

		// Get the fact that correspons with the key
		f, nameExists := a.DressedNameExists(k)
		if !nameExists {
			f2, _ := a.DressedInitialExists(k)
			f = f2
		}

		err := f.SetValue(v)
		if err != nil {
			if strict {
				a.PrintError(k + " " + err.Error())
			}

			return ErrWrongType
		}
	}

	return nil
}

// SplitArguments splits command-line arguments into
// their "positional" and "flag" categories. They are
// returned in that order. The passed arguments
// should not include the call to the binary.
func (a Argument) SplitArguments(arguments []string) ([]string, map[string]interface{}) {
	// Define regular expressions
	flagReg := regexp.MustCompile(`^(-\S|--\S+)$`)

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

// RequiredPositionals returns all positional facts
// marked as required in the received arguments.
func (a Argument) RequiredPositionals() []*Fact {
	var facts []*Fact
	for _, f := range a.PositionalFacts {
		if f.Required {
			facts = append(facts, f)
		}
	}

	return facts
}

// RequiredFlags returns all flag facts marked as
// required in the received arguments.
func (a Argument) RequiredFlags() []*Fact {
	var facts []*Fact
	for _, f := range a.FlagFacts {
		if f.Required {
			facts = append(facts, f)
		}
	}

	return facts
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

	fact := NewFact(help, name, 0, true, true, v)
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
