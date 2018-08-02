package argue

import (
	"fmt"
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
	return Argument{}
}

// AddBool adds a new bool fact to the argument with
// the given parameters.
func (agmt *Argument) AddBool(name string, help string) *Fact {
	var boolFact Fact
	boolFact.Type = FactTypeBool
	boolFact.FullName = name
	boolFact.ShortName = name[0]
	boolFact.Positional = false
	boolFact.Required = false

	agmt.Facts = append(agmt.Facts, &boolFact)
	agmt.SortFacts()
	return &boolFact
}

// Propose accepts a boolean that determines if the
// proposed argument must succeed or not.
func (agmt Argument) Propose(ms bool) bool {
	return true
}

// SortFacts sorts the facts in an argument by fact
// type.
func (agmt *Argument) SortFacts() {
	sort.Slice(agmt.Facts, func(i, j int) bool {
		return int(agmt.Facts[i].Type) < int(agmt.Facts[j].Type)
	})
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

// PrintUsage writes the usage information of the
// recieved argument to the standard output.
func (agmt Argument) PrintUsage() {
	agmt.PrintVersion()
	fmt.Println(agmt.Description)
	fmt.Println()
	fmt.Printf("Usage: %v", getBinaryName())
	for _, f := range agmt.Facts {
		fmt.Printf(" [--%v]", f.FullName)
	}
	fmt.Println()

	// Only show positional arguments if they exist
	if agmt.NumPositional() > 0 {
		fmt.Println()
		fmt.Println("Positional arguments:")
		for _, f := range agmt.PositionalFacts() {
			fmt.Printf("  %s\n", f.InfoString())
		}
	}

	// Only show optional arguments if they exist
	if agmt.NumFlags() > 0 {
		fmt.Println()
		fmt.Println("Flags:")
		for _, f := range agmt.FlagFacts() {
			fmt.Printf("  %s\n", f.InfoString())
		}
	}
}

// PrintVersion writes the version of the program
// to the standard output in the form of "<name>
// <version>"
func (agmt Argument) PrintVersion() {
	fmt.Printf("%v %v\n", getBinaryName(), agmt.Version)
}
