package argue

import (
	"fmt"
	"os"
	"strings"
)

func printFact(w int, s int, f Fact) {
	header := "  " + f.usageHeader()
	extra := w - len(header)
	space := strings.Repeat(" ", s+extra)
	fmt.Println(header + space + f.Help)
}

// PrintError accepts a message and will print it
// as an error message along with the received
// Argument's usage line. This will exit the program
// with an error code of 1.
func (a Argument) PrintError(msg string) {
	fmt.Printf("Error: %v\n", msg)
	fmt.Printf("Run \"%v", os.Args[0])
	if a.commandSuffix != "" {
		fmt.Print(" " + a.commandSuffix)
	}
	fmt.Print(" --help\" to see usage information\n")
	os.Exit(1)
}

// PrintUsage writes the usage information of the
// received Argument to the standard output.
func (a Argument) PrintUsage() {
	spacing := 4

	// Add help and version fact definitons and a dummy
	// variable to satisfy the NewFact signature
	var dummy bool
	helpFact := NewFact("display this help and exit", "help", byte("h"[0]), false, false, &dummy)
	versionFact := NewFact("display version and exit", "version", byte("v"[0]), false, false, &dummy)
	factBank := append(a.Facts(), &helpFact)
	if a.ShowVersion {
		factBank = append(factBank, &versionFact)
	}

	// Check which fact has the longest length to use as
	// a baseline for spacing
	a.SortFlagFacts()
	width := 0
	for _, f := range factBank {
		l := len(f.usageHeader()) + 2
		if l > width {
			width = l
		}
	}

	if a.ShowVersion {
		a.PrintVersion()
	}

	if a.ShowDesc {
		fmt.Println(a.Description + "\n")
	}

	// Print usage line
	fmt.Printf("Usage: %s", os.Args[0])
	if a.commandSuffix != "" {
		fmt.Print(" " + a.commandSuffix)
	}
	for _, f := range a.FlagFacts {
		if f.Type == FactTypeBool {
			fmt.Printf(" [--%s]", f.Name)
		} else {
			fmt.Printf(" [--%s VALUE]", f.Name)
		}
	}

	for _, f := range a.PositionalFacts {
		fmt.Printf(" %s", UpperFactName(f.Name))
	}
	fmt.Println()

	// Display positional facts
	if len(a.PositionalFacts) > 0 {
		fmt.Println()
		fmt.Println("Positional arguments:")
		for _, f := range a.PositionalFacts {
			printFact(width, spacing, *f)
		}
	}

	fmt.Println()
	fmt.Println("Flags:")
	for _, f := range a.FlagFacts {
		printFact(width, spacing, *f)
	}

	// Print default fact information
	printFact(width, spacing, helpFact)
	if a.ShowVersion {
		printFact(width, spacing, versionFact)
	}
}

// PrintVersion writes the version of the program to
// the standard output in the form of "<name>
// <version>"
func (a Argument) PrintVersion() {
	fmt.Printf("%v %v\n", binaryName(), a.Version)
}
