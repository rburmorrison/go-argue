package argue

// SubArgument represents one argument in a pool of
// arguments, typically managed by a Lawyer.
type SubArgument struct {
	Name     string
	Help     string
	Argument Argument
	handler  func(v interface{})
}

// SetHandler sets the Handler field of SubArgument
// to the function provided. If the argument that the
// SubArgument represents was auto-generated from a
// struct, that struct will be passed to the handler
// function. If it was created manually, nil will be
// passed to the handler instead.
func (sa *SubArgument) SetHandler(f func(v interface{})) {
	sa.handler = f
}
