package argue

import (
	"fmt"
)

// Argument represents the rules and information that
// argue must follow when parsing command-line
// arguments.
type Argument struct {
	Description string
	Version     string
	Facts       []*Fact
}

// NewArgument accepts a description and will return
// a new Argument with that description and default
// values.
func NewArgument(desc string, version string) Argument {
	var agmt Argument
	agmt.Description = desc
	agmt.Version = version

	return agmt
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
	return &boolFact
}

// Propose accepts a boolean that determines if the
// proposed argument must succeed or not.
func (agmt Argument) Propose(ms bool) bool {
	return true
}

// PrintUsage writes the usage information of the
// recieved argument to the standard output.
func (agmt Argument) PrintUsage() {
	agmt.PrintVersion()
	fmt.Printf("Usage: %v [--example-arg]\n", getBinaryName())
	fmt.Println()
	fmt.Println("Positional arguments:")
	fmt.Println()
	fmt.Println("Flags:")
}

// PrintVersion writes the version of the program
// to the standard output in the form of "<name>
// <version>"
func (agmt Argument) PrintVersion() {
	fmt.Printf("%v %v\n", getBinaryName(), agmt.Version)
}
