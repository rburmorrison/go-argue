package argue

// Argument represents the rules and information that
// argue must follow when parsing command-line
// arguments.
type Argument struct {
	Description string
}

// NewArgument accepts a description and will return
// a new Argument with that description and default
// values.
func NewArgument(desc string) Argument {
	var agmt Argument
	agmt.Description = desc

	return agmt
}
