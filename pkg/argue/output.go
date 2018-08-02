package argue

import (
	"fmt"
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
		replacer := strings.NewReplacer(" ", "", "-", "")
		name := replacer.Replace(f.FullName)
		p1 = fmt.Sprintf("  %s", strings.ToUpper(name))
	}

	p1 += strings.Repeat(" ", s-len(p1))
	p1 += f.Help
	fmt.Println(p1)
}

// PrintUsage writes the usage information of the
// recieved argument to the standard output.
func (agmt Argument) PrintUsage() {
	agmt.SortFacts()
	width := 0
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

	fmt.Printf("Usage: %v", getBinaryName())
	for _, f := range agmt.FlagFacts() {
		if f.Type == FactTypeBool {
			fmt.Printf(" [--%v]", f.FullName)
		} else {
			fmt.Printf(" [--%v VALUE]", f.FullName)
		}
	}

	for _, f := range agmt.PositionalFacts() {
		replacer := strings.NewReplacer(" ", "", "-", "")
		name := replacer.Replace(f.FullName)
		fmt.Printf(" %s", strings.ToUpper(name))
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
	printFact(width, NewFact(FactTypeBool, "display this help and exit", "help", byte("h"[0]), false, false))
	printFact(width, NewFact(FactTypeBool, "display version and exit", "version", byte("v"[0]), false, false))
}

// PrintVersion writes the version of the program
// to the standard output in the form of "<name>
// <version>"
func (agmt Argument) PrintVersion() {
	fmt.Printf("%v %v\n", getBinaryName(), agmt.Version)
}
