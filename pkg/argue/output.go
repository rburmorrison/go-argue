package argue

import (
	"fmt"
	"strings"
)

func printFact(s int, f Fact) {
	p1 := fmt.Sprintf("  -%s, --%s", string(f.ShortName), f.FullName)
	if f.ShortName == 0 {
		p1 = fmt.Sprintf("  --%s", f.FullName)
	}
	p1 += strings.Repeat(" ", s-len(p1)-1)
	p1 += f.Help
	fmt.Println(p1)
}

// PrintUsage writes the usage information of the
// recieved argument to the standard output.
func (agmt Argument) PrintUsage() {
	width := 30
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
			printFact(width, *f)
		}
	}

	// Only show optional arguments if they exist
	fmt.Println()
	fmt.Println("Flags:")
	for _, f := range agmt.FlagFacts() {
		printFact(width, *f)
	}
	// fmt.Println("  -h, --help\tdisplay this help and exit")
	// fmt.Println("  --version\tdisplay version and exit")
}

// PrintVersion writes the version of the program
// to the standard output in the form of "<name>
// <version>"
func (agmt Argument) PrintVersion() {
	fmt.Printf("%v %v\n", getBinaryName(), agmt.Version)
}
