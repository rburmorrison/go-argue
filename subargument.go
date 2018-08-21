package argue

// SubArgument represents one argument in a pool of
// arguments, typically managed by a Lawyer.
type SubArgument struct {
	Name     string
	Help     string
	Argument Argument
	handler  func(arg *Argument)
}

// SetHandler sets the Handler field of SubArgument
// to the function provided
func (sa *SubArgument) SetHandler(f func(arg *Argument)) {
	sa.handler = f
}
