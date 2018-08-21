package argue

// SubArgument represents one argument in a pool of
// arguments, typically managed by a Lawyer.
type SubArgument struct {
	Name     string
	Help     string
	Argument Argument
}
