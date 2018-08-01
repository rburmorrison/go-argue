package argue

// Argument represents the rules and information that
// argue must follow when parsing command-line
// arguments.
type Argument struct {
	Description string
	Version     string
	Facts       []Fact
}

// NewArgument accepts a description and will return
// a new Argument with that description and default
// values.
func NewArgument(desc string) Argument {
	var agmt Argument
	agmt.Description = desc

	return agmt
}

// Propose accepts a boolean that determines if the
// proposed argument must succeed or not.
func (agmt Argument) Propose(ms bool) bool {
	return true
}

// PrintUsage writes the usage information of the
// recieved argument to the standard output.
func (agmt Argument) PrintUsage() {
}
