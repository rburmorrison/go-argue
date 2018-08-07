package argue

import (
	"fmt"
	"os"
	"strings"
)

func printFact(s int, f Fact) {
	var p1 string
	if !f.Positional {
		p1 = fmt.Sprintf("  -%s, --%s", string(f.ShortName), f.FullName)
		if f.ShortName == 0 {
			p1 = fmt.Sprintf("  --%s", f.FullName)
		}

		if f.Type != FactTypeBool {
			p1 += " VALUE"
		}
	} else {
		p1 = fmt.Sprintf("  %s", f.UpperName())
	}

	p1 += strings.Repeat(" ", s-len(p1))
	p1 += f.Help
	fmt.Println(p1)
}

// PrintError accepts a message and will print it
// as an error message along with the received
// argument's usage line. This will exit the program
// with an error code of 1.
func (agmt Argument) PrintError(msg string) {
	fmt.Printf("Error: %v\n", msg)
	fmt.Printf("Run \"%v --help\" to see usage information\n", os.Args[0])
	os.Exit(1)
}

// PrintUsage writes the usage information of the
// recieved argument to the standard output.
func (agmt Argument) PrintUsage() {
	width := 0
	agmt.SortFacts()
	for _, f := range agmt.Facts {
		s := fmt.Sprintf("  -%s, --%s VALUE", string(f.ShortName), f.FullName)
		if len(s) > width {
			width = len(s) + 5
		}
	}

	if agmt.ShowVersion {
		agmt.PrintVersion()
	}

	if agmt.ShowDesc {
		fmt.Println(agmt.Description)
		fmt.Println()
	}

	// Print usage line
	fmt.Printf("Usage: %v", os.Args[0])
	for _, f := range agmt.FlagFacts() {
		if f.Type == FactTypeBool {
			fmt.Printf(" [--%v]", f.FullName)
		} else {
			fmt.Printf(" [--%v VALUE]", f.FullName)
		}
	}

	for _, f := range agmt.PositionalFacts() {
		fmt.Printf(" %s", f.UpperName())
	}
	fmt.Println()

	// Only show positional arguments if they exist
	if agmt.NumPositional() > 0 {
		fmt.Println()
		fmt.Println("Positional arguments:")
		for _, f := range agmt.PositionalFacts() {
			printFact(width, *f)
		}
	}

	// Only show optional arguments if they exist
	fmt.Println()
	fmt.Println("Flags:")
	for _, f := range agmt.FlagFacts() {
		printFact(width, *f)
	}

	// Print default fact information
	var emptyI bool
	printFact(width, NewFact(FactTypeBool, "display this help and exit", "help", byte("h"[0]), false, false, &emptyI))
	printFact(width, NewFact(FactTypeBool, "display version and exit", "version", byte("v"[0]), false, false, &emptyI))
}

// PrintVersion writes the version of the program
// to the standard output in the form of "<name>
// <version>"
func (agmt Argument) PrintVersion() {
	fmt.Printf("%v %v\n", getBinaryName(), agmt.Version)
}
