package argue

import (
	"fmt"
	"strings"
)

// PrintUsage writes the usage information of the
// recieved Lawyer to the standard output.
func (l Lawyer) PrintUsage() {
	spacing := 4

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
	// -> flagWidth := 11 // Start with --help width
	// -> if l.ShowVersion {
	// -> 	flagWidth = 14 // Change to --version if showing
	// -> }

	if l.ShowVersion {
		fmt.Println(binaryName() + " " + l.Version)
	}

	if l.ShowDesc {
		fmt.Println(l.Description + "\n")
	}

	// Temporary
	fmt.Println("Flags:")
	fmt.Println("  -h --help       display this help and exit")
	fmt.Println("  -v --version    display version and exit")
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
