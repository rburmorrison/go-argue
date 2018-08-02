package argue

import "fmt"

// PrintUsage writes the usage information of the
// recieved argument to the standard output.
func (agmt Argument) PrintUsage() {
	if agmt.ShowVersion {
		agmt.PrintVersion()
	}

	if agmt.ShowDesc {
		fmt.Println(agmt.Description)
		fmt.Println()
	}

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
