package argue

import (
	"fmt"
	"strings"
)

// PrintUsage writes the usage information of the
// recieved Lawyer to the standard output.
func (l Lawyer) PrintUsage() {
	spacing := 4

	// Define help and version flags
	var dummy bool
	helpFact := NewFact("display this help and exit", "help", byte("h"[0]), false, false, &dummy)
	versionFact := NewFact("display version and exit", "version", byte("v"[0]), false, false, &dummy)
	factBank := append(l.defaultArgument.FlagFacts, &helpFact)
	if l.ShowVersion {
		factBank = append(factBank, &versionFact)
	}

	// Calculate the largest sub-command name for spacing
	// purposes
	cmdWidth := 0
	for _, sa := range l.SubArguments {
		length := len(sa.Name) + 2
		if length > cmdWidth {
			cmdWidth = length
		}
	}

	// Calculate the largest flag header for spacing
	// purposes
	flagWidth := 0 // Start with --help width
	for _, f := range factBank {
		length := len(f.usageHeader()) + 2
		if length > flagWidth {
			flagWidth = length
		}
	}

	if l.ShowVersion {
		l.PrintVersion()
	}

	if l.ShowDesc {
		fmt.Println(l.Description + "\n")
	}

	// Print flags
	fmt.Println("Flags:")
	for _, f := range factBank {
		header := "  " + f.usageHeader()
		extra := flagWidth - len(header)
		space := strings.Repeat(" ", spacing+extra)
		fmt.Println(header + space + f.Help)
	}
	fmt.Println()

	// Display sub-commands
	fmt.Println("Commands:")
	for _, sa := range l.SubArguments {
		n := strings.ToLower(sa.Name)
		header := "  " + n
		extra := cmdWidth - len(header)
		space := strings.Repeat(" ", spacing+extra)
		fmt.Println(header + space + sa.Help)
	}

	fmt.Println()
	fmt.Printf("Run '%s <command> --help' for details about a command.\n", binaryName())
}

// PrintVersion prints the version specified by the
// Lawyer.
func (l Lawyer) PrintVersion() {
	fmt.Println(binaryName() + " " + l.Version)
}
